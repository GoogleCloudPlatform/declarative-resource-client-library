# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.monitoring import dashboard_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import dashboard_pb2_grpc

from typing import List


class Dashboard(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        grid_layout: dict = None,
        mosaic_layout: dict = None,
        row_layout: dict = None,
        column_layout: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.grid_layout = grid_layout
        self.mosaic_layout = mosaic_layout
        self.row_layout = row_layout
        self.column_layout = column_layout
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ApplyMonitoringDashboardRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if DashboardGridLayout.to_proto(self.grid_layout):
            request.resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            request.resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            request.resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            request.resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            request.resource.row_layout.CopyFrom(
                DashboardRowLayout.to_proto(self.row_layout)
            )
        else:
            request.resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            request.resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            request.resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringDashboard(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.grid_layout = DashboardGridLayout.from_proto(response.grid_layout)
        self.mosaic_layout = DashboardMosaicLayout.from_proto(response.mosaic_layout)
        self.row_layout = DashboardRowLayout.from_proto(response.row_layout)
        self.column_layout = DashboardColumnLayout.from_proto(response.column_layout)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.DeleteMonitoringDashboardRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if DashboardGridLayout.to_proto(self.grid_layout):
            request.resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            request.resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            request.resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            request.resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            request.resource.row_layout.CopyFrom(
                DashboardRowLayout.to_proto(self.row_layout)
            )
        else:
            request.resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            request.resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            request.resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringDashboard(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ListMonitoringDashboardRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringDashboard(request).items

    def to_proto(self):
        resource = dashboard_pb2.MonitoringDashboard()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if DashboardGridLayout.to_proto(self.grid_layout):
            resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            resource.row_layout.CopyFrom(DashboardRowLayout.to_proto(self.row_layout))
        else:
            resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class DashboardGridLayout(object):
    def __init__(self, columns: int = None, widgets: list = None):
        self.columns = columns
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardGridLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardWidgetArray.to_proto(resource.widgets):
            res.widgets.extend(DashboardWidgetArray.to_proto(resource.widgets))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayout(columns=resource.columns, widgets=resource.widgets,)


class DashboardGridLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayout.from_proto(i) for i in resources]


class DashboardMosaicLayout(object):
    def __init__(self, columns: int = None, tiles: list = None):
        self.columns = columns
        self.tiles = tiles

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardMosaicLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardMosaicLayoutTilesArray.to_proto(resource.tiles):
            res.tiles.extend(DashboardMosaicLayoutTilesArray.to_proto(resource.tiles))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayout(columns=resource.columns, tiles=resource.tiles,)


class DashboardMosaicLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayout.from_proto(i) for i in resources]


class DashboardMosaicLayoutTiles(object):
    def __init__(
        self,
        x_pos: int = None,
        y_pos: int = None,
        width: int = None,
        height: int = None,
        widget: dict = None,
    ):
        self.x_pos = x_pos
        self.y_pos = y_pos
        self.width = width
        self.height = height
        self.widget = widget

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTiles()
        if Primitive.to_proto(resource.x_pos):
            res.x_pos = Primitive.to_proto(resource.x_pos)
        if Primitive.to_proto(resource.y_pos):
            res.y_pos = Primitive.to_proto(resource.y_pos)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.height):
            res.height = Primitive.to_proto(resource.height)
        if DashboardWidget.to_proto(resource.widget):
            res.widget.CopyFrom(DashboardWidget.to_proto(resource.widget))
        else:
            res.ClearField("widget")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTiles(
            x_pos=resource.x_pos,
            y_pos=resource.y_pos,
            width=resource.width,
            height=resource.height,
            widget=resource.widget,
        )


class DashboardMosaicLayoutTilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayoutTiles.from_proto(i) for i in resources]


class DashboardRowLayout(object):
    def __init__(self, rows: list = None):
        self.rows = rows

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardRowLayout()
        if DashboardRowLayoutRowsArray.to_proto(resource.rows):
            res.rows.extend(DashboardRowLayoutRowsArray.to_proto(resource.rows))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayout(rows=resource.rows,)


class DashboardRowLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayout.from_proto(i) for i in resources]


class DashboardRowLayoutRows(object):
    def __init__(self, weight: int = None, widgets: list = None):
        self.weight = weight
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardRowLayoutRows()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if DashboardWidgetArray.to_proto(resource.widgets):
            res.widgets.extend(DashboardWidgetArray.to_proto(resource.widgets))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRows(weight=resource.weight, widgets=resource.widgets,)


class DashboardRowLayoutRowsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRows.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRows.from_proto(i) for i in resources]


class DashboardColumnLayout(object):
    def __init__(self, columns: list = None):
        self.columns = columns

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardColumnLayout()
        if DashboardColumnLayoutColumnsArray.to_proto(resource.columns):
            res.columns.extend(
                DashboardColumnLayoutColumnsArray.to_proto(resource.columns)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayout(columns=resource.columns,)


class DashboardColumnLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardColumnLayout.from_proto(i) for i in resources]


class DashboardColumnLayoutColumns(object):
    def __init__(self, weight: int = None, widgets: list = None):
        self.weight = weight
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumns()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if DashboardWidgetArray.to_proto(resource.widgets):
            res.widgets.extend(DashboardWidgetArray.to_proto(resource.widgets))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumns(
            weight=resource.weight, widgets=resource.widgets,
        )


class DashboardColumnLayoutColumnsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayoutColumns.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardColumnLayoutColumns.from_proto(i) for i in resources]


class DashboardWidget(object):
    def __init__(
        self,
        title: str = None,
        xy_chart: dict = None,
        scorecard: dict = None,
        text: dict = None,
        blank: dict = None,
    ):
        self.title = title
        self.xy_chart = xy_chart
        self.scorecard = scorecard
        self.text = text
        self.blank = blank

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidget()
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if DashboardWidgetXyChart.to_proto(resource.xy_chart):
            res.xy_chart.CopyFrom(DashboardWidgetXyChart.to_proto(resource.xy_chart))
        else:
            res.ClearField("xy_chart")
        if DashboardWidgetScorecard.to_proto(resource.scorecard):
            res.scorecard.CopyFrom(
                DashboardWidgetScorecard.to_proto(resource.scorecard)
            )
        else:
            res.ClearField("scorecard")
        if DashboardWidgetText.to_proto(resource.text):
            res.text.CopyFrom(DashboardWidgetText.to_proto(resource.text))
        else:
            res.ClearField("text")
        if DashboardWidgetBlank.to_proto(resource.blank):
            res.blank.CopyFrom(DashboardWidgetBlank.to_proto(resource.blank))
        else:
            res.ClearField("blank")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidget(
            title=resource.title,
            xy_chart=resource.xy_chart,
            scorecard=resource.scorecard,
            text=resource.text,
            blank=resource.blank,
        )


class DashboardWidgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidget.from_proto(i) for i in resources]


class DashboardWidgetXyChart(object):
    def __init__(
        self,
        data_sets: list = None,
        source_drilldown: dict = None,
        metric_drilldown: dict = None,
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.source_drilldown = source_drilldown
        self.metric_drilldown = metric_drilldown
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChart()
        if DashboardWidgetXyChartDataSetsArray.to_proto(resource.data_sets):
            res.data_sets.extend(
                DashboardWidgetXyChartDataSetsArray.to_proto(resource.data_sets)
            )
        if DashboardWidgetXyChartSourceDrilldown.to_proto(resource.source_drilldown):
            res.source_drilldown.CopyFrom(
                DashboardWidgetXyChartSourceDrilldown.to_proto(
                    resource.source_drilldown
                )
            )
        else:
            res.ClearField("source_drilldown")
        if DashboardWidgetXyChartMetricDrilldown.to_proto(resource.metric_drilldown):
            res.metric_drilldown.CopyFrom(
                DashboardWidgetXyChartMetricDrilldown.to_proto(
                    resource.metric_drilldown
                )
            )
        else:
            res.ClearField("metric_drilldown")
        if Primitive.to_proto(resource.timeshift_duration):
            res.timeshift_duration = Primitive.to_proto(resource.timeshift_duration)
        if DashboardWidgetXyChartThresholdsArray.to_proto(resource.thresholds):
            res.thresholds.extend(
                DashboardWidgetXyChartThresholdsArray.to_proto(resource.thresholds)
            )
        if DashboardWidgetXyChartXAxis.to_proto(resource.x_axis):
            res.x_axis.CopyFrom(DashboardWidgetXyChartXAxis.to_proto(resource.x_axis))
        else:
            res.ClearField("x_axis")
        if DashboardWidgetXyChartYAxis.to_proto(resource.y_axis):
            res.y_axis.CopyFrom(DashboardWidgetXyChartYAxis.to_proto(resource.y_axis))
        else:
            res.ClearField("y_axis")
        if DashboardWidgetXyChartChartOptions.to_proto(resource.chart_options):
            res.chart_options.CopyFrom(
                DashboardWidgetXyChartChartOptions.to_proto(resource.chart_options)
            )
        else:
            res.ClearField("chart_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChart(
            data_sets=resource.data_sets,
            source_drilldown=resource.source_drilldown,
            metric_drilldown=resource.metric_drilldown,
            timeshift_duration=resource.timeshift_duration,
            thresholds=resource.thresholds,
            x_axis=resource.x_axis,
            y_axis=resource.y_axis,
            chart_options=resource.chart_options,
        )


class DashboardWidgetXyChartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChart.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChart.from_proto(i) for i in resources]


class DashboardWidgetXyChartDataSets(object):
    def __init__(
        self,
        time_series_query: dict = None,
        plot_type: str = None,
        legend_template: str = None,
        min_alignment_period: str = None,
    ):
        self.time_series_query = time_series_query
        self.plot_type = plot_type
        self.legend_template = legend_template
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartDataSets()
        if DashboardWidgetXyChartDataSetsTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardWidgetXyChartDataSetsPlotTypeEnum.to_proto(resource.plot_type):
            res.plot_type = DashboardWidgetXyChartDataSetsPlotTypeEnum.to_proto(
                resource.plot_type
            )
        if Primitive.to_proto(resource.legend_template):
            res.legend_template = Primitive.to_proto(resource.legend_template)
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSets(
            time_series_query=resource.time_series_query,
            plot_type=resource.plot_type,
            legend_template=resource.legend_template,
            min_alignment_period=resource.min_alignment_period,
        )


class DashboardWidgetXyChartDataSetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartDataSets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartDataSets.from_proto(i) for i in resources]


class DashboardWidgetXyChartDataSetsTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        api_source: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.api_source = api_source
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery()
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum.to_proto(
            resource.api_source
        ):
            res.api_source = DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum.to_proto(
                resource.api_source
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQuery(
            time_series_filter=resource.time_series_filter,
            time_series_filter_ratio=resource.time_series_filter_ratio,
            time_series_query_language=resource.time_series_query_language,
            api_source=resource.api_source,
            unit_override=resource.unit_override,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQuery.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
            filter=resource.filter,
            aggregation=resource.aggregation,
            secondary_aggregation=resource.secondary_aggregation,
            pick_time_series_filter=resource.pick_time_series_filter,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=resource.ranking_method,
            num_time_series=resource.num_time_series,
            direction=resource.direction,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(object):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=resource.numerator,
            denominator=resource.denominator,
            secondary_aggregation=resource.secondary_aggregation,
            pick_time_series_filter=resource.pick_time_series_filter,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=resource.filter, aggregation=resource.aggregation,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=resource.filter, aggregation=resource.aggregation,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=resource.ranking_method,
            num_time_series=resource.num_time_series,
            direction=resource.direction,
        )


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldown(object):
    def __init__(
        self,
        resource_type_drilldown: dict = None,
        resource_label_drilldowns: list = None,
        metadata_system_label_drilldowns: list = None,
        metadata_user_label_drilldowns: list = None,
        group_name_drilldown: dict = None,
        service_name_drilldown: dict = None,
        service_type_drilldown: dict = None,
    ):
        self.resource_type_drilldown = resource_type_drilldown
        self.resource_label_drilldowns = resource_label_drilldowns
        self.metadata_system_label_drilldowns = metadata_system_label_drilldowns
        self.metadata_user_label_drilldowns = metadata_user_label_drilldowns
        self.group_name_drilldown = group_name_drilldown
        self.service_name_drilldown = service_name_drilldown
        self.service_type_drilldown = service_type_drilldown

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldown()
        if DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown.to_proto(
            resource.resource_type_drilldown
        ):
            res.resource_type_drilldown.CopyFrom(
                DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown.to_proto(
                    resource.resource_type_drilldown
                )
            )
        else:
            res.ClearField("resource_type_drilldown")
        if DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsArray.to_proto(
            resource.resource_label_drilldowns
        ):
            res.resource_label_drilldowns.extend(
                DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsArray.to_proto(
                    resource.resource_label_drilldowns
                )
            )
        if DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsArray.to_proto(
            resource.metadata_system_label_drilldowns
        ):
            res.metadata_system_label_drilldowns.extend(
                DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsArray.to_proto(
                    resource.metadata_system_label_drilldowns
                )
            )
        if DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsArray.to_proto(
            resource.metadata_user_label_drilldowns
        ):
            res.metadata_user_label_drilldowns.extend(
                DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsArray.to_proto(
                    resource.metadata_user_label_drilldowns
                )
            )
        if DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown.to_proto(
            resource.group_name_drilldown
        ):
            res.group_name_drilldown.CopyFrom(
                DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown.to_proto(
                    resource.group_name_drilldown
                )
            )
        else:
            res.ClearField("group_name_drilldown")
        if DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown.to_proto(
            resource.service_name_drilldown
        ):
            res.service_name_drilldown.CopyFrom(
                DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown.to_proto(
                    resource.service_name_drilldown
                )
            )
        else:
            res.ClearField("service_name_drilldown")
        if DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown.to_proto(
            resource.service_type_drilldown
        ):
            res.service_type_drilldown.CopyFrom(
                DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown.to_proto(
                    resource.service_type_drilldown
                )
            )
        else:
            res.ClearField("service_type_drilldown")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldown(
            resource_type_drilldown=resource.resource_type_drilldown,
            resource_label_drilldowns=resource.resource_label_drilldowns,
            metadata_system_label_drilldowns=resource.metadata_system_label_drilldowns,
            metadata_user_label_drilldowns=resource.metadata_user_label_drilldowns,
            group_name_drilldown=resource.group_name_drilldown,
            service_name_drilldown=resource.service_name_drilldown,
            service_type_drilldown=resource.service_type_drilldown,
        )


class DashboardWidgetXyChartSourceDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartSourceDrilldown.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartSourceDrilldown.from_proto(i) for i in resources]


class DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownGroupNameDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetXyChartSourceDrilldownGroupNameDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownServiceNameDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetXyChartSourceDrilldownServiceNameDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown(object):
    def __init__(self, types: list = None):
        self.types = types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown()
        )
        if DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnumArray.to_proto(
            resource.types
        ):
            res.types.extend(
                DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnumArray.to_proto(
                    resource.types
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown(
            types=resource.types,
        )


class DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartMetricDrilldown(object):
    def __init__(
        self,
        metric_type_drilldown: dict = None,
        metric_label_drilldowns: list = None,
        metric_group_by_drilldown: dict = None,
    ):
        self.metric_type_drilldown = metric_type_drilldown
        self.metric_label_drilldowns = metric_label_drilldowns
        self.metric_group_by_drilldown = metric_group_by_drilldown

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldown()
        if DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown.to_proto(
            resource.metric_type_drilldown
        ):
            res.metric_type_drilldown.CopyFrom(
                DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown.to_proto(
                    resource.metric_type_drilldown
                )
            )
        else:
            res.ClearField("metric_type_drilldown")
        if DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsArray.to_proto(
            resource.metric_label_drilldowns
        ):
            res.metric_label_drilldowns.extend(
                DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsArray.to_proto(
                    resource.metric_label_drilldowns
                )
            )
        if DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown.to_proto(
            resource.metric_group_by_drilldown
        ):
            res.metric_group_by_drilldown.CopyFrom(
                DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown.to_proto(
                    resource.metric_group_by_drilldown
                )
            )
        else:
            res.ClearField("metric_group_by_drilldown")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartMetricDrilldown(
            metric_type_drilldown=resource.metric_type_drilldown,
            metric_label_drilldowns=resource.metric_label_drilldowns,
            metric_group_by_drilldown=resource.metric_group_by_drilldown,
        )


class DashboardWidgetXyChartMetricDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartMetricDrilldown.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartMetricDrilldown.from_proto(i) for i in resources]


class DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown(object):
    def __init__(self, target_value: str = None):
        self.target_value = target_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown(
            target_value=resource.target_value,
        )


class DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown(object):
    def __init__(
        self,
        resource_labels: list = None,
        metric_labels: list = None,
        metadata_system_labels: list = None,
        metadata_user_labels: list = None,
        reducer: str = None,
    ):
        self.resource_labels = resource_labels
        self.metric_labels = metric_labels
        self.metadata_system_labels = metadata_system_labels
        self.metadata_user_labels = metadata_user_labels
        self.reducer = reducer

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown()
        )
        if Primitive.to_proto(resource.resource_labels):
            res.resource_labels.extend(Primitive.to_proto(resource.resource_labels))
        if Primitive.to_proto(resource.metric_labels):
            res.metric_labels.extend(Primitive.to_proto(resource.metric_labels))
        if Primitive.to_proto(resource.metadata_system_labels):
            res.metadata_system_labels.extend(
                Primitive.to_proto(resource.metadata_system_labels)
            )
        if Primitive.to_proto(resource.metadata_user_labels):
            res.metadata_user_labels.extend(
                Primitive.to_proto(resource.metadata_user_labels)
            )
        if DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum.to_proto(
            resource.reducer
        ):
            res.reducer = DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum.to_proto(
                resource.reducer
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown(
            resource_labels=resource.resource_labels,
            metric_labels=resource.metric_labels,
            metadata_system_labels=resource.metadata_system_labels,
            metadata_user_labels=resource.metadata_user_labels,
            reducer=resource.reducer,
        )


class DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetXyChartThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartThresholds()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardWidgetXyChartThresholdsColorEnum.to_proto(resource.color):
            res.color = DashboardWidgetXyChartThresholdsColorEnum.to_proto(
                resource.color
            )
        if DashboardWidgetXyChartThresholdsDirectionEnum.to_proto(resource.direction):
            res.direction = DashboardWidgetXyChartThresholdsDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartThresholds(
            label=resource.label,
            value=resource.value,
            color=resource.color,
            direction=resource.direction,
        )


class DashboardWidgetXyChartThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartThresholds.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartThresholds.from_proto(i) for i in resources]


class DashboardWidgetXyChartXAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartXAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartXAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardWidgetXyChartXAxisScaleEnum.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartXAxis(label=resource.label, scale=resource.scale,)


class DashboardWidgetXyChartXAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartXAxis.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartXAxis.from_proto(i) for i in resources]


class DashboardWidgetXyChartYAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartYAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartYAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardWidgetXyChartYAxisScaleEnum.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartYAxis(label=resource.label, scale=resource.scale,)


class DashboardWidgetXyChartYAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartYAxis.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartYAxis.from_proto(i) for i in resources]


class DashboardWidgetXyChartChartOptions(object):
    def __init__(self, mode: str = None, show_legend: bool = None):
        self.mode = mode
        self.show_legend = show_legend

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetXyChartChartOptions()
        if DashboardWidgetXyChartChartOptionsModeEnum.to_proto(resource.mode):
            res.mode = DashboardWidgetXyChartChartOptionsModeEnum.to_proto(
                resource.mode
            )
        if Primitive.to_proto(resource.show_legend):
            res.show_legend = Primitive.to_proto(resource.show_legend)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartChartOptions(
            mode=resource.mode, show_legend=resource.show_legend,
        )


class DashboardWidgetXyChartChartOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetXyChartChartOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetXyChartChartOptions.from_proto(i) for i in resources]


class DashboardWidgetScorecard(object):
    def __init__(
        self,
        time_series_query: dict = None,
        source_drilldown: dict = None,
        metric_drilldown: dict = None,
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.source_drilldown = source_drilldown
        self.metric_drilldown = metric_drilldown
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecard()
        if DashboardWidgetScorecardTimeSeriesQuery.to_proto(resource.time_series_query):
            res.time_series_query.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardWidgetScorecardSourceDrilldown.to_proto(resource.source_drilldown):
            res.source_drilldown.CopyFrom(
                DashboardWidgetScorecardSourceDrilldown.to_proto(
                    resource.source_drilldown
                )
            )
        else:
            res.ClearField("source_drilldown")
        if DashboardWidgetScorecardMetricDrilldown.to_proto(resource.metric_drilldown):
            res.metric_drilldown.CopyFrom(
                DashboardWidgetScorecardMetricDrilldown.to_proto(
                    resource.metric_drilldown
                )
            )
        else:
            res.ClearField("metric_drilldown")
        if DashboardWidgetScorecardGaugeView.to_proto(resource.gauge_view):
            res.gauge_view.CopyFrom(
                DashboardWidgetScorecardGaugeView.to_proto(resource.gauge_view)
            )
        else:
            res.ClearField("gauge_view")
        if DashboardWidgetScorecardSparkChartView.to_proto(resource.spark_chart_view):
            res.spark_chart_view.CopyFrom(
                DashboardWidgetScorecardSparkChartView.to_proto(
                    resource.spark_chart_view
                )
            )
        else:
            res.ClearField("spark_chart_view")
        if DashboardWidgetScorecardThresholdsArray.to_proto(resource.thresholds):
            res.thresholds.extend(
                DashboardWidgetScorecardThresholdsArray.to_proto(resource.thresholds)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecard(
            time_series_query=resource.time_series_query,
            source_drilldown=resource.source_drilldown,
            metric_drilldown=resource.metric_drilldown,
            gauge_view=resource.gauge_view,
            spark_chart_view=resource.spark_chart_view,
            thresholds=resource.thresholds,
        )


class DashboardWidgetScorecardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecard.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetScorecard.from_proto(i) for i in resources]


class DashboardWidgetScorecardTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        api_source: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.api_source = api_source
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQuery()
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum.to_proto(
            resource.api_source
        ):
            res.api_source = DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum.to_proto(
                resource.api_source
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQuery(
            time_series_filter=resource.time_series_filter,
            time_series_filter_ratio=resource.time_series_filter_ratio,
            time_series_query_language=resource.time_series_query_language,
            api_source=resource.api_source,
            unit_override=resource.unit_override,
        )


class DashboardWidgetScorecardTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardTimeSeriesQuery.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQuery.from_proto(i) for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(
            filter=resource.filter,
            aggregation=resource.aggregation,
            secondary_aggregation=resource.secondary_aggregation,
            pick_time_series_filter=resource.pick_time_series_filter,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=resource.ranking_method,
            num_time_series=resource.num_time_series,
            direction=resource.direction,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(object):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=resource.numerator,
            denominator=resource.denominator,
            secondary_aggregation=resource.secondary_aggregation,
            pick_time_series_filter=resource.pick_time_series_filter,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(object):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=resource.filter, aggregation=resource.aggregation,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(object):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=resource.filter, aggregation=resource.aggregation,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=resource.alignment_period,
            per_series_aligner=resource.per_series_aligner,
            cross_series_reducer=resource.cross_series_reducer,
            group_by_fields=resource.group_by_fields,
            reduce_fraction_less_than_params=resource.reduce_fraction_less_than_params,
            reduce_make_distribution_params=resource.reduce_make_distribution_params,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(
            threshold=resource.threshold,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(
            bucket_options=resource.bucket_options,
            exemplar_sampling=resource.exemplar_sampling,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(
            linear_buckets=resource.linear_buckets,
            exponential_buckets=resource.exponential_buckets,
            explicit_buckets=resource.explicit_buckets,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            width=resource.width,
            offset=resource.offset,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=resource.num_finite_buckets,
            growth_factor=resource.growth_factor,
            scale=resource.scale,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=resource.bounds,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(
            minimum_value=resource.minimum_value,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=resource.ranking_method,
            num_time_series=resource.num_time_series,
            direction=resource.direction,
        )


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldown(object):
    def __init__(
        self,
        resource_type_drilldown: dict = None,
        resource_label_drilldowns: list = None,
        metadata_system_label_drilldowns: list = None,
        metadata_user_label_drilldowns: list = None,
        group_name_drilldown: dict = None,
        service_name_drilldown: dict = None,
        service_type_drilldown: dict = None,
    ):
        self.resource_type_drilldown = resource_type_drilldown
        self.resource_label_drilldowns = resource_label_drilldowns
        self.metadata_system_label_drilldowns = metadata_system_label_drilldowns
        self.metadata_user_label_drilldowns = metadata_user_label_drilldowns
        self.group_name_drilldown = group_name_drilldown
        self.service_name_drilldown = service_name_drilldown
        self.service_type_drilldown = service_type_drilldown

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldown()
        if DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown.to_proto(
            resource.resource_type_drilldown
        ):
            res.resource_type_drilldown.CopyFrom(
                DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown.to_proto(
                    resource.resource_type_drilldown
                )
            )
        else:
            res.ClearField("resource_type_drilldown")
        if DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsArray.to_proto(
            resource.resource_label_drilldowns
        ):
            res.resource_label_drilldowns.extend(
                DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsArray.to_proto(
                    resource.resource_label_drilldowns
                )
            )
        if DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsArray.to_proto(
            resource.metadata_system_label_drilldowns
        ):
            res.metadata_system_label_drilldowns.extend(
                DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsArray.to_proto(
                    resource.metadata_system_label_drilldowns
                )
            )
        if DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsArray.to_proto(
            resource.metadata_user_label_drilldowns
        ):
            res.metadata_user_label_drilldowns.extend(
                DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsArray.to_proto(
                    resource.metadata_user_label_drilldowns
                )
            )
        if DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown.to_proto(
            resource.group_name_drilldown
        ):
            res.group_name_drilldown.CopyFrom(
                DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown.to_proto(
                    resource.group_name_drilldown
                )
            )
        else:
            res.ClearField("group_name_drilldown")
        if DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown.to_proto(
            resource.service_name_drilldown
        ):
            res.service_name_drilldown.CopyFrom(
                DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown.to_proto(
                    resource.service_name_drilldown
                )
            )
        else:
            res.ClearField("service_name_drilldown")
        if DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown.to_proto(
            resource.service_type_drilldown
        ):
            res.service_type_drilldown.CopyFrom(
                DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown.to_proto(
                    resource.service_type_drilldown
                )
            )
        else:
            res.ClearField("service_type_drilldown")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldown(
            resource_type_drilldown=resource.resource_type_drilldown,
            resource_label_drilldowns=resource.resource_label_drilldowns,
            metadata_system_label_drilldowns=resource.metadata_system_label_drilldowns,
            metadata_user_label_drilldowns=resource.metadata_user_label_drilldowns,
            group_name_drilldown=resource.group_name_drilldown,
            service_name_drilldown=resource.service_name_drilldown,
            service_type_drilldown=resource.service_type_drilldown,
        )


class DashboardWidgetScorecardSourceDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardSourceDrilldown.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldown.from_proto(i) for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownGroupNameDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetScorecardSourceDrilldownGroupNameDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown(object):
    def __init__(self, target_values: list = None):
        self.target_values = target_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownServiceNameDrilldown()
        )
        if Primitive.to_proto(resource.target_values):
            res.target_values.extend(Primitive.to_proto(resource.target_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown(
            target_values=resource.target_values,
        )


class DashboardWidgetScorecardSourceDrilldownServiceNameDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown(object):
    def __init__(self, types: list = None):
        self.types = types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown()
        )
        if DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnumArray.to_proto(
            resource.types
        ):
            res.types.extend(
                DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnumArray.to_proto(
                    resource.types
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown(
            types=resource.types,
        )


class DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardMetricDrilldown(object):
    def __init__(
        self,
        metric_type_drilldown: dict = None,
        metric_label_drilldowns: list = None,
        metric_group_by_drilldown: dict = None,
    ):
        self.metric_type_drilldown = metric_type_drilldown
        self.metric_label_drilldowns = metric_label_drilldowns
        self.metric_group_by_drilldown = metric_group_by_drilldown

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldown()
        if DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown.to_proto(
            resource.metric_type_drilldown
        ):
            res.metric_type_drilldown.CopyFrom(
                DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown.to_proto(
                    resource.metric_type_drilldown
                )
            )
        else:
            res.ClearField("metric_type_drilldown")
        if DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsArray.to_proto(
            resource.metric_label_drilldowns
        ):
            res.metric_label_drilldowns.extend(
                DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsArray.to_proto(
                    resource.metric_label_drilldowns
                )
            )
        if DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown.to_proto(
            resource.metric_group_by_drilldown
        ):
            res.metric_group_by_drilldown.CopyFrom(
                DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown.to_proto(
                    resource.metric_group_by_drilldown
                )
            )
        else:
            res.ClearField("metric_group_by_drilldown")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardMetricDrilldown(
            metric_type_drilldown=resource.metric_type_drilldown,
            metric_label_drilldowns=resource.metric_label_drilldowns,
            metric_group_by_drilldown=resource.metric_group_by_drilldown,
        )


class DashboardWidgetScorecardMetricDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardMetricDrilldown.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardMetricDrilldown.from_proto(i) for i in resources
        ]


class DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown(object):
    def __init__(self, target_value: str = None):
        self.target_value = target_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown(
            target_value=resource.target_value,
        )


class DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns(object):
    def __init__(
        self,
        label: str = None,
        logical_operator: str = None,
        value_restrictions: list = None,
    ):
        self.label = label
        self.logical_operator = logical_operator
        self.value_restrictions = value_restrictions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray.to_proto(
            resource.value_restrictions
        ):
            res.value_restrictions.extend(
                DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray.to_proto(
                    resource.value_restrictions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns(
            label=resource.label,
            logical_operator=resource.logical_operator,
            value_restrictions=resource.value_restrictions,
        )


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions(
    object
):
    def __init__(self, target_value: str = None, comparator: str = None):
        self.target_value = target_value
        self.comparator = comparator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions()
        )
        if Primitive.to_proto(resource.target_value):
            res.target_value = Primitive.to_proto(resource.target_value)
        if DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
            resource.comparator
        ):
            res.comparator = DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.to_proto(
                resource.comparator
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions(
            target_value=resource.target_value, comparator=resource.comparator,
        )


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions.from_proto(
                i
            )
            for i in resources
        ]


class DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown(object):
    def __init__(
        self,
        resource_labels: list = None,
        metric_labels: list = None,
        metadata_system_labels: list = None,
        metadata_user_labels: list = None,
        reducer: str = None,
    ):
        self.resource_labels = resource_labels
        self.metric_labels = metric_labels
        self.metadata_system_labels = metadata_system_labels
        self.metadata_user_labels = metadata_user_labels
        self.reducer = reducer

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown()
        )
        if Primitive.to_proto(resource.resource_labels):
            res.resource_labels.extend(Primitive.to_proto(resource.resource_labels))
        if Primitive.to_proto(resource.metric_labels):
            res.metric_labels.extend(Primitive.to_proto(resource.metric_labels))
        if Primitive.to_proto(resource.metadata_system_labels):
            res.metadata_system_labels.extend(
                Primitive.to_proto(resource.metadata_system_labels)
            )
        if Primitive.to_proto(resource.metadata_user_labels):
            res.metadata_user_labels.extend(
                Primitive.to_proto(resource.metadata_user_labels)
            )
        if DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum.to_proto(
            resource.reducer
        ):
            res.reducer = DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum.to_proto(
                resource.reducer
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown(
            resource_labels=resource.resource_labels,
            metric_labels=resource.metric_labels,
            metadata_system_labels=resource.metadata_system_labels,
            metadata_user_labels=resource.metadata_user_labels,
            reducer=resource.reducer,
        )


class DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown.from_proto(i)
            for i in resources
        ]


class DashboardWidgetScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardGaugeView()
        if Primitive.to_proto(resource.lower_bound):
            res.lower_bound = Primitive.to_proto(resource.lower_bound)
        if Primitive.to_proto(resource.upper_bound):
            res.upper_bound = Primitive.to_proto(resource.upper_bound)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardGaugeView(
            lower_bound=resource.lower_bound, upper_bound=resource.upper_bound,
        )


class DashboardWidgetScorecardGaugeViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardGaugeView.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetScorecardGaugeView.from_proto(i) for i in resources]


class DashboardWidgetScorecardSparkChartView(object):
    def __init__(self, spark_chart_type: str = None, min_alignment_period: str = None):
        self.spark_chart_type = spark_chart_type
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardSparkChartView()
        if DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.to_proto(
            resource.spark_chart_type
        ):
            res.spark_chart_type = DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.to_proto(
                resource.spark_chart_type
            )
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardSparkChartView(
            spark_chart_type=resource.spark_chart_type,
            min_alignment_period=resource.min_alignment_period,
        )


class DashboardWidgetScorecardSparkChartViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardSparkChartView.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetScorecardSparkChartView.from_proto(i) for i in resources]


class DashboardWidgetScorecardThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetScorecardThresholds()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardWidgetScorecardThresholdsColorEnum.to_proto(resource.color):
            res.color = DashboardWidgetScorecardThresholdsColorEnum.to_proto(
                resource.color
            )
        if DashboardWidgetScorecardThresholdsDirectionEnum.to_proto(resource.direction):
            res.direction = DashboardWidgetScorecardThresholdsDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardThresholds(
            label=resource.label,
            value=resource.value,
            color=resource.color,
            direction=resource.direction,
        )


class DashboardWidgetScorecardThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetScorecardThresholds.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetScorecardThresholds.from_proto(i) for i in resources]


class DashboardWidgetText(object):
    def __init__(self, content: str = None, format: str = None):
        self.content = content
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardWidgetTextFormatEnum.to_proto(resource.format):
            res.format = DashboardWidgetTextFormatEnum.to_proto(resource.format)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetText(content=resource.content, format=resource.format,)


class DashboardWidgetTextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetText.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetText.from_proto(i) for i in resources]


class DashboardWidgetBlank(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringDashboardWidgetBlank()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetBlank()


class DashboardWidgetBlankArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardWidgetBlank.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardWidgetBlank.from_proto(i) for i in resources]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum") :]


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum.Value(
            "MonitoringDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum"
            ) :
        ]


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum.Value(
            "MonitoringDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartThresholdsColorEnum.Value(
            "MonitoringDashboardWidgetXyChartThresholdsColorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartThresholdsColorEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetXyChartThresholdsColorEnum") :]


class DashboardWidgetXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum.Value(
            "MonitoringDashboardWidgetXyChartThresholdsDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len("MonitoringDashboardWidgetXyChartThresholdsDirectionEnum") :
        ]


class DashboardWidgetXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartXAxisScaleEnum.Value(
            "MonitoringDashboardWidgetXyChartXAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartXAxisScaleEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetXyChartXAxisScaleEnum") :]


class DashboardWidgetXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartYAxisScaleEnum.Value(
            "MonitoringDashboardWidgetXyChartYAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartYAxisScaleEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetXyChartYAxisScaleEnum") :]


class DashboardWidgetXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartChartOptionsModeEnum.Value(
            "MonitoringDashboardWidgetXyChartChartOptionsModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetXyChartChartOptionsModeEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetXyChartChartOptionsModeEnum") :]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum.Value(
            "MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum.Name(
            resource
        )[
            len("MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum") :
        ]


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum.Value(
            "MonitoringDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum"
            ) :
        ]


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.Value(
            "MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum"
            ) :
        ]


class DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.Value(
            "MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum"
            ) :
        ]


class DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum.Value(
            "MonitoringDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len("MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum") :
        ]


class DashboardWidgetScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardThresholdsColorEnum.Value(
            "MonitoringDashboardWidgetScorecardThresholdsColorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardThresholdsColorEnum.Name(
            resource
        )[len("MonitoringDashboardWidgetScorecardThresholdsColorEnum") :]


class DashboardWidgetScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum.Value(
            "MonitoringDashboardWidgetScorecardThresholdsDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len("MonitoringDashboardWidgetScorecardThresholdsDirectionEnum") :
        ]


class DashboardWidgetTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetTextFormatEnum.Value(
            "MonitoringDashboardWidgetTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardWidgetTextFormatEnum.Name(resource)[
            len("MonitoringDashboardWidgetTextFormatEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
