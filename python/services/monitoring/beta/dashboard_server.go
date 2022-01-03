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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/beta/monitoring_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
)

// Server implements the gRPC interface for Dashboard.
type DashboardServer struct{}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsPlotTypeEnum converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum(e betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum) *beta.DashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartDataSetsPlotTypeEnum(n[len("MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsColorEnum converts a DashboardWidgetXyChartThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartThresholdsColorEnum(e betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum) *beta.DashboardWidgetXyChartThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartThresholdsColorEnum(n[len("MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsDirectionEnum converts a DashboardWidgetXyChartThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum(e betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum) *beta.DashboardWidgetXyChartThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartThresholdsDirectionEnum(n[len("MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartXAxisScaleEnum converts a DashboardWidgetXyChartXAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartXAxisScaleEnum(e betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum) *beta.DashboardWidgetXyChartXAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartXAxisScaleEnum(n[len("MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartYAxisScaleEnum converts a DashboardWidgetXyChartYAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartYAxisScaleEnum(e betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum) *beta.DashboardWidgetXyChartYAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartYAxisScaleEnum(n[len("MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartChartOptionsModeEnum converts a DashboardWidgetXyChartChartOptionsModeEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum(e betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum) *beta.DashboardWidgetXyChartChartOptionsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetXyChartChartOptionsModeEnum(n[len("MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(e betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) *beta.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(n[len("MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsColorEnum converts a DashboardWidgetScorecardThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardThresholdsColorEnum(e betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum) *beta.DashboardWidgetScorecardThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardThresholdsColorEnum(n[len("MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsDirectionEnum converts a DashboardWidgetScorecardThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum(e betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum) *beta.DashboardWidgetScorecardThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetScorecardThresholdsDirectionEnum(n[len("MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetTextFormatEnum converts a DashboardWidgetTextFormatEnum enum from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetTextFormatEnum(e betapb.MonitoringBetaDashboardWidgetTextFormatEnum) *beta.DashboardWidgetTextFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaDashboardWidgetTextFormatEnum_name[int32(e)]; ok {
		e := beta.DashboardWidgetTextFormatEnum(n[len("MonitoringBetaDashboardWidgetTextFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardGridLayout converts a DashboardGridLayout object from its proto representation.
func ProtoToMonitoringBetaDashboardGridLayout(p *betapb.MonitoringBetaDashboardGridLayout) *beta.DashboardGridLayout {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardGridLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringBetaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayout converts a DashboardMosaicLayout object from its proto representation.
func ProtoToMonitoringBetaDashboardMosaicLayout(p *betapb.MonitoringBetaDashboardMosaicLayout) *beta.DashboardMosaicLayout {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardMosaicLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetTiles() {
		obj.Tiles = append(obj.Tiles, *ProtoToMonitoringBetaDashboardMosaicLayoutTiles(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayoutTiles converts a DashboardMosaicLayoutTiles object from its proto representation.
func ProtoToMonitoringBetaDashboardMosaicLayoutTiles(p *betapb.MonitoringBetaDashboardMosaicLayoutTiles) *beta.DashboardMosaicLayoutTiles {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardMosaicLayoutTiles{
		XPos:   dcl.Int64OrNil(p.GetXPos()),
		YPos:   dcl.Int64OrNil(p.GetYPos()),
		Width:  dcl.Int64OrNil(p.GetWidth()),
		Height: dcl.Int64OrNil(p.GetHeight()),
		Widget: ProtoToMonitoringBetaDashboardWidget(p.GetWidget()),
	}
	return obj
}

// ProtoToDashboardRowLayout converts a DashboardRowLayout object from its proto representation.
func ProtoToMonitoringBetaDashboardRowLayout(p *betapb.MonitoringBetaDashboardRowLayout) *beta.DashboardRowLayout {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardRowLayout{}
	for _, r := range p.GetRows() {
		obj.Rows = append(obj.Rows, *ProtoToMonitoringBetaDashboardRowLayoutRows(r))
	}
	return obj
}

// ProtoToDashboardRowLayoutRows converts a DashboardRowLayoutRows object from its proto representation.
func ProtoToMonitoringBetaDashboardRowLayoutRows(p *betapb.MonitoringBetaDashboardRowLayoutRows) *beta.DashboardRowLayoutRows {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardRowLayoutRows{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringBetaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardColumnLayout converts a DashboardColumnLayout object from its proto representation.
func ProtoToMonitoringBetaDashboardColumnLayout(p *betapb.MonitoringBetaDashboardColumnLayout) *beta.DashboardColumnLayout {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardColumnLayout{}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToMonitoringBetaDashboardColumnLayoutColumns(r))
	}
	return obj
}

// ProtoToDashboardColumnLayoutColumns converts a DashboardColumnLayoutColumns object from its proto representation.
func ProtoToMonitoringBetaDashboardColumnLayoutColumns(p *betapb.MonitoringBetaDashboardColumnLayoutColumns) *beta.DashboardColumnLayoutColumns {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardColumnLayoutColumns{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringBetaDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardWidget converts a DashboardWidget object from its proto representation.
func ProtoToMonitoringBetaDashboardWidget(p *betapb.MonitoringBetaDashboardWidget) *beta.DashboardWidget {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidget{
		Title:     dcl.StringOrNil(p.GetTitle()),
		XyChart:   ProtoToMonitoringBetaDashboardWidgetXyChart(p.GetXyChart()),
		Scorecard: ProtoToMonitoringBetaDashboardWidgetScorecard(p.GetScorecard()),
		Text:      ProtoToMonitoringBetaDashboardWidgetText(p.GetText()),
		Blank:     ProtoToMonitoringBetaDashboardWidgetBlank(p.GetBlank()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChart converts a DashboardWidgetXyChart object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChart(p *betapb.MonitoringBetaDashboardWidgetXyChart) *beta.DashboardWidgetXyChart {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChart{
		TimeshiftDuration: dcl.StringOrNil(p.GetTimeshiftDuration()),
		XAxis:             ProtoToMonitoringBetaDashboardWidgetXyChartXAxis(p.GetXAxis()),
		YAxis:             ProtoToMonitoringBetaDashboardWidgetXyChartYAxis(p.GetYAxis()),
		ChartOptions:      ProtoToMonitoringBetaDashboardWidgetXyChartChartOptions(p.GetChartOptions()),
	}
	for _, r := range p.GetDataSets() {
		obj.DataSets = append(obj.DataSets, *ProtoToMonitoringBetaDashboardWidgetXyChartDataSets(r))
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringBetaDashboardWidgetXyChartThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSets converts a DashboardWidgetXyChartDataSets object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSets(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSets) *beta.DashboardWidgetXyChartDataSets {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSets{
		TimeSeriesQuery:    ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery(p.GetTimeSeriesQuery()),
		PlotType:           ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum(p.GetPlotType()),
		LegendTemplate:     dcl.StringOrNil(p.GetLegendTemplate()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQuery converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartThresholds converts a DashboardWidgetXyChartThresholds object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartThresholds(p *betapb.MonitoringBetaDashboardWidgetXyChartThresholds) *beta.DashboardWidgetXyChartThresholds {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringBetaDashboardWidgetXyChartThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartXAxis converts a DashboardWidgetXyChartXAxis object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartXAxis(p *betapb.MonitoringBetaDashboardWidgetXyChartXAxis) *beta.DashboardWidgetXyChartXAxis {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartXAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringBetaDashboardWidgetXyChartXAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartYAxis converts a DashboardWidgetXyChartYAxis object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartYAxis(p *betapb.MonitoringBetaDashboardWidgetXyChartYAxis) *beta.DashboardWidgetXyChartYAxis {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartYAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringBetaDashboardWidgetXyChartYAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartChartOptions converts a DashboardWidgetXyChartChartOptions object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetXyChartChartOptions(p *betapb.MonitoringBetaDashboardWidgetXyChartChartOptions) *beta.DashboardWidgetXyChartChartOptions {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetXyChartChartOptions{
		Mode: ProtoToMonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecard converts a DashboardWidgetScorecard object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecard(p *betapb.MonitoringBetaDashboardWidgetScorecard) *beta.DashboardWidgetScorecard {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecard{
		TimeSeriesQuery: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQuery(p.GetTimeSeriesQuery()),
		GaugeView:       ProtoToMonitoringBetaDashboardWidgetScorecardGaugeView(p.GetGaugeView()),
		SparkChartView:  ProtoToMonitoringBetaDashboardWidgetScorecardSparkChartView(p.GetSparkChartView()),
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringBetaDashboardWidgetScorecardThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQuery converts a DashboardWidgetScorecardTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQuery(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQuery) *beta.DashboardWidgetScorecardTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardGaugeView converts a DashboardWidgetScorecardGaugeView object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardGaugeView(p *betapb.MonitoringBetaDashboardWidgetScorecardGaugeView) *beta.DashboardWidgetScorecardGaugeView {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardGaugeView{
		LowerBound: dcl.Float64OrNil(p.GetLowerBound()),
		UpperBound: dcl.Float64OrNil(p.GetUpperBound()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardSparkChartView converts a DashboardWidgetScorecardSparkChartView object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardSparkChartView(p *betapb.MonitoringBetaDashboardWidgetScorecardSparkChartView) *beta.DashboardWidgetScorecardSparkChartView {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardSparkChartView{
		SparkChartType:     ProtoToMonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(p.GetSparkChartType()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardThresholds converts a DashboardWidgetScorecardThresholds object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetScorecardThresholds(p *betapb.MonitoringBetaDashboardWidgetScorecardThresholds) *beta.DashboardWidgetScorecardThresholds {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetScorecardThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringBetaDashboardWidgetScorecardThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetText converts a DashboardWidgetText object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetText(p *betapb.MonitoringBetaDashboardWidgetText) *beta.DashboardWidgetText {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetText{
		Content: dcl.StringOrNil(p.GetContent()),
		Format:  ProtoToMonitoringBetaDashboardWidgetTextFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToDashboardWidgetBlank converts a DashboardWidgetBlank object from its proto representation.
func ProtoToMonitoringBetaDashboardWidgetBlank(p *betapb.MonitoringBetaDashboardWidgetBlank) *beta.DashboardWidgetBlank {
	if p == nil {
		return nil
	}
	obj := &beta.DashboardWidgetBlank{}
	return obj
}

// ProtoToDashboard converts a Dashboard resource from its proto representation.
func ProtoToDashboard(p *betapb.MonitoringBetaDashboard) *beta.Dashboard {
	obj := &beta.Dashboard{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		GridLayout:   ProtoToMonitoringBetaDashboardGridLayout(p.GetGridLayout()),
		MosaicLayout: ProtoToMonitoringBetaDashboardMosaicLayout(p.GetMosaicLayout()),
		RowLayout:    ProtoToMonitoringBetaDashboardRowLayout(p.GetRowLayout()),
		ColumnLayout: ProtoToMonitoringBetaDashboardColumnLayout(p.GetColumnLayout()),
		Project:      dcl.StringOrNil(p.GetProject()),
		Etag:         dcl.StringOrNil(p.GetEtag()),
	}
	return obj
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsPlotTypeEnumToProto converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(e *beta.DashboardWidgetXyChartDataSetsPlotTypeEnum) betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum_value["DashboardWidgetXyChartDataSetsPlotTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
}

// DashboardWidgetXyChartThresholdsColorEnumToProto converts a DashboardWidgetXyChartThresholdsColorEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartThresholdsColorEnumToProto(e *beta.DashboardWidgetXyChartThresholdsColorEnum) betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum_value["DashboardWidgetXyChartThresholdsColorEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum(0)
}

// DashboardWidgetXyChartThresholdsDirectionEnumToProto converts a DashboardWidgetXyChartThresholdsDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnumToProto(e *beta.DashboardWidgetXyChartThresholdsDirectionEnum) betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum_value["DashboardWidgetXyChartThresholdsDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum(0)
}

// DashboardWidgetXyChartXAxisScaleEnumToProto converts a DashboardWidgetXyChartXAxisScaleEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartXAxisScaleEnumToProto(e *beta.DashboardWidgetXyChartXAxisScaleEnum) betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum_value["DashboardWidgetXyChartXAxisScaleEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum(0)
}

// DashboardWidgetXyChartYAxisScaleEnumToProto converts a DashboardWidgetXyChartYAxisScaleEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartYAxisScaleEnumToProto(e *beta.DashboardWidgetXyChartYAxisScaleEnum) betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum_value["DashboardWidgetXyChartYAxisScaleEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum(0)
}

// DashboardWidgetXyChartChartOptionsModeEnumToProto converts a DashboardWidgetXyChartChartOptionsModeEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnumToProto(e *beta.DashboardWidgetXyChartChartOptionsModeEnum) betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum_value["DashboardWidgetXyChartChartOptionsModeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(e *beta.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_value["DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
}

// DashboardWidgetScorecardThresholdsColorEnumToProto converts a DashboardWidgetScorecardThresholdsColorEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardThresholdsColorEnumToProto(e *beta.DashboardWidgetScorecardThresholdsColorEnum) betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum_value["DashboardWidgetScorecardThresholdsColorEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum(0)
}

// DashboardWidgetScorecardThresholdsDirectionEnumToProto converts a DashboardWidgetScorecardThresholdsDirectionEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnumToProto(e *beta.DashboardWidgetScorecardThresholdsDirectionEnum) betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum_value["DashboardWidgetScorecardThresholdsDirectionEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum(0)
}

// DashboardWidgetTextFormatEnumToProto converts a DashboardWidgetTextFormatEnum enum to its proto representation.
func MonitoringBetaDashboardWidgetTextFormatEnumToProto(e *beta.DashboardWidgetTextFormatEnum) betapb.MonitoringBetaDashboardWidgetTextFormatEnum {
	if e == nil {
		return betapb.MonitoringBetaDashboardWidgetTextFormatEnum(0)
	}
	if v, ok := betapb.MonitoringBetaDashboardWidgetTextFormatEnum_value["DashboardWidgetTextFormatEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaDashboardWidgetTextFormatEnum(v)
	}
	return betapb.MonitoringBetaDashboardWidgetTextFormatEnum(0)
}

// DashboardGridLayoutToProto converts a DashboardGridLayout object to its proto representation.
func MonitoringBetaDashboardGridLayoutToProto(o *beta.DashboardGridLayout) *betapb.MonitoringBetaDashboardGridLayout {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardGridLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sWidgets := make([]*betapb.MonitoringBetaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringBetaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardMosaicLayoutToProto converts a DashboardMosaicLayout object to its proto representation.
func MonitoringBetaDashboardMosaicLayoutToProto(o *beta.DashboardMosaicLayout) *betapb.MonitoringBetaDashboardMosaicLayout {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardMosaicLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sTiles := make([]*betapb.MonitoringBetaDashboardMosaicLayoutTiles, len(o.Tiles))
	for i, r := range o.Tiles {
		sTiles[i] = MonitoringBetaDashboardMosaicLayoutTilesToProto(&r)
	}
	p.SetTiles(sTiles)
	return p
}

// DashboardMosaicLayoutTilesToProto converts a DashboardMosaicLayoutTiles object to its proto representation.
func MonitoringBetaDashboardMosaicLayoutTilesToProto(o *beta.DashboardMosaicLayoutTiles) *betapb.MonitoringBetaDashboardMosaicLayoutTiles {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardMosaicLayoutTiles{}
	p.SetXPos(dcl.ValueOrEmptyInt64(o.XPos))
	p.SetYPos(dcl.ValueOrEmptyInt64(o.YPos))
	p.SetWidth(dcl.ValueOrEmptyInt64(o.Width))
	p.SetHeight(dcl.ValueOrEmptyInt64(o.Height))
	p.SetWidget(MonitoringBetaDashboardWidgetToProto(o.Widget))
	return p
}

// DashboardRowLayoutToProto converts a DashboardRowLayout object to its proto representation.
func MonitoringBetaDashboardRowLayoutToProto(o *beta.DashboardRowLayout) *betapb.MonitoringBetaDashboardRowLayout {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardRowLayout{}
	sRows := make([]*betapb.MonitoringBetaDashboardRowLayoutRows, len(o.Rows))
	for i, r := range o.Rows {
		sRows[i] = MonitoringBetaDashboardRowLayoutRowsToProto(&r)
	}
	p.SetRows(sRows)
	return p
}

// DashboardRowLayoutRowsToProto converts a DashboardRowLayoutRows object to its proto representation.
func MonitoringBetaDashboardRowLayoutRowsToProto(o *beta.DashboardRowLayoutRows) *betapb.MonitoringBetaDashboardRowLayoutRows {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardRowLayoutRows{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*betapb.MonitoringBetaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringBetaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardColumnLayoutToProto converts a DashboardColumnLayout object to its proto representation.
func MonitoringBetaDashboardColumnLayoutToProto(o *beta.DashboardColumnLayout) *betapb.MonitoringBetaDashboardColumnLayout {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardColumnLayout{}
	sColumns := make([]*betapb.MonitoringBetaDashboardColumnLayoutColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = MonitoringBetaDashboardColumnLayoutColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// DashboardColumnLayoutColumnsToProto converts a DashboardColumnLayoutColumns object to its proto representation.
func MonitoringBetaDashboardColumnLayoutColumnsToProto(o *beta.DashboardColumnLayoutColumns) *betapb.MonitoringBetaDashboardColumnLayoutColumns {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardColumnLayoutColumns{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*betapb.MonitoringBetaDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringBetaDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardWidgetToProto converts a DashboardWidget object to its proto representation.
func MonitoringBetaDashboardWidgetToProto(o *beta.DashboardWidget) *betapb.MonitoringBetaDashboardWidget {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidget{}
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetXyChart(MonitoringBetaDashboardWidgetXyChartToProto(o.XyChart))
	p.SetScorecard(MonitoringBetaDashboardWidgetScorecardToProto(o.Scorecard))
	p.SetText(MonitoringBetaDashboardWidgetTextToProto(o.Text))
	p.SetBlank(MonitoringBetaDashboardWidgetBlankToProto(o.Blank))
	return p
}

// DashboardWidgetXyChartToProto converts a DashboardWidgetXyChart object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartToProto(o *beta.DashboardWidgetXyChart) *betapb.MonitoringBetaDashboardWidgetXyChart {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChart{}
	p.SetTimeshiftDuration(dcl.ValueOrEmptyString(o.TimeshiftDuration))
	p.SetXAxis(MonitoringBetaDashboardWidgetXyChartXAxisToProto(o.XAxis))
	p.SetYAxis(MonitoringBetaDashboardWidgetXyChartYAxisToProto(o.YAxis))
	p.SetChartOptions(MonitoringBetaDashboardWidgetXyChartChartOptionsToProto(o.ChartOptions))
	sDataSets := make([]*betapb.MonitoringBetaDashboardWidgetXyChartDataSets, len(o.DataSets))
	for i, r := range o.DataSets {
		sDataSets[i] = MonitoringBetaDashboardWidgetXyChartDataSetsToProto(&r)
	}
	p.SetDataSets(sDataSets)
	sThresholds := make([]*betapb.MonitoringBetaDashboardWidgetXyChartThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringBetaDashboardWidgetXyChartThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetXyChartDataSetsToProto converts a DashboardWidgetXyChartDataSets object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsToProto(o *beta.DashboardWidgetXyChartDataSets) *betapb.MonitoringBetaDashboardWidgetXyChartDataSets {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSets{}
	p.SetTimeSeriesQuery(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetPlotType(MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(o.PlotType))
	p.SetLegendTemplate(dcl.ValueOrEmptyString(o.LegendTemplate))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQuery) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *beta.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartThresholdsToProto converts a DashboardWidgetXyChartThresholds object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartThresholdsToProto(o *beta.DashboardWidgetXyChartThresholds) *betapb.MonitoringBetaDashboardWidgetXyChartThresholds {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringBetaDashboardWidgetXyChartThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartXAxisToProto converts a DashboardWidgetXyChartXAxis object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartXAxisToProto(o *beta.DashboardWidgetXyChartXAxis) *betapb.MonitoringBetaDashboardWidgetXyChartXAxis {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartXAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringBetaDashboardWidgetXyChartXAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartYAxisToProto converts a DashboardWidgetXyChartYAxis object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartYAxisToProto(o *beta.DashboardWidgetXyChartYAxis) *betapb.MonitoringBetaDashboardWidgetXyChartYAxis {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartYAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringBetaDashboardWidgetXyChartYAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartChartOptionsToProto converts a DashboardWidgetXyChartChartOptions object to its proto representation.
func MonitoringBetaDashboardWidgetXyChartChartOptionsToProto(o *beta.DashboardWidgetXyChartChartOptions) *betapb.MonitoringBetaDashboardWidgetXyChartChartOptions {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetXyChartChartOptions{}
	p.SetMode(MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnumToProto(o.Mode))
	return p
}

// DashboardWidgetScorecardToProto converts a DashboardWidgetScorecard object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardToProto(o *beta.DashboardWidgetScorecard) *betapb.MonitoringBetaDashboardWidgetScorecard {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecard{}
	p.SetTimeSeriesQuery(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetGaugeView(MonitoringBetaDashboardWidgetScorecardGaugeViewToProto(o.GaugeView))
	p.SetSparkChartView(MonitoringBetaDashboardWidgetScorecardSparkChartViewToProto(o.SparkChartView))
	sThresholds := make([]*betapb.MonitoringBetaDashboardWidgetScorecardThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringBetaDashboardWidgetScorecardThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryToProto converts a DashboardWidgetScorecardTimeSeriesQuery object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryToProto(o *beta.DashboardWidgetScorecardTimeSeriesQuery) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *beta.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardGaugeViewToProto converts a DashboardWidgetScorecardGaugeView object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardGaugeViewToProto(o *beta.DashboardWidgetScorecardGaugeView) *betapb.MonitoringBetaDashboardWidgetScorecardGaugeView {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardGaugeView{}
	p.SetLowerBound(dcl.ValueOrEmptyDouble(o.LowerBound))
	p.SetUpperBound(dcl.ValueOrEmptyDouble(o.UpperBound))
	return p
}

// DashboardWidgetScorecardSparkChartViewToProto converts a DashboardWidgetScorecardSparkChartView object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardSparkChartViewToProto(o *beta.DashboardWidgetScorecardSparkChartView) *betapb.MonitoringBetaDashboardWidgetScorecardSparkChartView {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardSparkChartView{}
	p.SetSparkChartType(MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(o.SparkChartType))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetScorecardThresholdsToProto converts a DashboardWidgetScorecardThresholds object to its proto representation.
func MonitoringBetaDashboardWidgetScorecardThresholdsToProto(o *beta.DashboardWidgetScorecardThresholds) *betapb.MonitoringBetaDashboardWidgetScorecardThresholds {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetScorecardThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringBetaDashboardWidgetScorecardThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetTextToProto converts a DashboardWidgetText object to its proto representation.
func MonitoringBetaDashboardWidgetTextToProto(o *beta.DashboardWidgetText) *betapb.MonitoringBetaDashboardWidgetText {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetText{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetFormat(MonitoringBetaDashboardWidgetTextFormatEnumToProto(o.Format))
	return p
}

// DashboardWidgetBlankToProto converts a DashboardWidgetBlank object to its proto representation.
func MonitoringBetaDashboardWidgetBlankToProto(o *beta.DashboardWidgetBlank) *betapb.MonitoringBetaDashboardWidgetBlank {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaDashboardWidgetBlank{}
	return p
}

// DashboardToProto converts a Dashboard resource to its proto representation.
func DashboardToProto(resource *beta.Dashboard) *betapb.MonitoringBetaDashboard {
	p := &betapb.MonitoringBetaDashboard{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetGridLayout(MonitoringBetaDashboardGridLayoutToProto(resource.GridLayout))
	p.SetMosaicLayout(MonitoringBetaDashboardMosaicLayoutToProto(resource.MosaicLayout))
	p.SetRowLayout(MonitoringBetaDashboardRowLayoutToProto(resource.RowLayout))
	p.SetColumnLayout(MonitoringBetaDashboardColumnLayoutToProto(resource.ColumnLayout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))

	return p
}

// applyDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) applyDashboard(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaDashboardRequest) (*betapb.MonitoringBetaDashboard, error) {
	p := ProtoToDashboard(request.GetResource())
	res, err := c.ApplyDashboard(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DashboardToProto(res)
	return r, nil
}

// applyMonitoringBetaDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) ApplyMonitoringBetaDashboard(ctx context.Context, request *betapb.ApplyMonitoringBetaDashboardRequest) (*betapb.MonitoringBetaDashboard, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDashboard(ctx, cl, request)
}

// DeleteDashboard handles the gRPC request by passing it to the underlying Dashboard Delete() method.
func (s *DashboardServer) DeleteMonitoringBetaDashboard(ctx context.Context, request *betapb.DeleteMonitoringBetaDashboardRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDashboard(ctx, ProtoToDashboard(request.GetResource()))

}

// ListMonitoringBetaDashboard handles the gRPC request by passing it to the underlying DashboardList() method.
func (s *DashboardServer) ListMonitoringBetaDashboard(ctx context.Context, request *betapb.ListMonitoringBetaDashboardRequest) (*betapb.ListMonitoringBetaDashboardResponse, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDashboard(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaDashboard
	for _, r := range resources.Items {
		rp := DashboardToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaDashboardResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDashboard(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
