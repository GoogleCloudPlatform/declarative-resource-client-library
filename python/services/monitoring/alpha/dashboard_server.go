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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// Server implements the gRPC interface for Dashboard.
type DashboardServer struct{}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsPlotTypeEnum converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum) *alpha.DashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartDataSetsPlotTypeEnum(n[len("MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsColorEnum converts a DashboardWidgetXyChartThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum) *alpha.DashboardWidgetXyChartThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartThresholdsColorEnum(n[len("MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsDirectionEnum converts a DashboardWidgetXyChartThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum) *alpha.DashboardWidgetXyChartThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartThresholdsDirectionEnum(n[len("MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartXAxisScaleEnum converts a DashboardWidgetXyChartXAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum) *alpha.DashboardWidgetXyChartXAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartXAxisScaleEnum(n[len("MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartYAxisScaleEnum converts a DashboardWidgetXyChartYAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum) *alpha.DashboardWidgetXyChartYAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartYAxisScaleEnum(n[len("MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartChartOptionsModeEnum converts a DashboardWidgetXyChartChartOptionsModeEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum(e alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum) *alpha.DashboardWidgetXyChartChartOptionsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetXyChartChartOptionsModeEnum(n[len("MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) *alpha.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(n[len("MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsColorEnum converts a DashboardWidgetScorecardThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum) *alpha.DashboardWidgetScorecardThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardThresholdsColorEnum(n[len("MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsDirectionEnum converts a DashboardWidgetScorecardThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum(e alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum) *alpha.DashboardWidgetScorecardThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetScorecardThresholdsDirectionEnum(n[len("MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetTextFormatEnum converts a DashboardWidgetTextFormatEnum enum from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetTextFormatEnum(e alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum) *alpha.DashboardWidgetTextFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum_name[int32(e)]; ok {
		e := alpha.DashboardWidgetTextFormatEnum(n[len("MonitoringAlphaDashboardWidgetTextFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardGridLayout converts a DashboardGridLayout object from its proto representation.
func ProtoToMonitoringAlphaDashboardGridLayout(p *alphapb.MonitoringAlphaDashboardGridLayout) *alpha.DashboardGridLayout {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardGridLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringAlphaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayout converts a DashboardMosaicLayout object from its proto representation.
func ProtoToMonitoringAlphaDashboardMosaicLayout(p *alphapb.MonitoringAlphaDashboardMosaicLayout) *alpha.DashboardMosaicLayout {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardMosaicLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetTiles() {
		obj.Tiles = append(obj.Tiles, *ProtoToMonitoringAlphaDashboardMosaicLayoutTiles(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayoutTiles converts a DashboardMosaicLayoutTiles object from its proto representation.
func ProtoToMonitoringAlphaDashboardMosaicLayoutTiles(p *alphapb.MonitoringAlphaDashboardMosaicLayoutTiles) *alpha.DashboardMosaicLayoutTiles {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardMosaicLayoutTiles{
		XPos:   dcl.Int64OrNil(p.GetXPos()),
		YPos:   dcl.Int64OrNil(p.GetYPos()),
		Width:  dcl.Int64OrNil(p.GetWidth()),
		Height: dcl.Int64OrNil(p.GetHeight()),
		Widget: ProtoToMonitoringAlphaDashboardWidget(p.GetWidget()),
	}
	return obj
}

// ProtoToDashboardRowLayout converts a DashboardRowLayout object from its proto representation.
func ProtoToMonitoringAlphaDashboardRowLayout(p *alphapb.MonitoringAlphaDashboardRowLayout) *alpha.DashboardRowLayout {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardRowLayout{}
	for _, r := range p.GetRows() {
		obj.Rows = append(obj.Rows, *ProtoToMonitoringAlphaDashboardRowLayoutRows(r))
	}
	return obj
}

// ProtoToDashboardRowLayoutRows converts a DashboardRowLayoutRows object from its proto representation.
func ProtoToMonitoringAlphaDashboardRowLayoutRows(p *alphapb.MonitoringAlphaDashboardRowLayoutRows) *alpha.DashboardRowLayoutRows {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardRowLayoutRows{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringAlphaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardColumnLayout converts a DashboardColumnLayout object from its proto representation.
func ProtoToMonitoringAlphaDashboardColumnLayout(p *alphapb.MonitoringAlphaDashboardColumnLayout) *alpha.DashboardColumnLayout {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardColumnLayout{}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToMonitoringAlphaDashboardColumnLayoutColumns(r))
	}
	return obj
}

// ProtoToDashboardColumnLayoutColumns converts a DashboardColumnLayoutColumns object from its proto representation.
func ProtoToMonitoringAlphaDashboardColumnLayoutColumns(p *alphapb.MonitoringAlphaDashboardColumnLayoutColumns) *alpha.DashboardColumnLayoutColumns {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardColumnLayoutColumns{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringAlphaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardWidget converts a DashboardWidget object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidget(p *alphapb.MonitoringAlphaDashboardWidget) *alpha.DashboardWidget {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidget{
		Title:     dcl.StringOrNil(p.GetTitle()),
		XyChart:   ProtoToMonitoringAlphaDashboardWidgetXyChart(p.GetXyChart()),
		Scorecard: ProtoToMonitoringAlphaDashboardWidgetScorecard(p.GetScorecard()),
		Text:      ProtoToMonitoringAlphaDashboardWidgetText(p.GetText()),
		Blank:     ProtoToMonitoringAlphaDashboardWidgetBlank(p.GetBlank()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChart converts a DashboardWidgetXyChart object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChart(p *alphapb.MonitoringAlphaDashboardWidgetXyChart) *alpha.DashboardWidgetXyChart {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChart{
		TimeshiftDuration: dcl.StringOrNil(p.GetTimeshiftDuration()),
		XAxis:             ProtoToMonitoringAlphaDashboardWidgetXyChartXAxis(p.GetXAxis()),
		YAxis:             ProtoToMonitoringAlphaDashboardWidgetXyChartYAxis(p.GetYAxis()),
		ChartOptions:      ProtoToMonitoringAlphaDashboardWidgetXyChartChartOptions(p.GetChartOptions()),
	}
	for _, r := range p.GetDataSets() {
		obj.DataSets = append(obj.DataSets, *ProtoToMonitoringAlphaDashboardWidgetXyChartDataSets(r))
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringAlphaDashboardWidgetXyChartThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSets converts a DashboardWidgetXyChartDataSets object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSets(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSets) *alpha.DashboardWidgetXyChartDataSets {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSets{
		TimeSeriesQuery:    ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQuery(p.GetTimeSeriesQuery()),
		PlotType:           ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum(p.GetPlotType()),
		LegendTemplate:     dcl.StringOrNil(p.GetLegendTemplate()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQuery converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQuery(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQuery) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartThresholds converts a DashboardWidgetXyChartThresholds object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartThresholds(p *alphapb.MonitoringAlphaDashboardWidgetXyChartThresholds) *alpha.DashboardWidgetXyChartThresholds {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartXAxis converts a DashboardWidgetXyChartXAxis object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartXAxis(p *alphapb.MonitoringAlphaDashboardWidgetXyChartXAxis) *alpha.DashboardWidgetXyChartXAxis {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartXAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartYAxis converts a DashboardWidgetXyChartYAxis object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartYAxis(p *alphapb.MonitoringAlphaDashboardWidgetXyChartYAxis) *alpha.DashboardWidgetXyChartYAxis {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartYAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartChartOptions converts a DashboardWidgetXyChartChartOptions object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetXyChartChartOptions(p *alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptions) *alpha.DashboardWidgetXyChartChartOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetXyChartChartOptions{
		Mode: ProtoToMonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecard converts a DashboardWidgetScorecard object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecard(p *alphapb.MonitoringAlphaDashboardWidgetScorecard) *alpha.DashboardWidgetScorecard {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecard{
		TimeSeriesQuery: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQuery(p.GetTimeSeriesQuery()),
		GaugeView:       ProtoToMonitoringAlphaDashboardWidgetScorecardGaugeView(p.GetGaugeView()),
		SparkChartView:  ProtoToMonitoringAlphaDashboardWidgetScorecardSparkChartView(p.GetSparkChartView()),
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringAlphaDashboardWidgetScorecardThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQuery converts a DashboardWidgetScorecardTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQuery(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQuery) *alpha.DashboardWidgetScorecardTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardGaugeView converts a DashboardWidgetScorecardGaugeView object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardGaugeView(p *alphapb.MonitoringAlphaDashboardWidgetScorecardGaugeView) *alpha.DashboardWidgetScorecardGaugeView {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardGaugeView{
		LowerBound: dcl.Float64OrNil(p.GetLowerBound()),
		UpperBound: dcl.Float64OrNil(p.GetUpperBound()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardSparkChartView converts a DashboardWidgetScorecardSparkChartView object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardSparkChartView(p *alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartView) *alpha.DashboardWidgetScorecardSparkChartView {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardSparkChartView{
		SparkChartType:     ProtoToMonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(p.GetSparkChartType()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardThresholds converts a DashboardWidgetScorecardThresholds object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetScorecardThresholds(p *alphapb.MonitoringAlphaDashboardWidgetScorecardThresholds) *alpha.DashboardWidgetScorecardThresholds {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetScorecardThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetText converts a DashboardWidgetText object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetText(p *alphapb.MonitoringAlphaDashboardWidgetText) *alpha.DashboardWidgetText {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetText{
		Content: dcl.StringOrNil(p.GetContent()),
		Format:  ProtoToMonitoringAlphaDashboardWidgetTextFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToDashboardWidgetBlank converts a DashboardWidgetBlank object from its proto representation.
func ProtoToMonitoringAlphaDashboardWidgetBlank(p *alphapb.MonitoringAlphaDashboardWidgetBlank) *alpha.DashboardWidgetBlank {
	if p == nil {
		return nil
	}
	obj := &alpha.DashboardWidgetBlank{}
	return obj
}

// ProtoToDashboard converts a Dashboard resource from its proto representation.
func ProtoToDashboard(p *alphapb.MonitoringAlphaDashboard) *alpha.Dashboard {
	obj := &alpha.Dashboard{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		GridLayout:   ProtoToMonitoringAlphaDashboardGridLayout(p.GetGridLayout()),
		MosaicLayout: ProtoToMonitoringAlphaDashboardMosaicLayout(p.GetMosaicLayout()),
		RowLayout:    ProtoToMonitoringAlphaDashboardRowLayout(p.GetRowLayout()),
		ColumnLayout: ProtoToMonitoringAlphaDashboardColumnLayout(p.GetColumnLayout()),
		Project:      dcl.StringOrNil(p.GetProject()),
		Etag:         dcl.StringOrNil(p.GetEtag()),
	}
	return obj
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsPlotTypeEnumToProto converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(e *alpha.DashboardWidgetXyChartDataSetsPlotTypeEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum_value["DashboardWidgetXyChartDataSetsPlotTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
}

// DashboardWidgetXyChartThresholdsColorEnumToProto converts a DashboardWidgetXyChartThresholdsColorEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnumToProto(e *alpha.DashboardWidgetXyChartThresholdsColorEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum_value["DashboardWidgetXyChartThresholdsColorEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnum(0)
}

// DashboardWidgetXyChartThresholdsDirectionEnumToProto converts a DashboardWidgetXyChartThresholdsDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnumToProto(e *alpha.DashboardWidgetXyChartThresholdsDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum_value["DashboardWidgetXyChartThresholdsDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnum(0)
}

// DashboardWidgetXyChartXAxisScaleEnumToProto converts a DashboardWidgetXyChartXAxisScaleEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnumToProto(e *alpha.DashboardWidgetXyChartXAxisScaleEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum_value["DashboardWidgetXyChartXAxisScaleEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnum(0)
}

// DashboardWidgetXyChartYAxisScaleEnumToProto converts a DashboardWidgetXyChartYAxisScaleEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnumToProto(e *alpha.DashboardWidgetXyChartYAxisScaleEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum_value["DashboardWidgetXyChartYAxisScaleEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnum(0)
}

// DashboardWidgetXyChartChartOptionsModeEnumToProto converts a DashboardWidgetXyChartChartOptionsModeEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnumToProto(e *alpha.DashboardWidgetXyChartChartOptionsModeEnum) alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum_value["DashboardWidgetXyChartChartOptionsModeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(e *alpha.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_value["DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
}

// DashboardWidgetScorecardThresholdsColorEnumToProto converts a DashboardWidgetScorecardThresholdsColorEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnumToProto(e *alpha.DashboardWidgetScorecardThresholdsColorEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum_value["DashboardWidgetScorecardThresholdsColorEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnum(0)
}

// DashboardWidgetScorecardThresholdsDirectionEnumToProto converts a DashboardWidgetScorecardThresholdsDirectionEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnumToProto(e *alpha.DashboardWidgetScorecardThresholdsDirectionEnum) alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum_value["DashboardWidgetScorecardThresholdsDirectionEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnum(0)
}

// DashboardWidgetTextFormatEnumToProto converts a DashboardWidgetTextFormatEnum enum to its proto representation.
func MonitoringAlphaDashboardWidgetTextFormatEnumToProto(e *alpha.DashboardWidgetTextFormatEnum) alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum {
	if e == nil {
		return alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum_value["DashboardWidgetTextFormatEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum(v)
	}
	return alphapb.MonitoringAlphaDashboardWidgetTextFormatEnum(0)
}

// DashboardGridLayoutToProto converts a DashboardGridLayout object to its proto representation.
func MonitoringAlphaDashboardGridLayoutToProto(o *alpha.DashboardGridLayout) *alphapb.MonitoringAlphaDashboardGridLayout {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardGridLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sWidgets := make([]*alphapb.MonitoringAlphaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringAlphaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardMosaicLayoutToProto converts a DashboardMosaicLayout object to its proto representation.
func MonitoringAlphaDashboardMosaicLayoutToProto(o *alpha.DashboardMosaicLayout) *alphapb.MonitoringAlphaDashboardMosaicLayout {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardMosaicLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sTiles := make([]*alphapb.MonitoringAlphaDashboardMosaicLayoutTiles, len(o.Tiles))
	for i, r := range o.Tiles {
		sTiles[i] = MonitoringAlphaDashboardMosaicLayoutTilesToProto(&r)
	}
	p.SetTiles(sTiles)
	return p
}

// DashboardMosaicLayoutTilesToProto converts a DashboardMosaicLayoutTiles object to its proto representation.
func MonitoringAlphaDashboardMosaicLayoutTilesToProto(o *alpha.DashboardMosaicLayoutTiles) *alphapb.MonitoringAlphaDashboardMosaicLayoutTiles {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardMosaicLayoutTiles{}
	p.SetXPos(dcl.ValueOrEmptyInt64(o.XPos))
	p.SetYPos(dcl.ValueOrEmptyInt64(o.YPos))
	p.SetWidth(dcl.ValueOrEmptyInt64(o.Width))
	p.SetHeight(dcl.ValueOrEmptyInt64(o.Height))
	p.SetWidget(MonitoringAlphaDashboardWidgetToProto(o.Widget))
	return p
}

// DashboardRowLayoutToProto converts a DashboardRowLayout object to its proto representation.
func MonitoringAlphaDashboardRowLayoutToProto(o *alpha.DashboardRowLayout) *alphapb.MonitoringAlphaDashboardRowLayout {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardRowLayout{}
	sRows := make([]*alphapb.MonitoringAlphaDashboardRowLayoutRows, len(o.Rows))
	for i, r := range o.Rows {
		sRows[i] = MonitoringAlphaDashboardRowLayoutRowsToProto(&r)
	}
	p.SetRows(sRows)
	return p
}

// DashboardRowLayoutRowsToProto converts a DashboardRowLayoutRows object to its proto representation.
func MonitoringAlphaDashboardRowLayoutRowsToProto(o *alpha.DashboardRowLayoutRows) *alphapb.MonitoringAlphaDashboardRowLayoutRows {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardRowLayoutRows{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*alphapb.MonitoringAlphaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringAlphaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardColumnLayoutToProto converts a DashboardColumnLayout object to its proto representation.
func MonitoringAlphaDashboardColumnLayoutToProto(o *alpha.DashboardColumnLayout) *alphapb.MonitoringAlphaDashboardColumnLayout {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardColumnLayout{}
	sColumns := make([]*alphapb.MonitoringAlphaDashboardColumnLayoutColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = MonitoringAlphaDashboardColumnLayoutColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// DashboardColumnLayoutColumnsToProto converts a DashboardColumnLayoutColumns object to its proto representation.
func MonitoringAlphaDashboardColumnLayoutColumnsToProto(o *alpha.DashboardColumnLayoutColumns) *alphapb.MonitoringAlphaDashboardColumnLayoutColumns {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardColumnLayoutColumns{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*alphapb.MonitoringAlphaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringAlphaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardWidgetToProto converts a DashboardWidget object to its proto representation.
func MonitoringAlphaDashboardWidgetToProto(o *alpha.DashboardWidget) *alphapb.MonitoringAlphaDashboardWidget {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidget{}
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetXyChart(MonitoringAlphaDashboardWidgetXyChartToProto(o.XyChart))
	p.SetScorecard(MonitoringAlphaDashboardWidgetScorecardToProto(o.Scorecard))
	p.SetText(MonitoringAlphaDashboardWidgetTextToProto(o.Text))
	p.SetBlank(MonitoringAlphaDashboardWidgetBlankToProto(o.Blank))
	return p
}

// DashboardWidgetXyChartToProto converts a DashboardWidgetXyChart object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartToProto(o *alpha.DashboardWidgetXyChart) *alphapb.MonitoringAlphaDashboardWidgetXyChart {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChart{}
	p.SetTimeshiftDuration(dcl.ValueOrEmptyString(o.TimeshiftDuration))
	p.SetXAxis(MonitoringAlphaDashboardWidgetXyChartXAxisToProto(o.XAxis))
	p.SetYAxis(MonitoringAlphaDashboardWidgetXyChartYAxisToProto(o.YAxis))
	p.SetChartOptions(MonitoringAlphaDashboardWidgetXyChartChartOptionsToProto(o.ChartOptions))
	sDataSets := make([]*alphapb.MonitoringAlphaDashboardWidgetXyChartDataSets, len(o.DataSets))
	for i, r := range o.DataSets {
		sDataSets[i] = MonitoringAlphaDashboardWidgetXyChartDataSetsToProto(&r)
	}
	p.SetDataSets(sDataSets)
	sThresholds := make([]*alphapb.MonitoringAlphaDashboardWidgetXyChartThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringAlphaDashboardWidgetXyChartThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetXyChartDataSetsToProto converts a DashboardWidgetXyChartDataSets object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsToProto(o *alpha.DashboardWidgetXyChartDataSets) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSets {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSets{}
	p.SetTimeSeriesQuery(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetPlotType(MonitoringAlphaDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(o.PlotType))
	p.SetLegendTemplate(dcl.ValueOrEmptyString(o.LegendTemplate))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQuery) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *alpha.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringAlphaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartThresholdsToProto converts a DashboardWidgetXyChartThresholds object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartThresholdsToProto(o *alpha.DashboardWidgetXyChartThresholds) *alphapb.MonitoringAlphaDashboardWidgetXyChartThresholds {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringAlphaDashboardWidgetXyChartThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringAlphaDashboardWidgetXyChartThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartXAxisToProto converts a DashboardWidgetXyChartXAxis object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartXAxisToProto(o *alpha.DashboardWidgetXyChartXAxis) *alphapb.MonitoringAlphaDashboardWidgetXyChartXAxis {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartXAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringAlphaDashboardWidgetXyChartXAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartYAxisToProto converts a DashboardWidgetXyChartYAxis object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartYAxisToProto(o *alpha.DashboardWidgetXyChartYAxis) *alphapb.MonitoringAlphaDashboardWidgetXyChartYAxis {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartYAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringAlphaDashboardWidgetXyChartYAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartChartOptionsToProto converts a DashboardWidgetXyChartChartOptions object to its proto representation.
func MonitoringAlphaDashboardWidgetXyChartChartOptionsToProto(o *alpha.DashboardWidgetXyChartChartOptions) *alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetXyChartChartOptions{}
	p.SetMode(MonitoringAlphaDashboardWidgetXyChartChartOptionsModeEnumToProto(o.Mode))
	return p
}

// DashboardWidgetScorecardToProto converts a DashboardWidgetScorecard object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardToProto(o *alpha.DashboardWidgetScorecard) *alphapb.MonitoringAlphaDashboardWidgetScorecard {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecard{}
	p.SetTimeSeriesQuery(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetGaugeView(MonitoringAlphaDashboardWidgetScorecardGaugeViewToProto(o.GaugeView))
	p.SetSparkChartView(MonitoringAlphaDashboardWidgetScorecardSparkChartViewToProto(o.SparkChartView))
	sThresholds := make([]*alphapb.MonitoringAlphaDashboardWidgetScorecardThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringAlphaDashboardWidgetScorecardThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryToProto converts a DashboardWidgetScorecardTimeSeriesQuery object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQuery) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *alpha.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringAlphaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardGaugeViewToProto converts a DashboardWidgetScorecardGaugeView object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardGaugeViewToProto(o *alpha.DashboardWidgetScorecardGaugeView) *alphapb.MonitoringAlphaDashboardWidgetScorecardGaugeView {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardGaugeView{}
	p.SetLowerBound(dcl.ValueOrEmptyDouble(o.LowerBound))
	p.SetUpperBound(dcl.ValueOrEmptyDouble(o.UpperBound))
	return p
}

// DashboardWidgetScorecardSparkChartViewToProto converts a DashboardWidgetScorecardSparkChartView object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardSparkChartViewToProto(o *alpha.DashboardWidgetScorecardSparkChartView) *alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartView {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardSparkChartView{}
	p.SetSparkChartType(MonitoringAlphaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(o.SparkChartType))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetScorecardThresholdsToProto converts a DashboardWidgetScorecardThresholds object to its proto representation.
func MonitoringAlphaDashboardWidgetScorecardThresholdsToProto(o *alpha.DashboardWidgetScorecardThresholds) *alphapb.MonitoringAlphaDashboardWidgetScorecardThresholds {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetScorecardThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringAlphaDashboardWidgetScorecardThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringAlphaDashboardWidgetScorecardThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetTextToProto converts a DashboardWidgetText object to its proto representation.
func MonitoringAlphaDashboardWidgetTextToProto(o *alpha.DashboardWidgetText) *alphapb.MonitoringAlphaDashboardWidgetText {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetText{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetFormat(MonitoringAlphaDashboardWidgetTextFormatEnumToProto(o.Format))
	return p
}

// DashboardWidgetBlankToProto converts a DashboardWidgetBlank object to its proto representation.
func MonitoringAlphaDashboardWidgetBlankToProto(o *alpha.DashboardWidgetBlank) *alphapb.MonitoringAlphaDashboardWidgetBlank {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaDashboardWidgetBlank{}
	return p
}

// DashboardToProto converts a Dashboard resource to its proto representation.
func DashboardToProto(resource *alpha.Dashboard) *alphapb.MonitoringAlphaDashboard {
	p := &alphapb.MonitoringAlphaDashboard{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetGridLayout(MonitoringAlphaDashboardGridLayoutToProto(resource.GridLayout))
	p.SetMosaicLayout(MonitoringAlphaDashboardMosaicLayoutToProto(resource.MosaicLayout))
	p.SetRowLayout(MonitoringAlphaDashboardRowLayoutToProto(resource.RowLayout))
	p.SetColumnLayout(MonitoringAlphaDashboardColumnLayoutToProto(resource.ColumnLayout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))

	return p
}

// applyDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) applyDashboard(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaDashboardRequest) (*alphapb.MonitoringAlphaDashboard, error) {
	p := ProtoToDashboard(request.GetResource())
	res, err := c.ApplyDashboard(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DashboardToProto(res)
	return r, nil
}

// applyMonitoringAlphaDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) ApplyMonitoringAlphaDashboard(ctx context.Context, request *alphapb.ApplyMonitoringAlphaDashboardRequest) (*alphapb.MonitoringAlphaDashboard, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDashboard(ctx, cl, request)
}

// DeleteDashboard handles the gRPC request by passing it to the underlying Dashboard Delete() method.
func (s *DashboardServer) DeleteMonitoringAlphaDashboard(ctx context.Context, request *alphapb.DeleteMonitoringAlphaDashboardRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDashboard(ctx, ProtoToDashboard(request.GetResource()))

}

// ListMonitoringAlphaDashboard handles the gRPC request by passing it to the underlying DashboardList() method.
func (s *DashboardServer) ListMonitoringAlphaDashboard(ctx context.Context, request *alphapb.ListMonitoringAlphaDashboardRequest) (*alphapb.ListMonitoringAlphaDashboardResponse, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDashboard(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaDashboard
	for _, r := range resources.Items {
		rp := DashboardToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaDashboardResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDashboard(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
