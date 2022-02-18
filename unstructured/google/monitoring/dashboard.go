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
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Dashboard struct{}

func DashboardToUnstructured(r *dclService.Dashboard) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "ga",
			Type:    "Dashboard",
		},
		Object: make(map[string]interface{}),
	}
	if r.ColumnLayout != nil && r.ColumnLayout != dclService.EmptyDashboardColumnLayout {
		rColumnLayout := make(map[string]interface{})
		var rColumnLayoutColumns []interface{}
		for _, rColumnLayoutColumnsVal := range r.ColumnLayout.Columns {
			rColumnLayoutColumnsObject := make(map[string]interface{})
			if rColumnLayoutColumnsVal.Weight != nil {
				rColumnLayoutColumnsObject["weight"] = *rColumnLayoutColumnsVal.Weight
			}
			var rColumnLayoutColumnsValWidgets []interface{}
			for _, rColumnLayoutColumnsValWidgetsVal := range rColumnLayoutColumnsVal.Widgets {
				rColumnLayoutColumnsValWidgetsObject := make(map[string]interface{})
				if rColumnLayoutColumnsValWidgetsVal.Blank != nil && rColumnLayoutColumnsValWidgetsVal.Blank != dclService.EmptyDashboardWidgetBlank {
					rColumnLayoutColumnsValWidgetsValBlank := make(map[string]interface{})
					rColumnLayoutColumnsValWidgetsObject["blank"] = rColumnLayoutColumnsValWidgetsValBlank
				}
				if rColumnLayoutColumnsValWidgetsVal.Scorecard != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard != dclService.EmptyDashboardWidgetScorecard {
					rColumnLayoutColumnsValWidgetsValScorecard := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView != dclService.EmptyDashboardWidgetScorecardGaugeView {
						rColumnLayoutColumnsValWidgetsValScorecardGaugeView := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.LowerBound != nil {
							rColumnLayoutColumnsValWidgetsValScorecardGaugeView["lowerBound"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.LowerBound
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.UpperBound != nil {
							rColumnLayoutColumnsValWidgetsValScorecardGaugeView["upperBound"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.UpperBound
						}
						rColumnLayoutColumnsValWidgetsValScorecard["gaugeView"] = rColumnLayoutColumnsValWidgetsValScorecardGaugeView
					}
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView != dclService.EmptyDashboardWidgetScorecardSparkChartView {
						rColumnLayoutColumnsValWidgetsValScorecardSparkChartView := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
							rColumnLayoutColumnsValWidgetsValScorecardSparkChartView["minAlignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.SparkChartType != nil {
							rColumnLayoutColumnsValWidgetsValScorecardSparkChartView["sparkChartType"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.SparkChartType)
						}
						rColumnLayoutColumnsValWidgetsValScorecard["sparkChartView"] = rColumnLayoutColumnsValWidgetsValScorecardSparkChartView
					}
					var rColumnLayoutColumnsValWidgetsValScorecardThresholds []interface{}
					for _, rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.Thresholds {
						rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Color != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["color"] = string(*rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Color)
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Direction != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["direction"] = string(*rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Direction)
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Label != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["label"] = *rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Label
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Value != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["value"] = *rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Value
						}
						rColumnLayoutColumnsValWidgetsValScorecardThresholds = append(rColumnLayoutColumnsValWidgetsValScorecardThresholds, rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject)
					}
					rColumnLayoutColumnsValWidgetsValScorecard["thresholds"] = rColumnLayoutColumnsValWidgetsValScorecardThresholds
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardWidgetScorecardTimeSeriesQuery {
						rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
							}
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
							}
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride != nil {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["unitOverride"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride
						}
						rColumnLayoutColumnsValWidgetsValScorecard["timeSeriesQuery"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery
					}
					rColumnLayoutColumnsValWidgetsObject["scorecard"] = rColumnLayoutColumnsValWidgetsValScorecard
				}
				if rColumnLayoutColumnsValWidgetsVal.Text != nil && rColumnLayoutColumnsValWidgetsVal.Text != dclService.EmptyDashboardWidgetText {
					rColumnLayoutColumnsValWidgetsValText := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.Text.Content != nil {
						rColumnLayoutColumnsValWidgetsValText["content"] = *rColumnLayoutColumnsValWidgetsVal.Text.Content
					}
					if rColumnLayoutColumnsValWidgetsVal.Text.Format != nil {
						rColumnLayoutColumnsValWidgetsValText["format"] = string(*rColumnLayoutColumnsValWidgetsVal.Text.Format)
					}
					rColumnLayoutColumnsValWidgetsObject["text"] = rColumnLayoutColumnsValWidgetsValText
				}
				if rColumnLayoutColumnsValWidgetsVal.Title != nil {
					rColumnLayoutColumnsValWidgetsObject["title"] = *rColumnLayoutColumnsValWidgetsVal.Title
				}
				if rColumnLayoutColumnsValWidgetsVal.XyChart != nil && rColumnLayoutColumnsValWidgetsVal.XyChart != dclService.EmptyDashboardWidgetXyChart {
					rColumnLayoutColumnsValWidgetsValXyChart := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions != dclService.EmptyDashboardWidgetXyChartChartOptions {
						rColumnLayoutColumnsValWidgetsValXyChartChartOptions := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions.Mode != nil {
							rColumnLayoutColumnsValWidgetsValXyChartChartOptions["mode"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions.Mode)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["chartOptions"] = rColumnLayoutColumnsValWidgetsValXyChartChartOptions
					}
					var rColumnLayoutColumnsValWidgetsValXyChartDataSets []interface{}
					for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal := range rColumnLayoutColumnsValWidgetsVal.XyChart.DataSets {
						rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.LegendTemplate != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["legendTemplate"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.LegendTemplate
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["minAlignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.PlotType != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["plotType"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.PlotType)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
								}
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
										}
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
										}
										var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
										for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
										}
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
										}
										var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
										for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
								}
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
							}
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["timeSeriesQuery"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery
						}
						rColumnLayoutColumnsValWidgetsValXyChartDataSets = append(rColumnLayoutColumnsValWidgetsValXyChartDataSets, rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject)
					}
					rColumnLayoutColumnsValWidgetsValXyChart["dataSets"] = rColumnLayoutColumnsValWidgetsValXyChartDataSets
					var rColumnLayoutColumnsValWidgetsValXyChartThresholds []interface{}
					for _, rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal := range rColumnLayoutColumnsValWidgetsVal.XyChart.Thresholds {
						rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Color != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["color"] = string(*rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Color)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Direction != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Direction)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["label"] = *rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Label
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Value != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["value"] = *rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Value
						}
						rColumnLayoutColumnsValWidgetsValXyChartThresholds = append(rColumnLayoutColumnsValWidgetsValXyChartThresholds, rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject)
					}
					rColumnLayoutColumnsValWidgetsValXyChart["thresholds"] = rColumnLayoutColumnsValWidgetsValXyChartThresholds
					if rColumnLayoutColumnsValWidgetsVal.XyChart.TimeshiftDuration != nil {
						rColumnLayoutColumnsValWidgetsValXyChart["timeshiftDuration"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.TimeshiftDuration
					}
					if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis != dclService.EmptyDashboardWidgetXyChartXAxis {
						rColumnLayoutColumnsValWidgetsValXyChartXAxis := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartXAxis["label"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Label
						}
						if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Scale != nil {
							rColumnLayoutColumnsValWidgetsValXyChartXAxis["scale"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Scale)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["xAxis"] = rColumnLayoutColumnsValWidgetsValXyChartXAxis
					}
					if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis != dclService.EmptyDashboardWidgetXyChartYAxis {
						rColumnLayoutColumnsValWidgetsValXyChartYAxis := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartYAxis["label"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Label
						}
						if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Scale != nil {
							rColumnLayoutColumnsValWidgetsValXyChartYAxis["scale"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Scale)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["yAxis"] = rColumnLayoutColumnsValWidgetsValXyChartYAxis
					}
					rColumnLayoutColumnsValWidgetsObject["xyChart"] = rColumnLayoutColumnsValWidgetsValXyChart
				}
				rColumnLayoutColumnsValWidgets = append(rColumnLayoutColumnsValWidgets, rColumnLayoutColumnsValWidgetsObject)
			}
			rColumnLayoutColumnsObject["widgets"] = rColumnLayoutColumnsValWidgets
			rColumnLayoutColumns = append(rColumnLayoutColumns, rColumnLayoutColumnsObject)
		}
		rColumnLayout["columns"] = rColumnLayoutColumns
		u.Object["columnLayout"] = rColumnLayout
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.GridLayout != nil && r.GridLayout != dclService.EmptyDashboardGridLayout {
		rGridLayout := make(map[string]interface{})
		if r.GridLayout.Columns != nil {
			rGridLayout["columns"] = *r.GridLayout.Columns
		}
		var rGridLayoutWidgets []interface{}
		for _, rGridLayoutWidgetsVal := range r.GridLayout.Widgets {
			rGridLayoutWidgets = append(rGridLayoutWidgets, DashboardWidgetToUnstructured(&rGridLayoutWidgetsVal))
		}
		rGridLayout["widgets"] = rGridLayoutWidgets
		u.Object["gridLayout"] = rGridLayout
	}
	if r.MosaicLayout != nil && r.MosaicLayout != dclService.EmptyDashboardMosaicLayout {
		rMosaicLayout := make(map[string]interface{})
		if r.MosaicLayout.Columns != nil {
			rMosaicLayout["columns"] = *r.MosaicLayout.Columns
		}
		var rMosaicLayoutTiles []interface{}
		for _, rMosaicLayoutTilesVal := range r.MosaicLayout.Tiles {
			rMosaicLayoutTilesObject := make(map[string]interface{})
			if rMosaicLayoutTilesVal.Height != nil {
				rMosaicLayoutTilesObject["height"] = *rMosaicLayoutTilesVal.Height
			}
			if rMosaicLayoutTilesVal.Widget != nil {
				rMosaicLayoutTilesObject["widget"] = DashboardWidgetToUnstructured(rMosaicLayoutTilesVal.Widget)
			}
			if rMosaicLayoutTilesVal.Width != nil {
				rMosaicLayoutTilesObject["width"] = *rMosaicLayoutTilesVal.Width
			}
			if rMosaicLayoutTilesVal.XPos != nil {
				rMosaicLayoutTilesObject["xPos"] = *rMosaicLayoutTilesVal.XPos
			}
			if rMosaicLayoutTilesVal.YPos != nil {
				rMosaicLayoutTilesObject["yPos"] = *rMosaicLayoutTilesVal.YPos
			}
			rMosaicLayoutTiles = append(rMosaicLayoutTiles, rMosaicLayoutTilesObject)
		}
		rMosaicLayout["tiles"] = rMosaicLayoutTiles
		u.Object["mosaicLayout"] = rMosaicLayout
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RowLayout != nil && r.RowLayout != dclService.EmptyDashboardRowLayout {
		rRowLayout := make(map[string]interface{})
		var rRowLayoutRows []interface{}
		for _, rRowLayoutRowsVal := range r.RowLayout.Rows {
			rRowLayoutRowsObject := make(map[string]interface{})
			if rRowLayoutRowsVal.Weight != nil {
				rRowLayoutRowsObject["weight"] = *rRowLayoutRowsVal.Weight
			}
			var rRowLayoutRowsValWidgets []interface{}
			for _, rRowLayoutRowsValWidgetsVal := range rRowLayoutRowsVal.Widgets {
				rRowLayoutRowsValWidgets = append(rRowLayoutRowsValWidgets, DashboardWidgetToUnstructured(&rRowLayoutRowsValWidgetsVal))
			}
			rRowLayoutRowsObject["widgets"] = rRowLayoutRowsValWidgets
			rRowLayoutRows = append(rRowLayoutRows, rRowLayoutRowsObject)
		}
		rRowLayout["rows"] = rRowLayoutRows
		u.Object["rowLayout"] = rRowLayout
	}
	return u
}

func DashboardWidgetToUnstructured(r *dclService.DashboardWidget) map[string]interface{} {
	result := make(map[string]interface{})
	if r.Blank != nil && r.Blank != dclService.EmptyDashboardWidgetBlank {
		rBlank := make(map[string]interface{})
		result["blank"] = rBlank
	}
	if r.Scorecard != nil && r.Scorecard != dclService.EmptyDashboardWidgetScorecard {
		rScorecard := make(map[string]interface{})
		if r.Scorecard.GaugeView != nil && r.Scorecard.GaugeView != dclService.EmptyDashboardWidgetScorecardGaugeView {
			rScorecardGaugeView := make(map[string]interface{})
			if r.Scorecard.GaugeView.LowerBound != nil {
				rScorecardGaugeView["lowerBound"] = *r.Scorecard.GaugeView.LowerBound
			}
			if r.Scorecard.GaugeView.UpperBound != nil {
				rScorecardGaugeView["upperBound"] = *r.Scorecard.GaugeView.UpperBound
			}
			rScorecard["gaugeView"] = rScorecardGaugeView
		}
		if r.Scorecard.SparkChartView != nil && r.Scorecard.SparkChartView != dclService.EmptyDashboardWidgetScorecardSparkChartView {
			rScorecardSparkChartView := make(map[string]interface{})
			if r.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
				rScorecardSparkChartView["minAlignmentPeriod"] = *r.Scorecard.SparkChartView.MinAlignmentPeriod
			}
			if r.Scorecard.SparkChartView.SparkChartType != nil {
				rScorecardSparkChartView["sparkChartType"] = string(*r.Scorecard.SparkChartView.SparkChartType)
			}
			rScorecard["sparkChartView"] = rScorecardSparkChartView
		}
		var rScorecardThresholds []interface{}
		for _, rScorecardThresholdsVal := range r.Scorecard.Thresholds {
			rScorecardThresholdsObject := make(map[string]interface{})
			if rScorecardThresholdsVal.Color != nil {
				rScorecardThresholdsObject["color"] = string(*rScorecardThresholdsVal.Color)
			}
			if rScorecardThresholdsVal.Direction != nil {
				rScorecardThresholdsObject["direction"] = string(*rScorecardThresholdsVal.Direction)
			}
			if rScorecardThresholdsVal.Label != nil {
				rScorecardThresholdsObject["label"] = *rScorecardThresholdsVal.Label
			}
			if rScorecardThresholdsVal.Value != nil {
				rScorecardThresholdsObject["value"] = *rScorecardThresholdsVal.Value
			}
			rScorecardThresholds = append(rScorecardThresholds, rScorecardThresholdsObject)
		}
		rScorecard["thresholds"] = rScorecardThresholds
		if r.Scorecard.TimeSeriesQuery != nil && r.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardWidgetScorecardTimeSeriesQuery {
			rScorecardTimeSeriesQuery := make(map[string]interface{})
			if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
				rScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
					rScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
					}
					var rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
					for _, rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
						rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rScorecardTimeSeriesQueryTimeSeriesFilterAggregation
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
					rScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
					rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
					rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
					}
					var rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
					for _, rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
						rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
				}
				rScorecardTimeSeriesQuery["timeSeriesFilter"] = rScorecardTimeSeriesQueryTimeSeriesFilter
			}
			if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
				rScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
					rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
						}
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
						}
						var rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
						for _, rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
						}
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
						}
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
					rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
						}
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
						}
						var rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
						for _, rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
						}
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
						if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
							rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
						}
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
					rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
				}
				if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
					rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
					}
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
					}
					var rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
					for _, rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
					if r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
						rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
					}
					rScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
				}
				rScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rScorecardTimeSeriesQueryTimeSeriesFilterRatio
			}
			if r.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
				rScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *r.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
			}
			if r.Scorecard.TimeSeriesQuery.UnitOverride != nil {
				rScorecardTimeSeriesQuery["unitOverride"] = *r.Scorecard.TimeSeriesQuery.UnitOverride
			}
			rScorecard["timeSeriesQuery"] = rScorecardTimeSeriesQuery
		}
		result["scorecard"] = rScorecard
	}
	if r.Text != nil && r.Text != dclService.EmptyDashboardWidgetText {
		rText := make(map[string]interface{})
		if r.Text.Content != nil {
			rText["content"] = *r.Text.Content
		}
		if r.Text.Format != nil {
			rText["format"] = string(*r.Text.Format)
		}
		result["text"] = rText
	}
	if r.Title != nil {
		result["title"] = *r.Title
	}
	if r.XyChart != nil && r.XyChart != dclService.EmptyDashboardWidgetXyChart {
		rXyChart := make(map[string]interface{})
		if r.XyChart.ChartOptions != nil && r.XyChart.ChartOptions != dclService.EmptyDashboardWidgetXyChartChartOptions {
			rXyChartChartOptions := make(map[string]interface{})
			if r.XyChart.ChartOptions.Mode != nil {
				rXyChartChartOptions["mode"] = string(*r.XyChart.ChartOptions.Mode)
			}
			rXyChart["chartOptions"] = rXyChartChartOptions
		}
		var rXyChartDataSets []interface{}
		for _, rXyChartDataSetsVal := range r.XyChart.DataSets {
			rXyChartDataSetsObject := make(map[string]interface{})
			if rXyChartDataSetsVal.LegendTemplate != nil {
				rXyChartDataSetsObject["legendTemplate"] = *rXyChartDataSetsVal.LegendTemplate
			}
			if rXyChartDataSetsVal.MinAlignmentPeriod != nil {
				rXyChartDataSetsObject["minAlignmentPeriod"] = *rXyChartDataSetsVal.MinAlignmentPeriod
			}
			if rXyChartDataSetsVal.PlotType != nil {
				rXyChartDataSetsObject["plotType"] = string(*rXyChartDataSetsVal.PlotType)
			}
			if rXyChartDataSetsVal.TimeSeriesQuery != nil && rXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery {
				rXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
				if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
					rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
						}
						var rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
						for _, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
						}
						var rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
						for _, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
					}
					rXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
				}
				if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
					rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
							}
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
							}
							var rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
							for _, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
							}
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
							}
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
							}
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
							}
							var rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
							for _, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
							}
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
							if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
								rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
							}
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
					}
					if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
						}
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
						}
						var rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
						for _, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
						if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
							rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
						}
						rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
					}
					rXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
				}
				if rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
					rXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
				}
				if rXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
					rXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
				}
				rXyChartDataSetsObject["timeSeriesQuery"] = rXyChartDataSetsValTimeSeriesQuery
			}
			rXyChartDataSets = append(rXyChartDataSets, rXyChartDataSetsObject)
		}
		rXyChart["dataSets"] = rXyChartDataSets
		var rXyChartThresholds []interface{}
		for _, rXyChartThresholdsVal := range r.XyChart.Thresholds {
			rXyChartThresholdsObject := make(map[string]interface{})
			if rXyChartThresholdsVal.Color != nil {
				rXyChartThresholdsObject["color"] = string(*rXyChartThresholdsVal.Color)
			}
			if rXyChartThresholdsVal.Direction != nil {
				rXyChartThresholdsObject["direction"] = string(*rXyChartThresholdsVal.Direction)
			}
			if rXyChartThresholdsVal.Label != nil {
				rXyChartThresholdsObject["label"] = *rXyChartThresholdsVal.Label
			}
			if rXyChartThresholdsVal.Value != nil {
				rXyChartThresholdsObject["value"] = *rXyChartThresholdsVal.Value
			}
			rXyChartThresholds = append(rXyChartThresholds, rXyChartThresholdsObject)
		}
		rXyChart["thresholds"] = rXyChartThresholds
		if r.XyChart.TimeshiftDuration != nil {
			rXyChart["timeshiftDuration"] = *r.XyChart.TimeshiftDuration
		}
		if r.XyChart.XAxis != nil && r.XyChart.XAxis != dclService.EmptyDashboardWidgetXyChartXAxis {
			rXyChartXAxis := make(map[string]interface{})
			if r.XyChart.XAxis.Label != nil {
				rXyChartXAxis["label"] = *r.XyChart.XAxis.Label
			}
			if r.XyChart.XAxis.Scale != nil {
				rXyChartXAxis["scale"] = string(*r.XyChart.XAxis.Scale)
			}
			rXyChart["xAxis"] = rXyChartXAxis
		}
		if r.XyChart.YAxis != nil && r.XyChart.YAxis != dclService.EmptyDashboardWidgetXyChartYAxis {
			rXyChartYAxis := make(map[string]interface{})
			if r.XyChart.YAxis.Label != nil {
				rXyChartYAxis["label"] = *r.XyChart.YAxis.Label
			}
			if r.XyChart.YAxis.Scale != nil {
				rXyChartYAxis["scale"] = string(*r.XyChart.YAxis.Scale)
			}
			rXyChart["yAxis"] = rXyChartYAxis
		}
		result["xyChart"] = rXyChart
	}
	return result
}

func UnstructuredToDashboard(u *unstructured.Resource) (*dclService.Dashboard, error) {
	r := &dclService.Dashboard{}
	if _, ok := u.Object["columnLayout"]; ok {
		if rColumnLayout, ok := u.Object["columnLayout"].(map[string]interface{}); ok {
			r.ColumnLayout = &dclService.DashboardColumnLayout{}
			if _, ok := rColumnLayout["columns"]; ok {
				if s, ok := rColumnLayout["columns"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rColumnLayoutColumns dclService.DashboardColumnLayoutColumns
							if _, ok := objval["weight"]; ok {
								if i, ok := objval["weight"].(int64); ok {
									rColumnLayoutColumns.Weight = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rColumnLayoutColumns.Weight: expected int64")
								}
							}
							if _, ok := objval["widgets"]; ok {
								if s, ok := objval["widgets"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rColumnLayoutColumnsWidgets dclService.DashboardWidget
											if _, ok := objval["blank"]; ok {
												if _, ok := objval["blank"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Blank = &dclService.DashboardWidgetBlank{}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Blank: expected map[string]interface{}")
												}
											}
											if _, ok := objval["scorecard"]; ok {
												if rColumnLayoutColumnsWidgetsScorecard, ok := objval["scorecard"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Scorecard = &dclService.DashboardWidgetScorecard{}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["gaugeView"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardGaugeView, ok := rColumnLayoutColumnsWidgetsScorecard["gaugeView"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.GaugeView = &dclService.DashboardWidgetScorecardGaugeView{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["lowerBound"]; ok {
																if f, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["lowerBound"].(float64); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView.LowerBound: expected float64")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["upperBound"]; ok {
																if f, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["upperBound"].(float64); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView.UpperBound: expected float64")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["sparkChartView"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardSparkChartView, ok := rColumnLayoutColumnsWidgetsScorecard["sparkChartView"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.SparkChartView = &dclService.DashboardWidgetScorecardSparkChartView{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["minAlignmentPeriod"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["sparkChartType"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["sparkChartType"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.SparkChartType = dclService.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.SparkChartType: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["thresholds"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsScorecard["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsScorecardThresholds dclService.DashboardWidgetScorecardThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Color = dclService.DashboardWidgetScorecardThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Direction = dclService.DashboardWidgetScorecardThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Value: expected float64")
																		}
																	}
																	rColumnLayoutColumnsWidgets.Scorecard.Thresholds = append(rColumnLayoutColumnsWidgets.Scorecard.Thresholds, rColumnLayoutColumnsWidgetsScorecardThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["timeSeriesQuery"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery, ok := rColumnLayoutColumnsWidgetsScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery = &dclService.DashboardWidgetScorecardTimeSeriesQuery{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
																if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																		if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																				if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																				if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["unitOverride"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["unitOverride"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard: expected map[string]interface{}")
												}
											}
											if _, ok := objval["text"]; ok {
												if rColumnLayoutColumnsWidgetsText, ok := objval["text"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Text = &dclService.DashboardWidgetText{}
													if _, ok := rColumnLayoutColumnsWidgetsText["content"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsText["content"].(string); ok {
															rColumnLayoutColumnsWidgets.Text.Content = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text.Content: expected string")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsText["format"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsText["format"].(string); ok {
															rColumnLayoutColumnsWidgets.Text.Format = dclService.DashboardWidgetTextFormatEnumRef(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text.Format: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text: expected map[string]interface{}")
												}
											}
											if _, ok := objval["title"]; ok {
												if s, ok := objval["title"].(string); ok {
													rColumnLayoutColumnsWidgets.Title = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Title: expected string")
												}
											}
											if _, ok := objval["xyChart"]; ok {
												if rColumnLayoutColumnsWidgetsXyChart, ok := objval["xyChart"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.XyChart = &dclService.DashboardWidgetXyChart{}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["chartOptions"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartChartOptions, ok := rColumnLayoutColumnsWidgetsXyChart["chartOptions"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.ChartOptions = &dclService.DashboardWidgetXyChartChartOptions{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartChartOptions["mode"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartChartOptions["mode"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.ChartOptions.Mode = dclService.DashboardWidgetXyChartChartOptionsModeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.ChartOptions.Mode: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.ChartOptions: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["dataSets"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["dataSets"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsXyChartDataSets dclService.DashboardWidgetXyChartDataSets
																	if _, ok := objval["legendTemplate"]; ok {
																		if s, ok := objval["legendTemplate"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.LegendTemplate = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.LegendTemplate: expected string")
																		}
																	}
																	if _, ok := objval["minAlignmentPeriod"]; ok {
																		if s, ok := objval["minAlignmentPeriod"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.MinAlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := objval["plotType"]; ok {
																		if s, ok := objval["plotType"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.PlotType = dclService.DashboardWidgetXyChartDataSetsPlotTypeEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.PlotType: expected string")
																		}
																	}
																	if _, ok := objval["timeSeriesQuery"]; ok {
																		if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
																				if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																				if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																								if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																								if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
																		}
																	}
																	rColumnLayoutColumnsWidgets.XyChart.DataSets = append(rColumnLayoutColumnsWidgets.XyChart.DataSets, rColumnLayoutColumnsWidgetsXyChartDataSets)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.DataSets: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["thresholds"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsXyChartThresholds dclService.DashboardWidgetXyChartThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Color = dclService.DashboardWidgetXyChartThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Direction = dclService.DashboardWidgetXyChartThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Value: expected float64")
																		}
																	}
																	rColumnLayoutColumnsWidgets.XyChart.Thresholds = append(rColumnLayoutColumnsWidgets.XyChart.Thresholds, rColumnLayoutColumnsWidgetsXyChartThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["timeshiftDuration"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["timeshiftDuration"].(string); ok {
															rColumnLayoutColumnsWidgets.XyChart.TimeshiftDuration = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.TimeshiftDuration: expected string")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["xAxis"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartXAxis, ok := rColumnLayoutColumnsWidgetsXyChart["xAxis"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.XAxis = &dclService.DashboardWidgetXyChartXAxis{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["label"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["label"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.XAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis.Label: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["scale"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["scale"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.XAxis.Scale = dclService.DashboardWidgetXyChartXAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["yAxis"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartYAxis, ok := rColumnLayoutColumnsWidgetsXyChart["yAxis"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.YAxis = &dclService.DashboardWidgetXyChartYAxis{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["label"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["label"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.YAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis.Label: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["scale"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["scale"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.YAxis.Scale = dclService.DashboardWidgetXyChartYAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart: expected map[string]interface{}")
												}
											}
											rColumnLayoutColumns.Widgets = append(rColumnLayoutColumns.Widgets, rColumnLayoutColumnsWidgets)
										}
									}
								} else {
									return nil, fmt.Errorf("rColumnLayoutColumns.Widgets: expected []interface{}")
								}
							}
							r.ColumnLayout.Columns = append(r.ColumnLayout.Columns, rColumnLayoutColumns)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ColumnLayout.Columns: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ColumnLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["gridLayout"]; ok {
		if rGridLayout, ok := u.Object["gridLayout"].(map[string]interface{}); ok {
			r.GridLayout = &dclService.DashboardGridLayout{}
			if _, ok := rGridLayout["columns"]; ok {
				if i, ok := rGridLayout["columns"].(int64); ok {
					r.GridLayout.Columns = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.GridLayout.Columns: expected int64")
				}
			}
			if _, ok := rGridLayout["widgets"]; ok {
				if s, ok := rGridLayout["widgets"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							unstructuredObjval, err := UnstructuredToDashboardWidget(objval)
							if err != nil {
								return nil, err
							}
							r.GridLayout.Widgets = append(r.GridLayout.Widgets, *unstructuredObjval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.GridLayout.Widgets: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GridLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["mosaicLayout"]; ok {
		if rMosaicLayout, ok := u.Object["mosaicLayout"].(map[string]interface{}); ok {
			r.MosaicLayout = &dclService.DashboardMosaicLayout{}
			if _, ok := rMosaicLayout["columns"]; ok {
				if i, ok := rMosaicLayout["columns"].(int64); ok {
					r.MosaicLayout.Columns = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.MosaicLayout.Columns: expected int64")
				}
			}
			if _, ok := rMosaicLayout["tiles"]; ok {
				if s, ok := rMosaicLayout["tiles"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMosaicLayoutTiles dclService.DashboardMosaicLayoutTiles
							if _, ok := objval["height"]; ok {
								if i, ok := objval["height"].(int64); ok {
									rMosaicLayoutTiles.Height = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Height: expected int64")
								}
							}
							if _, ok := objval["widget"]; ok {
								if rMosaicLayoutTilesWidget, ok := objval["widget"].(map[string]interface{}); ok {
									var err error
									rMosaicLayoutTiles.Widget, err = UnstructuredToDashboardWidget(rMosaicLayoutTilesWidget)
									if err != nil {
										return nil, err
									}
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Widget: expected map[string]interface{}")
								}
							}
							if _, ok := objval["width"]; ok {
								if i, ok := objval["width"].(int64); ok {
									rMosaicLayoutTiles.Width = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Width: expected int64")
								}
							}
							if _, ok := objval["xPos"]; ok {
								if i, ok := objval["xPos"].(int64); ok {
									rMosaicLayoutTiles.XPos = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.XPos: expected int64")
								}
							}
							if _, ok := objval["yPos"]; ok {
								if i, ok := objval["yPos"].(int64); ok {
									rMosaicLayoutTiles.YPos = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.YPos: expected int64")
								}
							}
							r.MosaicLayout.Tiles = append(r.MosaicLayout.Tiles, rMosaicLayoutTiles)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MosaicLayout.Tiles: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MosaicLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["rowLayout"]; ok {
		if rRowLayout, ok := u.Object["rowLayout"].(map[string]interface{}); ok {
			r.RowLayout = &dclService.DashboardRowLayout{}
			if _, ok := rRowLayout["rows"]; ok {
				if s, ok := rRowLayout["rows"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rRowLayoutRows dclService.DashboardRowLayoutRows
							if _, ok := objval["weight"]; ok {
								if i, ok := objval["weight"].(int64); ok {
									rRowLayoutRows.Weight = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rRowLayoutRows.Weight: expected int64")
								}
							}
							if _, ok := objval["widgets"]; ok {
								if s, ok := objval["widgets"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											unstructuredObjval, err := UnstructuredToDashboardWidget(objval)
											if err != nil {
												return nil, err
											}
											rRowLayoutRows.Widgets = append(rRowLayoutRows.Widgets, *unstructuredObjval)
										}
									}
								} else {
									return nil, fmt.Errorf("rRowLayoutRows.Widgets: expected []interface{}")
								}
							}
							r.RowLayout.Rows = append(r.RowLayout.Rows, rRowLayoutRows)
						}
					}
				} else {
					return nil, fmt.Errorf("r.RowLayout.Rows: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RowLayout: expected map[string]interface{}")
		}
	}
	return r, nil
}

func UnstructuredToDashboardWidget(obj map[string]interface{}) (*dclService.DashboardWidget, error) {
	r := &dclService.DashboardWidget{}
	if _, ok := obj["blank"]; ok {
		if _, ok := obj["blank"].(map[string]interface{}); ok {
			r.Blank = &dclService.DashboardWidgetBlank{}
		} else {
			return nil, fmt.Errorf("r.Blank: expected map[string]interface{}")
		}
	}
	if _, ok := obj["scorecard"]; ok {
		if rScorecard, ok := obj["scorecard"].(map[string]interface{}); ok {
			r.Scorecard = &dclService.DashboardWidgetScorecard{}
			if _, ok := rScorecard["gaugeView"]; ok {
				if rScorecardGaugeView, ok := rScorecard["gaugeView"].(map[string]interface{}); ok {
					r.Scorecard.GaugeView = &dclService.DashboardWidgetScorecardGaugeView{}
					if _, ok := rScorecardGaugeView["lowerBound"]; ok {
						if f, ok := rScorecardGaugeView["lowerBound"].(float64); ok {
							r.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.Scorecard.GaugeView.LowerBound: expected float64")
						}
					}
					if _, ok := rScorecardGaugeView["upperBound"]; ok {
						if f, ok := rScorecardGaugeView["upperBound"].(float64); ok {
							r.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.Scorecard.GaugeView.UpperBound: expected float64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Scorecard.GaugeView: expected map[string]interface{}")
				}
			}
			if _, ok := rScorecard["sparkChartView"]; ok {
				if rScorecardSparkChartView, ok := rScorecard["sparkChartView"].(map[string]interface{}); ok {
					r.Scorecard.SparkChartView = &dclService.DashboardWidgetScorecardSparkChartView{}
					if _, ok := rScorecardSparkChartView["minAlignmentPeriod"]; ok {
						if s, ok := rScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
							r.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
						}
					}
					if _, ok := rScorecardSparkChartView["sparkChartType"]; ok {
						if s, ok := rScorecardSparkChartView["sparkChartType"].(string); ok {
							r.Scorecard.SparkChartView.SparkChartType = dclService.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Scorecard.SparkChartView.SparkChartType: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Scorecard.SparkChartView: expected map[string]interface{}")
				}
			}
			if _, ok := rScorecard["thresholds"]; ok {
				if s, ok := rScorecard["thresholds"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rScorecardThresholds dclService.DashboardWidgetScorecardThresholds
							if _, ok := objval["color"]; ok {
								if s, ok := objval["color"].(string); ok {
									rScorecardThresholds.Color = dclService.DashboardWidgetScorecardThresholdsColorEnumRef(s)
								} else {
									return nil, fmt.Errorf("rScorecardThresholds.Color: expected string")
								}
							}
							if _, ok := objval["direction"]; ok {
								if s, ok := objval["direction"].(string); ok {
									rScorecardThresholds.Direction = dclService.DashboardWidgetScorecardThresholdsDirectionEnumRef(s)
								} else {
									return nil, fmt.Errorf("rScorecardThresholds.Direction: expected string")
								}
							}
							if _, ok := objval["label"]; ok {
								if s, ok := objval["label"].(string); ok {
									rScorecardThresholds.Label = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rScorecardThresholds.Label: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if f, ok := objval["value"].(float64); ok {
									rScorecardThresholds.Value = dcl.Float64(f)
								} else {
									return nil, fmt.Errorf("rScorecardThresholds.Value: expected float64")
								}
							}
							r.Scorecard.Thresholds = append(r.Scorecard.Thresholds, rScorecardThresholds)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Scorecard.Thresholds: expected []interface{}")
				}
			}
			if _, ok := rScorecard["timeSeriesQuery"]; ok {
				if rScorecardTimeSeriesQuery, ok := rScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
					r.Scorecard.TimeSeriesQuery = &dclService.DashboardWidgetScorecardTimeSeriesQuery{}
					if _, ok := rScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
						if rScorecardTimeSeriesQueryTimeSeriesFilter, ok := rScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
							r.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
								if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
										if i, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
						}
					}
					if _, ok := rScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
						if rScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
							r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
										if rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
										if rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
												}
											}
											if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
												if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
										if i, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
								}
							}
							if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
								if rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
									r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
										}
									}
									if _, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
										if s, ok := rScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
											r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
						}
					}
					if _, ok := rScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
						if s, ok := rScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
							r.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
						}
					}
					if _, ok := rScorecardTimeSeriesQuery["unitOverride"]; ok {
						if s, ok := rScorecardTimeSeriesQuery["unitOverride"].(string); ok {
							r.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Scorecard: expected map[string]interface{}")
		}
	}
	if _, ok := obj["text"]; ok {
		if rText, ok := obj["text"].(map[string]interface{}); ok {
			r.Text = &dclService.DashboardWidgetText{}
			if _, ok := rText["content"]; ok {
				if s, ok := rText["content"].(string); ok {
					r.Text.Content = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Text.Content: expected string")
				}
			}
			if _, ok := rText["format"]; ok {
				if s, ok := rText["format"].(string); ok {
					r.Text.Format = dclService.DashboardWidgetTextFormatEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Text.Format: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Text: expected map[string]interface{}")
		}
	}
	if _, ok := obj["title"]; ok {
		if s, ok := obj["title"].(string); ok {
			r.Title = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Title: expected string")
		}
	}
	if _, ok := obj["xyChart"]; ok {
		if rXyChart, ok := obj["xyChart"].(map[string]interface{}); ok {
			r.XyChart = &dclService.DashboardWidgetXyChart{}
			if _, ok := rXyChart["chartOptions"]; ok {
				if rXyChartChartOptions, ok := rXyChart["chartOptions"].(map[string]interface{}); ok {
					r.XyChart.ChartOptions = &dclService.DashboardWidgetXyChartChartOptions{}
					if _, ok := rXyChartChartOptions["mode"]; ok {
						if s, ok := rXyChartChartOptions["mode"].(string); ok {
							r.XyChart.ChartOptions.Mode = dclService.DashboardWidgetXyChartChartOptionsModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.XyChart.ChartOptions.Mode: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.XyChart.ChartOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rXyChart["dataSets"]; ok {
				if s, ok := rXyChart["dataSets"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rXyChartDataSets dclService.DashboardWidgetXyChartDataSets
							if _, ok := objval["legendTemplate"]; ok {
								if s, ok := objval["legendTemplate"].(string); ok {
									rXyChartDataSets.LegendTemplate = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rXyChartDataSets.LegendTemplate: expected string")
								}
							}
							if _, ok := objval["minAlignmentPeriod"]; ok {
								if s, ok := objval["minAlignmentPeriod"].(string); ok {
									rXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rXyChartDataSets.MinAlignmentPeriod: expected string")
								}
							}
							if _, ok := objval["plotType"]; ok {
								if s, ok := objval["plotType"].(string); ok {
									rXyChartDataSets.PlotType = dclService.DashboardWidgetXyChartDataSetsPlotTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rXyChartDataSets.PlotType: expected string")
								}
							}
							if _, ok := objval["timeSeriesQuery"]; ok {
								if rXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
									rXyChartDataSets.TimeSeriesQuery = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
									if _, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
										if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
											rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
												if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
														if i, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
										}
									}
									if _, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
										if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
											rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
														if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
														if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																}
															}
															if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
														if i, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
												}
											}
											if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
												if rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
													rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
														}
													}
													if _, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
														if s, ok := rXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
															rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
														} else {
															return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
										}
									}
									if _, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
										if s, ok := rXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
											rXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
										}
									}
									if _, ok := rXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
										if s, ok := rXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
											rXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
								}
							}
							r.XyChart.DataSets = append(r.XyChart.DataSets, rXyChartDataSets)
						}
					}
				} else {
					return nil, fmt.Errorf("r.XyChart.DataSets: expected []interface{}")
				}
			}
			if _, ok := rXyChart["thresholds"]; ok {
				if s, ok := rXyChart["thresholds"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rXyChartThresholds dclService.DashboardWidgetXyChartThresholds
							if _, ok := objval["color"]; ok {
								if s, ok := objval["color"].(string); ok {
									rXyChartThresholds.Color = dclService.DashboardWidgetXyChartThresholdsColorEnumRef(s)
								} else {
									return nil, fmt.Errorf("rXyChartThresholds.Color: expected string")
								}
							}
							if _, ok := objval["direction"]; ok {
								if s, ok := objval["direction"].(string); ok {
									rXyChartThresholds.Direction = dclService.DashboardWidgetXyChartThresholdsDirectionEnumRef(s)
								} else {
									return nil, fmt.Errorf("rXyChartThresholds.Direction: expected string")
								}
							}
							if _, ok := objval["label"]; ok {
								if s, ok := objval["label"].(string); ok {
									rXyChartThresholds.Label = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rXyChartThresholds.Label: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if f, ok := objval["value"].(float64); ok {
									rXyChartThresholds.Value = dcl.Float64(f)
								} else {
									return nil, fmt.Errorf("rXyChartThresholds.Value: expected float64")
								}
							}
							r.XyChart.Thresholds = append(r.XyChart.Thresholds, rXyChartThresholds)
						}
					}
				} else {
					return nil, fmt.Errorf("r.XyChart.Thresholds: expected []interface{}")
				}
			}
			if _, ok := rXyChart["timeshiftDuration"]; ok {
				if s, ok := rXyChart["timeshiftDuration"].(string); ok {
					r.XyChart.TimeshiftDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.XyChart.TimeshiftDuration: expected string")
				}
			}
			if _, ok := rXyChart["xAxis"]; ok {
				if rXyChartXAxis, ok := rXyChart["xAxis"].(map[string]interface{}); ok {
					r.XyChart.XAxis = &dclService.DashboardWidgetXyChartXAxis{}
					if _, ok := rXyChartXAxis["label"]; ok {
						if s, ok := rXyChartXAxis["label"].(string); ok {
							r.XyChart.XAxis.Label = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.XyChart.XAxis.Label: expected string")
						}
					}
					if _, ok := rXyChartXAxis["scale"]; ok {
						if s, ok := rXyChartXAxis["scale"].(string); ok {
							r.XyChart.XAxis.Scale = dclService.DashboardWidgetXyChartXAxisScaleEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.XyChart.XAxis.Scale: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.XyChart.XAxis: expected map[string]interface{}")
				}
			}
			if _, ok := rXyChart["yAxis"]; ok {
				if rXyChartYAxis, ok := rXyChart["yAxis"].(map[string]interface{}); ok {
					r.XyChart.YAxis = &dclService.DashboardWidgetXyChartYAxis{}
					if _, ok := rXyChartYAxis["label"]; ok {
						if s, ok := rXyChartYAxis["label"].(string); ok {
							r.XyChart.YAxis.Label = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.XyChart.YAxis.Label: expected string")
						}
					}
					if _, ok := rXyChartYAxis["scale"]; ok {
						if s, ok := rXyChartYAxis["scale"].(string); ok {
							r.XyChart.YAxis.Scale = dclService.DashboardWidgetXyChartYAxisScaleEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.XyChart.YAxis.Scale: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.XyChart.YAxis: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.XyChart: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDashboard(ctx, r)
	if err != nil {
		return nil, err
	}
	return DashboardToUnstructured(r), nil
}

func ListDashboard(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDashboard(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DashboardToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDashboard(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDashboard(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DashboardToUnstructured(r), nil
}

func DashboardHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDashboard(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDashboard(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return err
	}
	return c.DeleteDashboard(ctx, r)
}

func DashboardID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Dashboard) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"Dashboard",
		"ga",
	}
}

func (r *Dashboard) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDashboard(ctx, config, resource)
}

func (r *Dashboard) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDashboard(ctx, config, resource, opts...)
}

func (r *Dashboard) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DashboardHasDiff(ctx, config, resource, opts...)
}

func (r *Dashboard) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDashboard(ctx, config, resource)
}

func (r *Dashboard) ID(resource *unstructured.Resource) (string, error) {
	return DashboardID(resource)
}

func init() {
	unstructured.Register(&Dashboard{})
}
