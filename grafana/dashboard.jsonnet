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
.addTemplate(
  grafana.template.new(
    'database',
    '$PROMETHEUS_DS',
    'label_values(sgw_database_sequence_get_count{instance=~"$instance"}, database)',
    label='Database',
    refresh='load',
  )
)
.addRow(
  row.new(
    title='Resources',
    collapse=true,
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
      nullPointMode='null as zero',
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
      nullPointMode='null as zero',
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
      nullPointMode='null as zero',
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
      nullPointMode='null as zero',
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
  .addPanel(
    graphPanel.new(
      'Garbage Collection',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ns',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_go_memstats_pausetotalns{instance=~"$instance"}',
        legendFormat='{{ instance }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Logging',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_error_count{instance=~"$instance"}',
        legendFormat='{{ instance }} errors',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_resource_utilization_warn_count{instance=~"$instance"}',
        legendFormat='{{ instance }} warns',
      )
    )
  )
)
.addRow(
  row.new(
    title='Cache',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Channel Cache Utilization',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_chan_cache_active_revs{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} active revs',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_chan_cache_tombstone_revs{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} thombstone revs',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_chan_cache_removal_revs{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} removal revs',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Channel Hit/Miss',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
    )
    .addSeriesOverride(
      {
        alias: '/misses/',
        transform: 'negative-Y',
      }
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_cache_chan_cache_hits{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }} hits',
      )
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_cache_chan_cache_misses{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }} misses',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Channel Cache Size',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        '(
            sgw_cache_chan_cache_active_revs{instance=~"$instance",database=~"$database"} +
            sgw_cache_chan_cache_tombstone_revs{instance=~"$instance",database=~"$database"} +
            sgw_cache_chan_cache_removal_revs{instance=~"$instance",database=~"$database"}
          ) / sgw_cache_chan_cache_num_channels{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} average',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_chan_cache_max_entries{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} max',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Channel Cache Count',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_chan_cache_num_channels{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Channel Hit/Miss',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
    )
    .addTarget(
      prometheus.target(
        'sgw_cache_rev_cache_hits{instance=~"$instance",database=~"$database"} / (
          sgw_cache_rev_cache_hits{instance=~"$instance",database=~"$database"} +
          sgw_cache_rev_cache_misses{instance=~"$instance",database=~"$database"}
        )',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='Database Stats',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Number of Active Replications',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_database_num_replications_active{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'New Replications Per Second',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_database_num_replications_total{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Closed Replications',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_database_num_replications_total{instance=~"$instance",database=~"$database"} - sgw_database_num_replications_active{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Document writes/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_database_num_doc_writes{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      '% of docs in conflict',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='percent',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_push_conflict_write_count{instance=~"$instance",database=~"$database"} / sgw_database_num_doc_writes{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Document reads/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_database_num_doc_reads_rest{instance=~"$instance",database=~"$database"}[5m]) +
          increase(sgw_database_num_doc_reads_blip{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='Delta Sync',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Delta Cache Hit Ratio',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_delta_sync_delta_cache_hit{instance=~"$instance",database=~"$database"} /
          sgw_delta_sync_delta_cache_miss',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Delta Hit/Miss rate',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='short',
    )
    .addTarget(
      prometheus.target(
        'sgw_delta_sync_deltas_requested{instance=~"$instance",database=~"$database"} /
          sgw_delta_sync_deltas_sent',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of documents sent to SG as a delta',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_delta_sync_delta_push_doc_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }} average',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of pull replications using deltas',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_delta_sync_delta_pull_replication_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }} average',
      )
    )
  )
)
.addRow(
  row.new(
    title='Import',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Number of docs imported',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_shared_bucket_import_import_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of errors as a result of doc import',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_shared_bucket_import_import_error_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='CBL Push Replication',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Average Write Processing Time',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_push_write_processing_time{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Documents pushed per second',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_push_doc_push_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Average Sync Function Processing Time',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ns',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_push_sync_function_time{instance=~"$instance",database=~"$database"}[5m]) /
          increase(sgw_replication_push_sync_function_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Average ProposeChange Processing Time',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_push_propose_change_time{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of ProposeChange messages / sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_push_propose_change_time{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of attachments pushed per second',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_push_attachment_push_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Average size of attachments',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_push_attachment_push_count{instance=~"$instance",database=~"$database"} / sgw_replication_push_attachment_push_bytes',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='CBL Pull Replication',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Changes Request Processing Latency',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_request_changes_count{instance=~"$instance",database=~"$database"} /
          sgw_replication_pull_request_changes_time',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Server DCP Feed Processing Latency',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_database_dcp_caching_time{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Revision Send Message Processing Latency ',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_rev_send_latency{instance=~"$instance",database=~"$database"} /
          sgw_replication_pull_rev_send_count',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of attachments pulled per second',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_pull_attachment_pull_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Average size of attachments pulled',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_attachment_pull_count{instance=~"$instance",database=~"$database"} /
          sgw_replication_pull_attachment_pull_bytes',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Documents pulled per second in 2.x',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_database_num_doc_reads_blip{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Total Replications by type',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_num_pull_repl_total_one_shot{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} one shot',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_num_pull_repl_total_continuous{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} continuous',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Active Replications by type',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_num_pull_repl_active_one_shot{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} one shot',
      )
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_num_pull_repl_active_continuous{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }} continuous',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of New Initial Replications per second',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_replication_pull_num_pull_repl_since_zero{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Replications Caught up Count',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='bytes',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw_replication_pull_num_pull_repl_caught_up{instance=~"$instance",database=~"$database"}',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='Security',
    collapse=false,
  )
  .addPanel(
    graphPanel.new(
      'Sync Function Doc Rejections Count',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_security_num_docs_rejected{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Sync Function Access Failure Count',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'increase(sgw_security_num_access_errors{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='GSI VIew / Query',
    collapse=false,
  )
  .addPanel(
    graphPanel.new(
      'Number of queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'rate(sgw::gsi::total_queries{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of channel queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_gsi_views_channels_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of access queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_gsi_views_access_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of all Docs queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_gsi_views_allDocs_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Number of role access queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'rate(sgw_gsi_views_roleAccess_count{instance=~"$instance",database=~"$database"}[5m])',
        legendFormat='{{ database }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Housekeeping Queries/sec',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='ops',
      min=0,
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'sgw::gsi::total_queries{instance=~"$instance",database=~"$database"} - (
          sgw_gsi_views_channels_count +
          sgw_gsi_views_access_count +
          sgw_gsi_views_roleAccess_count
        )',
        legendFormat='{{ database }}',
      )
    )
  )
)
