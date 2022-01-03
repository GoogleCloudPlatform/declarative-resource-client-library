# Copyright 2022 Google LLC. All Rights Reserved.
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
        etag: str = None,
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
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ApplyMonitoringBetaDashboardRequest()
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

        response = stub.ApplyMonitoringBetaDashboard(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.grid_layout = DashboardGridLayout.from_proto(response.grid_layout)
        self.mosaic_layout = DashboardMosaicLayout.from_proto(response.mosaic_layout)
        self.row_layout = DashboardRowLayout.from_proto(response.row_layout)
        self.column_layout = DashboardColumnLayout.from_proto(response.column_layout)
        self.project = Primitive.from_proto(response.project)
        self.etag = Primitive.from_proto(response.etag)

    def delete(self):
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.DeleteMonitoringBetaDashboardRequest()
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

        response = stub.DeleteMonitoringBetaDashboard(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ListMonitoringBetaDashboardRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringBetaDashboard(request).items

    def to_proto(self):
        resource = dashboard_pb2.MonitoringBetaDashboard()
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

        res = dashboard_pb2.MonitoringBetaDashboardGridLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardWidgetArray.to_proto(resource.widgets):
            res.widgets.extend(DashboardWidgetArray.to_proto(resource.widgets))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayout(
            columns=Primitive.from_proto(resource.columns),
            widgets=DashboardWidgetArray.from_proto(resource.widgets),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardMosaicLayoutTilesArray.to_proto(resource.tiles):
            res.tiles.extend(DashboardMosaicLayoutTilesArray.to_proto(resource.tiles))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayout(
            columns=Primitive.from_proto(resource.columns),
            tiles=DashboardMosaicLayoutTilesArray.from_proto(resource.tiles),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTiles()
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
            x_pos=Primitive.from_proto(resource.x_pos),
            y_pos=Primitive.from_proto(resource.y_pos),
            width=Primitive.from_proto(resource.width),
            height=Primitive.from_proto(resource.height),
            widget=DashboardWidget.from_proto(resource.widget),
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

        res = dashboard_pb2.MonitoringBetaDashboardRowLayout()
        if DashboardRowLayoutRowsArray.to_proto(resource.rows):
            res.rows.extend(DashboardRowLayoutRowsArray.to_proto(resource.rows))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayout(
            rows=DashboardRowLayoutRowsArray.from_proto(resource.rows),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRows()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if DashboardWidgetArray.to_proto(resource.widgets):
            res.widgets.extend(DashboardWidgetArray.to_proto(resource.widgets))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRows(
            weight=Primitive.from_proto(resource.weight),
            widgets=DashboardWidgetArray.from_proto(resource.widgets),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayout()
        if DashboardColumnLayoutColumnsArray.to_proto(resource.columns):
            res.columns.extend(
                DashboardColumnLayoutColumnsArray.to_proto(resource.columns)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayout(
            columns=DashboardColumnLayoutColumnsArray.from_proto(resource.columns),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumns()
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
            weight=Primitive.from_proto(resource.weight),
            widgets=DashboardWidgetArray.from_proto(resource.widgets),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidget()
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
            title=Primitive.from_proto(resource.title),
            xy_chart=DashboardWidgetXyChart.from_proto(resource.xy_chart),
            scorecard=DashboardWidgetScorecard.from_proto(resource.scorecard),
            text=DashboardWidgetText.from_proto(resource.text),
            blank=DashboardWidgetBlank.from_proto(resource.blank),
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
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChart()
        if DashboardWidgetXyChartDataSetsArray.to_proto(resource.data_sets):
            res.data_sets.extend(
                DashboardWidgetXyChartDataSetsArray.to_proto(resource.data_sets)
            )
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
            data_sets=DashboardWidgetXyChartDataSetsArray.from_proto(
                resource.data_sets
            ),
            timeshift_duration=Primitive.from_proto(resource.timeshift_duration),
            thresholds=DashboardWidgetXyChartThresholdsArray.from_proto(
                resource.thresholds
            ),
            x_axis=DashboardWidgetXyChartXAxis.from_proto(resource.x_axis),
            y_axis=DashboardWidgetXyChartYAxis.from_proto(resource.y_axis),
            chart_options=DashboardWidgetXyChartChartOptions.from_proto(
                resource.chart_options
            ),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSets()
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
            time_series_query=DashboardWidgetXyChartDataSetsTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            plot_type=DashboardWidgetXyChartDataSetsPlotTypeEnum.from_proto(
                resource.plot_type
            ),
            legend_template=Primitive.from_proto(resource.legend_template),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
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
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQuery()
        )
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
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQuery(
            time_series_filter=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            ranking_method=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
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
            numerator=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            ranking_method=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChartThresholds()
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
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardWidgetXyChartThresholdsColorEnum.from_proto(resource.color),
            direction=DashboardWidgetXyChartThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChartXAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartXAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardWidgetXyChartXAxisScaleEnum.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartXAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardWidgetXyChartXAxisScaleEnum.from_proto(resource.scale),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChartYAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardWidgetXyChartYAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardWidgetXyChartYAxisScaleEnum.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartYAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardWidgetXyChartYAxisScaleEnum.from_proto(resource.scale),
        )


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
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardWidgetXyChartChartOptions()
        if DashboardWidgetXyChartChartOptionsModeEnum.to_proto(resource.mode):
            res.mode = DashboardWidgetXyChartChartOptionsModeEnum.to_proto(
                resource.mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetXyChartChartOptions(
            mode=DashboardWidgetXyChartChartOptionsModeEnum.from_proto(resource.mode),
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
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardWidgetScorecard()
        if DashboardWidgetScorecardTimeSeriesQuery.to_proto(resource.time_series_query):
            res.time_series_query.CopyFrom(
                DashboardWidgetScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
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
            time_series_query=DashboardWidgetScorecardTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            gauge_view=DashboardWidgetScorecardGaugeView.from_proto(
                resource.gauge_view
            ),
            spark_chart_view=DashboardWidgetScorecardSparkChartView.from_proto(
                resource.spark_chart_view
            ),
            thresholds=DashboardWidgetScorecardThresholdsArray.from_proto(
                resource.thresholds
            ),
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
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQuery()
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
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQuery(
            time_series_filter=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            ranking_method=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio()
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
            numerator=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(object):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
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
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
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
            dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            ranking_method=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
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


class DashboardWidgetScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardWidgetScorecardGaugeView()
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
            lower_bound=Primitive.from_proto(resource.lower_bound),
            upper_bound=Primitive.from_proto(resource.upper_bound),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetScorecardSparkChartView()
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
            spark_chart_type=DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.from_proto(
                resource.spark_chart_type
            ),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetScorecardThresholds()
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
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardWidgetScorecardThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardWidgetScorecardThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardWidgetTextFormatEnum.to_proto(resource.format):
            res.format = DashboardWidgetTextFormatEnum.to_proto(resource.format)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardWidgetText(
            content=Primitive.from_proto(resource.content),
            format=DashboardWidgetTextFormatEnum.from_proto(resource.format),
        )


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

        res = dashboard_pb2.MonitoringBetaDashboardWidgetBlank()
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
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetXyChartDataSetsPlotTypeEnum") :
        ]


class DashboardWidgetXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetXyChartThresholdsColorEnum") :
        ]


class DashboardWidgetXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetXyChartThresholdsDirectionEnum") :
        ]


class DashboardWidgetXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum.Name(
            resource
        )[len("MonitoringBetaDashboardWidgetXyChartXAxisScaleEnum") :]


class DashboardWidgetXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum.Name(
            resource
        )[len("MonitoringBetaDashboardWidgetXyChartYAxisScaleEnum") :]


class DashboardWidgetXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum.Value(
            "MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetXyChartChartOptionsModeEnum") :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardWidgetScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetScorecardThresholdsColorEnum") :
        ]


class DashboardWidgetScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardWidgetScorecardThresholdsDirectionEnum") :
        ]


class DashboardWidgetTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetTextFormatEnum.Value(
            "MonitoringBetaDashboardWidgetTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardWidgetTextFormatEnum.Name(resource)[
            len("MonitoringBetaDashboardWidgetTextFormatEnum") :
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
