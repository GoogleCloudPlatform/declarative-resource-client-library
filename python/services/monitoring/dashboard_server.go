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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// Server implements the gRPC interface for Dashboard.
type DashboardServer struct{}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsPlotTypeEnum converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum) *monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum(n[len("MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsColorEnum converts a DashboardWidgetXyChartThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholdsColorEnum(e monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum) *monitoring.DashboardWidgetXyChartThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartThresholdsColorEnum(n[len("MonitoringDashboardWidgetXyChartThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsDirectionEnum converts a DashboardWidgetXyChartThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholdsDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum) *monitoring.DashboardWidgetXyChartThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartThresholdsDirectionEnum(n[len("MonitoringDashboardWidgetXyChartThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartXAxisScaleEnum converts a DashboardWidgetXyChartXAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartXAxisScaleEnum(e monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum) *monitoring.DashboardWidgetXyChartXAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartXAxisScaleEnum(n[len("MonitoringDashboardWidgetXyChartXAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartYAxisScaleEnum converts a DashboardWidgetXyChartYAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartYAxisScaleEnum(e monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum) *monitoring.DashboardWidgetXyChartYAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartYAxisScaleEnum(n[len("MonitoringDashboardWidgetXyChartYAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartChartOptionsModeEnum converts a DashboardWidgetXyChartChartOptionsModeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartChartOptionsModeEnum(e monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum) *monitoring.DashboardWidgetXyChartChartOptionsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartChartOptionsModeEnum(n[len("MonitoringDashboardWidgetXyChartChartOptionsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(e monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) *monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(n[len("MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsColorEnum converts a DashboardWidgetScorecardThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholdsColorEnum(e monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum) *monitoring.DashboardWidgetScorecardThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardThresholdsColorEnum(n[len("MonitoringDashboardWidgetScorecardThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsDirectionEnum converts a DashboardWidgetScorecardThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholdsDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum) *monitoring.DashboardWidgetScorecardThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardThresholdsDirectionEnum(n[len("MonitoringDashboardWidgetScorecardThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetTextFormatEnum converts a DashboardWidgetTextFormatEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetTextFormatEnum(e monitoringpb.MonitoringDashboardWidgetTextFormatEnum) *monitoring.DashboardWidgetTextFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetTextFormatEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetTextFormatEnum(n[len("MonitoringDashboardWidgetTextFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardGridLayout converts a DashboardGridLayout object from its proto representation.
func ProtoToMonitoringDashboardGridLayout(p *monitoringpb.MonitoringDashboardGridLayout) *monitoring.DashboardGridLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardGridLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayout converts a DashboardMosaicLayout object from its proto representation.
func ProtoToMonitoringDashboardMosaicLayout(p *monitoringpb.MonitoringDashboardMosaicLayout) *monitoring.DashboardMosaicLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardMosaicLayout{
		Columns: dcl.Int64OrNil(p.GetColumns()),
	}
	for _, r := range p.GetTiles() {
		obj.Tiles = append(obj.Tiles, *ProtoToMonitoringDashboardMosaicLayoutTiles(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayoutTiles converts a DashboardMosaicLayoutTiles object from its proto representation.
func ProtoToMonitoringDashboardMosaicLayoutTiles(p *monitoringpb.MonitoringDashboardMosaicLayoutTiles) *monitoring.DashboardMosaicLayoutTiles {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardMosaicLayoutTiles{
		XPos:   dcl.Int64OrNil(p.GetXPos()),
		YPos:   dcl.Int64OrNil(p.GetYPos()),
		Width:  dcl.Int64OrNil(p.GetWidth()),
		Height: dcl.Int64OrNil(p.GetHeight()),
		Widget: ProtoToMonitoringDashboardWidget(p.GetWidget()),
	}
	return obj
}

// ProtoToDashboardRowLayout converts a DashboardRowLayout object from its proto representation.
func ProtoToMonitoringDashboardRowLayout(p *monitoringpb.MonitoringDashboardRowLayout) *monitoring.DashboardRowLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardRowLayout{}
	for _, r := range p.GetRows() {
		obj.Rows = append(obj.Rows, *ProtoToMonitoringDashboardRowLayoutRows(r))
	}
	return obj
}

// ProtoToDashboardRowLayoutRows converts a DashboardRowLayoutRows object from its proto representation.
func ProtoToMonitoringDashboardRowLayoutRows(p *monitoringpb.MonitoringDashboardRowLayoutRows) *monitoring.DashboardRowLayoutRows {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardRowLayoutRows{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardColumnLayout converts a DashboardColumnLayout object from its proto representation.
func ProtoToMonitoringDashboardColumnLayout(p *monitoringpb.MonitoringDashboardColumnLayout) *monitoring.DashboardColumnLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardColumnLayout{}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToMonitoringDashboardColumnLayoutColumns(r))
	}
	return obj
}

// ProtoToDashboardColumnLayoutColumns converts a DashboardColumnLayoutColumns object from its proto representation.
func ProtoToMonitoringDashboardColumnLayoutColumns(p *monitoringpb.MonitoringDashboardColumnLayoutColumns) *monitoring.DashboardColumnLayoutColumns {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardColumnLayoutColumns{
		Weight: dcl.Int64OrNil(p.GetWeight()),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardWidget converts a DashboardWidget object from its proto representation.
func ProtoToMonitoringDashboardWidget(p *monitoringpb.MonitoringDashboardWidget) *monitoring.DashboardWidget {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidget{
		Title:     dcl.StringOrNil(p.GetTitle()),
		XyChart:   ProtoToMonitoringDashboardWidgetXyChart(p.GetXyChart()),
		Scorecard: ProtoToMonitoringDashboardWidgetScorecard(p.GetScorecard()),
		Text:      ProtoToMonitoringDashboardWidgetText(p.GetText()),
		Blank:     ProtoToMonitoringDashboardWidgetBlank(p.GetBlank()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChart converts a DashboardWidgetXyChart object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChart(p *monitoringpb.MonitoringDashboardWidgetXyChart) *monitoring.DashboardWidgetXyChart {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChart{
		TimeshiftDuration: dcl.StringOrNil(p.GetTimeshiftDuration()),
		XAxis:             ProtoToMonitoringDashboardWidgetXyChartXAxis(p.GetXAxis()),
		YAxis:             ProtoToMonitoringDashboardWidgetXyChartYAxis(p.GetYAxis()),
		ChartOptions:      ProtoToMonitoringDashboardWidgetXyChartChartOptions(p.GetChartOptions()),
	}
	for _, r := range p.GetDataSets() {
		obj.DataSets = append(obj.DataSets, *ProtoToMonitoringDashboardWidgetXyChartDataSets(r))
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringDashboardWidgetXyChartThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSets converts a DashboardWidgetXyChartDataSets object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSets) *monitoring.DashboardWidgetXyChartDataSets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSets{
		TimeSeriesQuery:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery(p.GetTimeSeriesQuery()),
		PlotType:           ProtoToMonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(p.GetPlotType()),
		LegendTemplate:     dcl.StringOrNil(p.GetLegendTemplate()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQuery converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartThresholds converts a DashboardWidgetXyChartThresholds object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholds(p *monitoringpb.MonitoringDashboardWidgetXyChartThresholds) *monitoring.DashboardWidgetXyChartThresholds {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringDashboardWidgetXyChartThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringDashboardWidgetXyChartThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartXAxis converts a DashboardWidgetXyChartXAxis object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartXAxis(p *monitoringpb.MonitoringDashboardWidgetXyChartXAxis) *monitoring.DashboardWidgetXyChartXAxis {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartXAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringDashboardWidgetXyChartXAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartYAxis converts a DashboardWidgetXyChartYAxis object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartYAxis(p *monitoringpb.MonitoringDashboardWidgetXyChartYAxis) *monitoring.DashboardWidgetXyChartYAxis {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartYAxis{
		Label: dcl.StringOrNil(p.GetLabel()),
		Scale: ProtoToMonitoringDashboardWidgetXyChartYAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartChartOptions converts a DashboardWidgetXyChartChartOptions object from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartChartOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartChartOptions) *monitoring.DashboardWidgetXyChartChartOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartChartOptions{
		Mode: ProtoToMonitoringDashboardWidgetXyChartChartOptionsModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecard converts a DashboardWidgetScorecard object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecard(p *monitoringpb.MonitoringDashboardWidgetScorecard) *monitoring.DashboardWidgetScorecard {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecard{
		TimeSeriesQuery: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQuery(p.GetTimeSeriesQuery()),
		GaugeView:       ProtoToMonitoringDashboardWidgetScorecardGaugeView(p.GetGaugeView()),
		SparkChartView:  ProtoToMonitoringDashboardWidgetScorecardSparkChartView(p.GetSparkChartView()),
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringDashboardWidgetScorecardThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQuery converts a DashboardWidgetScorecardTimeSeriesQuery object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQuery(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery) *monitoring.DashboardWidgetScorecardTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.GetTimeSeriesQueryLanguage()),
		UnitOverride:            dcl.StringOrNil(p.GetUnitOverride()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.GetFilter()),
		Aggregation:          ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Aggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:    dcl.StringOrNil(p.GetAlignmentPeriod()),
		PerSeriesAligner:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.GetNumTimeSeries()),
		Direction:     ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardGaugeView converts a DashboardWidgetScorecardGaugeView object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardGaugeView(p *monitoringpb.MonitoringDashboardWidgetScorecardGaugeView) *monitoring.DashboardWidgetScorecardGaugeView {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardGaugeView{
		LowerBound: dcl.Float64OrNil(p.GetLowerBound()),
		UpperBound: dcl.Float64OrNil(p.GetUpperBound()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardSparkChartView converts a DashboardWidgetScorecardSparkChartView object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardSparkChartView(p *monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView) *monitoring.DashboardWidgetScorecardSparkChartView {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardSparkChartView{
		SparkChartType:     ProtoToMonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(p.GetSparkChartType()),
		MinAlignmentPeriod: dcl.StringOrNil(p.GetMinAlignmentPeriod()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardThresholds converts a DashboardWidgetScorecardThresholds object from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholds(p *monitoringpb.MonitoringDashboardWidgetScorecardThresholds) *monitoring.DashboardWidgetScorecardThresholds {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardThresholds{
		Label:     dcl.StringOrNil(p.GetLabel()),
		Value:     dcl.Float64OrNil(p.GetValue()),
		Color:     ProtoToMonitoringDashboardWidgetScorecardThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringDashboardWidgetScorecardThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetText converts a DashboardWidgetText object from its proto representation.
func ProtoToMonitoringDashboardWidgetText(p *monitoringpb.MonitoringDashboardWidgetText) *monitoring.DashboardWidgetText {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetText{
		Content: dcl.StringOrNil(p.GetContent()),
		Format:  ProtoToMonitoringDashboardWidgetTextFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToDashboardWidgetBlank converts a DashboardWidgetBlank object from its proto representation.
func ProtoToMonitoringDashboardWidgetBlank(p *monitoringpb.MonitoringDashboardWidgetBlank) *monitoring.DashboardWidgetBlank {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetBlank{}
	return obj
}

// ProtoToDashboard converts a Dashboard resource from its proto representation.
func ProtoToDashboard(p *monitoringpb.MonitoringDashboard) *monitoring.Dashboard {
	obj := &monitoring.Dashboard{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		GridLayout:   ProtoToMonitoringDashboardGridLayout(p.GetGridLayout()),
		MosaicLayout: ProtoToMonitoringDashboardMosaicLayout(p.GetMosaicLayout()),
		RowLayout:    ProtoToMonitoringDashboardRowLayout(p.GetRowLayout()),
		ColumnLayout: ProtoToMonitoringDashboardColumnLayout(p.GetColumnLayout()),
		Project:      dcl.StringOrNil(p.GetProject()),
		Etag:         dcl.StringOrNil(p.GetEtag()),
	}
	return obj
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsPlotTypeEnumToProto converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum_value["DashboardWidgetXyChartDataSetsPlotTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
}

// DashboardWidgetXyChartThresholdsColorEnumToProto converts a DashboardWidgetXyChartThresholdsColorEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsColorEnumToProto(e *monitoring.DashboardWidgetXyChartThresholdsColorEnum) monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum_value["DashboardWidgetXyChartThresholdsColorEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(0)
}

// DashboardWidgetXyChartThresholdsDirectionEnumToProto converts a DashboardWidgetXyChartThresholdsDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartThresholdsDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum_value["DashboardWidgetXyChartThresholdsDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(0)
}

// DashboardWidgetXyChartXAxisScaleEnumToProto converts a DashboardWidgetXyChartXAxisScaleEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartXAxisScaleEnumToProto(e *monitoring.DashboardWidgetXyChartXAxisScaleEnum) monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum_value["DashboardWidgetXyChartXAxisScaleEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(0)
}

// DashboardWidgetXyChartYAxisScaleEnumToProto converts a DashboardWidgetXyChartYAxisScaleEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartYAxisScaleEnumToProto(e *monitoring.DashboardWidgetXyChartYAxisScaleEnum) monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum_value["DashboardWidgetXyChartYAxisScaleEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(0)
}

// DashboardWidgetXyChartChartOptionsModeEnumToProto converts a DashboardWidgetXyChartChartOptionsModeEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartChartOptionsModeEnumToProto(e *monitoring.DashboardWidgetXyChartChartOptionsModeEnum) monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum_value["DashboardWidgetXyChartChartOptionsModeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(e *monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_value["DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
}

// DashboardWidgetScorecardThresholdsColorEnumToProto converts a DashboardWidgetScorecardThresholdsColorEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsColorEnumToProto(e *monitoring.DashboardWidgetScorecardThresholdsColorEnum) monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum_value["DashboardWidgetScorecardThresholdsColorEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(0)
}

// DashboardWidgetScorecardThresholdsDirectionEnumToProto converts a DashboardWidgetScorecardThresholdsDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardThresholdsDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum_value["DashboardWidgetScorecardThresholdsDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(0)
}

// DashboardWidgetTextFormatEnumToProto converts a DashboardWidgetTextFormatEnum enum to its proto representation.
func MonitoringDashboardWidgetTextFormatEnumToProto(e *monitoring.DashboardWidgetTextFormatEnum) monitoringpb.MonitoringDashboardWidgetTextFormatEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetTextFormatEnum_value["DashboardWidgetTextFormatEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(0)
}

// DashboardGridLayoutToProto converts a DashboardGridLayout object to its proto representation.
func MonitoringDashboardGridLayoutToProto(o *monitoring.DashboardGridLayout) *monitoringpb.MonitoringDashboardGridLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardGridLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sWidgets := make([]*monitoringpb.MonitoringDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardMosaicLayoutToProto converts a DashboardMosaicLayout object to its proto representation.
func MonitoringDashboardMosaicLayoutToProto(o *monitoring.DashboardMosaicLayout) *monitoringpb.MonitoringDashboardMosaicLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardMosaicLayout{}
	p.SetColumns(dcl.ValueOrEmptyInt64(o.Columns))
	sTiles := make([]*monitoringpb.MonitoringDashboardMosaicLayoutTiles, len(o.Tiles))
	for i, r := range o.Tiles {
		sTiles[i] = MonitoringDashboardMosaicLayoutTilesToProto(&r)
	}
	p.SetTiles(sTiles)
	return p
}

// DashboardMosaicLayoutTilesToProto converts a DashboardMosaicLayoutTiles object to its proto representation.
func MonitoringDashboardMosaicLayoutTilesToProto(o *monitoring.DashboardMosaicLayoutTiles) *monitoringpb.MonitoringDashboardMosaicLayoutTiles {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardMosaicLayoutTiles{}
	p.SetXPos(dcl.ValueOrEmptyInt64(o.XPos))
	p.SetYPos(dcl.ValueOrEmptyInt64(o.YPos))
	p.SetWidth(dcl.ValueOrEmptyInt64(o.Width))
	p.SetHeight(dcl.ValueOrEmptyInt64(o.Height))
	p.SetWidget(MonitoringDashboardWidgetToProto(o.Widget))
	return p
}

// DashboardRowLayoutToProto converts a DashboardRowLayout object to its proto representation.
func MonitoringDashboardRowLayoutToProto(o *monitoring.DashboardRowLayout) *monitoringpb.MonitoringDashboardRowLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardRowLayout{}
	sRows := make([]*monitoringpb.MonitoringDashboardRowLayoutRows, len(o.Rows))
	for i, r := range o.Rows {
		sRows[i] = MonitoringDashboardRowLayoutRowsToProto(&r)
	}
	p.SetRows(sRows)
	return p
}

// DashboardRowLayoutRowsToProto converts a DashboardRowLayoutRows object to its proto representation.
func MonitoringDashboardRowLayoutRowsToProto(o *monitoring.DashboardRowLayoutRows) *monitoringpb.MonitoringDashboardRowLayoutRows {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardRowLayoutRows{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*monitoringpb.MonitoringDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardColumnLayoutToProto converts a DashboardColumnLayout object to its proto representation.
func MonitoringDashboardColumnLayoutToProto(o *monitoring.DashboardColumnLayout) *monitoringpb.MonitoringDashboardColumnLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardColumnLayout{}
	sColumns := make([]*monitoringpb.MonitoringDashboardColumnLayoutColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = MonitoringDashboardColumnLayoutColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// DashboardColumnLayoutColumnsToProto converts a DashboardColumnLayoutColumns object to its proto representation.
func MonitoringDashboardColumnLayoutColumnsToProto(o *monitoring.DashboardColumnLayoutColumns) *monitoringpb.MonitoringDashboardColumnLayoutColumns {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardColumnLayoutColumns{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	sWidgets := make([]*monitoringpb.MonitoringDashboardWidget, len(o.Widgets))
	for i, r := range o.Widgets {
		sWidgets[i] = MonitoringDashboardWidgetToProto(&r)
	}
	p.SetWidgets(sWidgets)
	return p
}

// DashboardWidgetToProto converts a DashboardWidget object to its proto representation.
func MonitoringDashboardWidgetToProto(o *monitoring.DashboardWidget) *monitoringpb.MonitoringDashboardWidget {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidget{}
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetXyChart(MonitoringDashboardWidgetXyChartToProto(o.XyChart))
	p.SetScorecard(MonitoringDashboardWidgetScorecardToProto(o.Scorecard))
	p.SetText(MonitoringDashboardWidgetTextToProto(o.Text))
	p.SetBlank(MonitoringDashboardWidgetBlankToProto(o.Blank))
	return p
}

// DashboardWidgetXyChartToProto converts a DashboardWidgetXyChart object to its proto representation.
func MonitoringDashboardWidgetXyChartToProto(o *monitoring.DashboardWidgetXyChart) *monitoringpb.MonitoringDashboardWidgetXyChart {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChart{}
	p.SetTimeshiftDuration(dcl.ValueOrEmptyString(o.TimeshiftDuration))
	p.SetXAxis(MonitoringDashboardWidgetXyChartXAxisToProto(o.XAxis))
	p.SetYAxis(MonitoringDashboardWidgetXyChartYAxisToProto(o.YAxis))
	p.SetChartOptions(MonitoringDashboardWidgetXyChartChartOptionsToProto(o.ChartOptions))
	sDataSets := make([]*monitoringpb.MonitoringDashboardWidgetXyChartDataSets, len(o.DataSets))
	for i, r := range o.DataSets {
		sDataSets[i] = MonitoringDashboardWidgetXyChartDataSetsToProto(&r)
	}
	p.SetDataSets(sDataSets)
	sThresholds := make([]*monitoringpb.MonitoringDashboardWidgetXyChartThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringDashboardWidgetXyChartThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetXyChartDataSetsToProto converts a DashboardWidgetXyChartDataSets object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsToProto(o *monitoring.DashboardWidgetXyChartDataSets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSets{}
	p.SetTimeSeriesQuery(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetPlotType(MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(o.PlotType))
	p.SetLegendTemplate(dcl.ValueOrEmptyString(o.LegendTemplate))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartThresholdsToProto converts a DashboardWidgetXyChartThresholds object to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsToProto(o *monitoring.DashboardWidgetXyChartThresholds) *monitoringpb.MonitoringDashboardWidgetXyChartThresholds {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringDashboardWidgetXyChartThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringDashboardWidgetXyChartThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetXyChartXAxisToProto converts a DashboardWidgetXyChartXAxis object to its proto representation.
func MonitoringDashboardWidgetXyChartXAxisToProto(o *monitoring.DashboardWidgetXyChartXAxis) *monitoringpb.MonitoringDashboardWidgetXyChartXAxis {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartXAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringDashboardWidgetXyChartXAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartYAxisToProto converts a DashboardWidgetXyChartYAxis object to its proto representation.
func MonitoringDashboardWidgetXyChartYAxisToProto(o *monitoring.DashboardWidgetXyChartYAxis) *monitoringpb.MonitoringDashboardWidgetXyChartYAxis {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartYAxis{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetScale(MonitoringDashboardWidgetXyChartYAxisScaleEnumToProto(o.Scale))
	return p
}

// DashboardWidgetXyChartChartOptionsToProto converts a DashboardWidgetXyChartChartOptions object to its proto representation.
func MonitoringDashboardWidgetXyChartChartOptionsToProto(o *monitoring.DashboardWidgetXyChartChartOptions) *monitoringpb.MonitoringDashboardWidgetXyChartChartOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartChartOptions{}
	p.SetMode(MonitoringDashboardWidgetXyChartChartOptionsModeEnumToProto(o.Mode))
	return p
}

// DashboardWidgetScorecardToProto converts a DashboardWidgetScorecard object to its proto representation.
func MonitoringDashboardWidgetScorecardToProto(o *monitoring.DashboardWidgetScorecard) *monitoringpb.MonitoringDashboardWidgetScorecard {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecard{}
	p.SetTimeSeriesQuery(MonitoringDashboardWidgetScorecardTimeSeriesQueryToProto(o.TimeSeriesQuery))
	p.SetGaugeView(MonitoringDashboardWidgetScorecardGaugeViewToProto(o.GaugeView))
	p.SetSparkChartView(MonitoringDashboardWidgetScorecardSparkChartViewToProto(o.SparkChartView))
	sThresholds := make([]*monitoringpb.MonitoringDashboardWidgetScorecardThresholds, len(o.Thresholds))
	for i, r := range o.Thresholds {
		sThresholds[i] = MonitoringDashboardWidgetScorecardThresholdsToProto(&r)
	}
	p.SetThresholds(sThresholds)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryToProto converts a DashboardWidgetScorecardTimeSeriesQuery object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQuery) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery{}
	p.SetTimeSeriesFilter(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter))
	p.SetTimeSeriesFilterRatio(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio))
	p.SetTimeSeriesQueryLanguage(dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage))
	p.SetUnitOverride(dcl.ValueOrEmptyString(o.UnitOverride))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation))
	p.SetSecondaryAggregation(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	p.SetNumerator(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator))
	p.SetDenominator(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator))
	p.SetSecondaryAggregation(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation))
	p.SetPickTimeSeriesFilter(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	p.SetFilter(dcl.ValueOrEmptyString(o.Filter))
	p.SetAggregation(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation))
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	p.SetAlignmentPeriod(dcl.ValueOrEmptyString(o.AlignmentPeriod))
	p.SetPerSeriesAligner(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner))
	p.SetCrossSeriesReducer(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer))
	sGroupByFields := make([]string, len(o.GroupByFields))
	for i, r := range o.GroupByFields {
		sGroupByFields[i] = r
	}
	p.SetGroupByFields(sGroupByFields)
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter object to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	p.SetRankingMethod(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod))
	p.SetNumTimeSeries(dcl.ValueOrEmptyInt64(o.NumTimeSeries))
	p.SetDirection(MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetScorecardGaugeViewToProto converts a DashboardWidgetScorecardGaugeView object to its proto representation.
func MonitoringDashboardWidgetScorecardGaugeViewToProto(o *monitoring.DashboardWidgetScorecardGaugeView) *monitoringpb.MonitoringDashboardWidgetScorecardGaugeView {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardGaugeView{}
	p.SetLowerBound(dcl.ValueOrEmptyDouble(o.LowerBound))
	p.SetUpperBound(dcl.ValueOrEmptyDouble(o.UpperBound))
	return p
}

// DashboardWidgetScorecardSparkChartViewToProto converts a DashboardWidgetScorecardSparkChartView object to its proto representation.
func MonitoringDashboardWidgetScorecardSparkChartViewToProto(o *monitoring.DashboardWidgetScorecardSparkChartView) *monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView{}
	p.SetSparkChartType(MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(o.SparkChartType))
	p.SetMinAlignmentPeriod(dcl.ValueOrEmptyString(o.MinAlignmentPeriod))
	return p
}

// DashboardWidgetScorecardThresholdsToProto converts a DashboardWidgetScorecardThresholds object to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsToProto(o *monitoring.DashboardWidgetScorecardThresholds) *monitoringpb.MonitoringDashboardWidgetScorecardThresholds {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardThresholds{}
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	p.SetValue(dcl.ValueOrEmptyDouble(o.Value))
	p.SetColor(MonitoringDashboardWidgetScorecardThresholdsColorEnumToProto(o.Color))
	p.SetDirection(MonitoringDashboardWidgetScorecardThresholdsDirectionEnumToProto(o.Direction))
	return p
}

// DashboardWidgetTextToProto converts a DashboardWidgetText object to its proto representation.
func MonitoringDashboardWidgetTextToProto(o *monitoring.DashboardWidgetText) *monitoringpb.MonitoringDashboardWidgetText {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetText{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetFormat(MonitoringDashboardWidgetTextFormatEnumToProto(o.Format))
	return p
}

// DashboardWidgetBlankToProto converts a DashboardWidgetBlank object to its proto representation.
func MonitoringDashboardWidgetBlankToProto(o *monitoring.DashboardWidgetBlank) *monitoringpb.MonitoringDashboardWidgetBlank {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetBlank{}
	return p
}

// DashboardToProto converts a Dashboard resource to its proto representation.
func DashboardToProto(resource *monitoring.Dashboard) *monitoringpb.MonitoringDashboard {
	p := &monitoringpb.MonitoringDashboard{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetGridLayout(MonitoringDashboardGridLayoutToProto(resource.GridLayout))
	p.SetMosaicLayout(MonitoringDashboardMosaicLayoutToProto(resource.MosaicLayout))
	p.SetRowLayout(MonitoringDashboardRowLayoutToProto(resource.RowLayout))
	p.SetColumnLayout(MonitoringDashboardColumnLayoutToProto(resource.ColumnLayout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))

	return p
}

// applyDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) applyDashboard(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringDashboardRequest) (*monitoringpb.MonitoringDashboard, error) {
	p := ProtoToDashboard(request.GetResource())
	res, err := c.ApplyDashboard(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DashboardToProto(res)
	return r, nil
}

// applyMonitoringDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) ApplyMonitoringDashboard(ctx context.Context, request *monitoringpb.ApplyMonitoringDashboardRequest) (*monitoringpb.MonitoringDashboard, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDashboard(ctx, cl, request)
}

// DeleteDashboard handles the gRPC request by passing it to the underlying Dashboard Delete() method.
func (s *DashboardServer) DeleteMonitoringDashboard(ctx context.Context, request *monitoringpb.DeleteMonitoringDashboardRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDashboard(ctx, ProtoToDashboard(request.GetResource()))

}

// ListMonitoringDashboard handles the gRPC request by passing it to the underlying DashboardList() method.
func (s *DashboardServer) ListMonitoringDashboard(ctx context.Context, request *monitoringpb.ListMonitoringDashboardRequest) (*monitoringpb.ListMonitoringDashboardResponse, error) {
	cl, err := createConfigDashboard(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDashboard(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringDashboard
	for _, r := range resources.Items {
		rp := DashboardToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringDashboardResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDashboard(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
