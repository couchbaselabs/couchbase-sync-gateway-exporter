local grafana = import 'grafonnet/grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local singlestat = grafana.singlestat;
local graphPanel = grafana.graphPanel;
local prometheus = grafana.prometheus;

dashboard.new(
  'Couchbase Sync Gateway Dashboard',
  description='',
  refresh='10s',
  time_from='now-1h',
  tags=['couchbase'],
  editable=true,
)
.addTemplate(
  grafana.template.datasource(
    'PROMETHEUS_DS',
    'prometheus',
    'Prometheus',
    hide='label',
  )
)
.addTemplate(
  grafana.template.new(
    'instance',
    '$PROMETHEUS_DS',
    'label_values(sgw_up, instance)',
    label='Instance',
    refresh='load',
  )
)
.addRow(
  row.new(
    title='Resources',
    collapse=false,
  )
  .addPanel(
    graphPanel.new(
      'CPU Utilization',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='percent',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_process_cpu_percent_utilization{instance=~"$instance"}',
        legendFormat='{{ instance }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Memory Utilization',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_process_memory_resident{instance=~"$instance"}',
        legendFormat='{{ instance }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Network Transfer',
      span=12,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='Bps',
    )
    .addSeriesOverride(
      {
        alias: '/sent/',
        transform: 'negative-Y',
      }
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_resource_utilization_pub_net_bytes_sent{instance=~"$instance"}[5m]) + rate(sgw_resource_utilization_admin_net_bytes_sent{instance=~"$instance"}[5m])',
        legendFormat='{{ instance }} sent',
      )
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_resource_utilization_pub_net_bytes_recv{instance=~"$instance"}[5m]) + rate(sgw_resource_utilization_admin_net_bytes_recv{instance=~"$instance"}[5m])',
        legendFormat='{{ instance }} recv',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Overall Heap Usage',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_sys{instance=~"$instance"}',
        legendFormat='{{ instance }} sys',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_heapalloc{instance=~"$instance"}',
        legendFormat='{{ instance }} heapalloc',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_heapidle{instance=~"$instance"}',
        legendFormat='{{ instance }} heapidle',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_heapreleased{instance=~"$instance"}',
        legendFormat='{{ instance }} heapreleased',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Overall Stack Usage',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_stacksys{instance=~"$instance"}',
        legendFormat='{{ instance }} stacksys',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_stackinuse{instance=~"$instance"}',
        legendFormat='{{ instance }} stackinuse',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_stackinuse{instance=~"$instance"}',
        legendFormat='{{ instance }} stackinuse',
      )
    )
  )
)
