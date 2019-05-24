package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"
)

func TestMetricsSimple(t *testing.T) {
	var collector = NewCollector(newFakeClient(t, "testdata/metrics2.json", nil))
	testCollector(t, collector, func(t *testing.T, status int, body string) {
		require.Equal(t, 200, status)
		requireGauge(t, body, `sgw_cache_chan_cache_active_revs{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_chan_cache_hits{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_max_entries{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_chan_cache_misses{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_num_channels{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_pending_queries{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_removal_revs{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_tombstone_revs{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_num_skipped_seqs{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_rev_cache_hits{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_rev_cache_misses{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_abandoned_seqs{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_crc32c_match_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_dcp_caching_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_dcp_caching_time{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_dcp_received_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_dcp_received_time{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_doc_reads_bytes_blip{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_doc_writes_bytes{database="travel-sample"} 2.274404e+06`)
		requireCounter(t, body, `sgw_database_doc_writes_bytes_blip{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_num_doc_reads_blip{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_num_doc_reads_rest{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_num_doc_writes{database="travel-sample"} 2100`)
		requireGauge(t, body, `sgw_database_num_replications_active{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_num_replications_total{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_sequence_get_count{database="travel-sample"} 1`)
		requireCounter(t, body, `sgw_database_sequence_released_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_sequence_reserved_count{database="travel-sample"} 2105`)
		requireCounter(t, body, `sgw_database_warn_channels_per_doc_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_warn_grants_per_doc_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_warn_xattr_size_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_gsi_views_access_count{database="travel-sample"} 3`)
		requireGauge(t, body, `sgw_gsi_views_channels_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_gsi_views_roleAccess_count{database="travel-sample"} 3`)
		requireCounter(t, body, `sgw_replication_pull_attachment_pull_bytes{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_attachment_pull_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_max_pending{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_active_continuous{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_active_one_shot{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_caught_up{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_num_pull_repl_since_zero{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_total_continuous{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_total_one_shot{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_replications_active{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_request_changes_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_request_changes_time{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_rev_processing_time{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_rev_send_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_rev_send_latency{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_attachment_push_bytes{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_attachment_push_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_push_conflict_write_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_push_doc_push_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_propose_change_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_propose_change_time{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_sync_function_count{database="travel-sample"} 2101`)
		requireCounter(t, body, `sgw_replication_push_sync_function_time{database="travel-sample"} 3.21153379e+08`)
		requireGauge(t, body, `sgw_replication_push_write_processing_time{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_resource_utilization_admin_net_bytes_recv 0`)
		requireCounter(t, body, `sgw_resource_utilization_admin_net_bytes_sent 0`)
		requireCounter(t, body, `sgw_resource_utilization_error_count 2`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapalloc 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapidle 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapinuse 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapreleased 0`)
		requireCounter(t, body, `sgw_resource_utilization_go_memstats_pausetotalns 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_stackinuse 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_stacksys 0`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_sys 0`)
		requireGauge(t, body, `sgw_resource_utilization_goroutines_high_watermark 0`)
		requireGauge(t, body, `sgw_resource_utilization_num_goroutines 0`)
		requireGauge(t, body, `sgw_resource_utilization_process_cpu_percent_utilization 0`)
		requireGauge(t, body, `sgw_resource_utilization_process_memory_resident 0`)
		requireCounter(t, body, `sgw_resource_utilization_pub_net_bytes_recv 0`)
		requireCounter(t, body, `sgw_resource_utilization_pub_net_bytes_sent 0`)
		requireGauge(t, body, `sgw_resource_utilization_system_memory_total 0`)
		requireCounter(t, body, `sgw_resource_utilization_warn_count 4`)
		requireGauge(t, body, `sgw_security_auth_failed_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_security_auth_success_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_security_num_access_errors{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_security_num_docs_rejected{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_security_total_auth_time{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_shared_bucket_import_import_count{database="travel-sample"} 2100`)
		requireCounter(t, body, `sgw_shared_bucket_import_import_error_count{database="travel-sample"} 0`)
		// nolint: lll
		requireGauge(t, body, `sgw_shared_bucket_import_import_processing_time{database="travel-sample"} 1.9499618564e+10`)
		require.Regexp(t, "sgw_scrape_duration_seconds \\d+\\.\\d+", body)
		requireGauge(t, body, `sgw_up 1`)

	})
}

func TestMetricsWithReplication(t *testing.T) {
	var collector = NewCollector(newFakeClient(t, "testdata/metrics1.json", nil))
	testCollector(t, collector, func(t *testing.T, status int, body string) {
		require.Equal(t, 200, status)
		requireGauge(t, body, `sgw_cache_chan_cache_active_revs{database="travel-sample"} 102`)
		requireCounter(t, body, `sgw_cache_chan_cache_hits{database="travel-sample"} 3685`)
		requireGauge(t, body, `sgw_cache_chan_cache_max_entries{database="travel-sample"} 50`)
		requireCounter(t, body, `sgw_cache_chan_cache_misses{database="travel-sample"} 563`)
		requireGauge(t, body, `sgw_cache_chan_cache_num_channels{database="travel-sample"} 3`)
		requireGauge(t, body, `sgw_cache_chan_cache_pending_queries{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_removal_revs{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_cache_chan_cache_tombstone_revs{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_cache_num_skipped_seqs{database="travel-sample"} 2400`)
		requireCounter(t, body, `sgw_cache_rev_cache_hits{database="travel-sample"} 76`)
		requireCounter(t, body, `sgw_cache_rev_cache_misses{database="travel-sample"} 37935`)
		requireCounter(t, body, `sgw_database_abandoned_seqs{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_crc32c_match_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_database_dcp_caching_count{database="travel-sample"} 10033`)
		requireGauge(t, body, `sgw_database_dcp_caching_time{database="travel-sample"} 2.73613199067e+13`)
		requireGauge(t, body, `sgw_database_dcp_received_count{database="travel-sample"} 10033`)
		requireGauge(t, body, `sgw_database_dcp_received_time{database="travel-sample"} 7.263339887408e+14`)
		requireCounter(t, body, `sgw_database_doc_reads_bytes_blip{database="travel-sample"} 2.109825e+07`)
		requireCounter(t, body, `sgw_database_doc_writes_bytes{database="travel-sample"} 2.1098263e+07`)
		requireCounter(t, body, `sgw_database_doc_writes_bytes_blip{database="travel-sample"} 13`)
		requireCounter(t, body, `sgw_database_num_doc_reads_blip{database="travel-sample"} 10032`)
		requireCounter(t, body, `sgw_database_num_doc_reads_rest{database="travel-sample"} 27979`)
		requireCounter(t, body, `sgw_database_num_doc_writes{database="travel-sample"} 10033`)
		requireGauge(t, body, `sgw_database_num_replications_active{database="travel-sample"} 568`)
		requireGauge(t, body, `sgw_database_num_replications_total{database="travel-sample"} 570`)
		requireCounter(t, body, `sgw_database_sequence_get_count{database="travel-sample"} 1`)
		requireCounter(t, body, `sgw_database_sequence_released_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_sequence_reserved_count{database="travel-sample"} 10034`)
		requireCounter(t, body, `sgw_database_warn_channels_per_doc_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_warn_grants_per_doc_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_database_warn_xattr_size_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_gsi_views_access_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_gsi_views_channels_count{database="travel-sample"} 563`)
		requireGauge(t, body, `sgw_gsi_views_roleAccess_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_attachment_pull_bytes{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_pull_attachment_pull_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_max_pending{database="travel-sample"} 5343`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_active_continuous{database="travel-sample"} 2`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_active_one_shot{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_caught_up{database="travel-sample"} 2`)
		requireCounter(t, body, `sgw_replication_pull_num_pull_repl_since_zero{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_total_continuous{database="travel-sample"} 5`)
		requireGauge(t, body, `sgw_replication_pull_num_pull_repl_total_one_shot{database="travel-sample"} 565`)
		requireGauge(t, body, `sgw_replication_pull_num_replications_active{database="travel-sample"} -566`)
		requireCounter(t, body, `sgw_replication_pull_request_changes_count{database="travel-sample"} 1984`)
		requireCounter(t, body, `sgw_replication_pull_request_changes_time{database="travel-sample"} 2.8729243e+10`)
		requireGauge(t, body, `sgw_replication_pull_rev_processing_time{database="travel-sample"} 1.028721848e+11`)
		requireCounter(t, body, `sgw_replication_pull_rev_send_count{database="travel-sample"} 10032`)
		requireGauge(t, body, `sgw_replication_pull_rev_send_latency{database="travel-sample"} 8.322484988e+11`)
		requireCounter(t, body, `sgw_replication_push_attachment_push_bytes{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_replication_push_attachment_push_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_push_conflict_write_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_replication_push_doc_push_count{database="travel-sample"} 1`)
		requireCounter(t, body, `sgw_replication_push_propose_change_count{database="travel-sample"} 1`)
		requireCounter(t, body, `sgw_replication_push_propose_change_time{database="travel-sample"} 6.0633e+06`)
		requireCounter(t, body, `sgw_replication_push_sync_function_count{database="travel-sample"} 10033`)
		requireCounter(t, body, `sgw_replication_push_sync_function_time{database="travel-sample"} 1.2717184e+09`)
		requireGauge(t, body, `sgw_replication_push_write_processing_time{database="travel-sample"} 1.002994e+08`)
		requireCounter(t, body, `sgw_replication_sgr_docs_checked_sent{replication="repl-1"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_docs_checked_sent{replication="repl-101"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_docs_checked_sent{replication="repl-102"} 7828`)
		requireCounter(t, body, `sgw_replication_sgr_docs_checked_sent{replication="repl-102-pull"} 27978`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachment_bytes_transferred{replication="repl-1"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachment_bytes_transferred{replication="repl-101"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachment_bytes_transferred{replication="repl-102"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachment_bytes_transferred{replication="repl-102-pull"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachments_transferred{replication="repl-1"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachments_transferred{replication="repl-101"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachments_transferred{replication="repl-102"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_attachments_transferred{replication="repl-102-pull"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_failed_to_push{replication="repl-1"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_failed_to_push{replication="repl-101"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_failed_to_push{replication="repl-102"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_failed_to_push{replication="repl-102-pull"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_pushed{replication="repl-1"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_pushed{replication="repl-101"} 0`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_pushed{replication="repl-102"} 7778`)
		requireCounter(t, body, `sgw_replication_sgr_num_docs_pushed{replication="repl-102-pull"} 0`)
		requireCounter(t, body, `sgw_resource_utilization_admin_net_bytes_recv 6.97856009e+08`)
		requireCounter(t, body, `sgw_resource_utilization_admin_net_bytes_sent 1.66156024e+08`)
		requireCounter(t, body, `sgw_resource_utilization_error_count 9`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapalloc 5.5281624e+07`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapidle 1.16293632e+08`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapinuse 8.4017152e+07`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_heapreleased 7.8716928e+07`)
		requireCounter(t, body, `sgw_resource_utilization_go_memstats_pausetotalns 2.3147789e+09`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_stackinuse 1.015808e+06`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_stacksys 1.015808e+06`)
		requireGauge(t, body, `sgw_resource_utilization_go_memstats_sys 2.15701752e+08`)
		requireGauge(t, body, `sgw_resource_utilization_goroutines_high_watermark 52`)
		requireGauge(t, body, `sgw_resource_utilization_num_goroutines 52`)
		requireGauge(t, body, `sgw_resource_utilization_process_cpu_percent_utilization 199.8688155922039`)
		requireGauge(t, body, `sgw_resource_utilization_process_memory_resident 1.48062208e+08`)
		requireCounter(t, body, `sgw_resource_utilization_pub_net_bytes_recv 6.97856009e+08`)
		requireCounter(t, body, `sgw_resource_utilization_pub_net_bytes_sent 1.66156024e+08`)
		requireGauge(t, body, `sgw_resource_utilization_system_memory_total 4.139307008e+09`)
		requireCounter(t, body, `sgw_resource_utilization_warn_count 1`)
		requireGauge(t, body, `sgw_security_auth_failed_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_security_auth_success_count{database="travel-sample"} 3`)
		requireCounter(t, body, `sgw_security_num_access_errors{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_security_num_docs_rejected{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_security_total_auth_time{database="travel-sample"} 1.643028e+08`)
		requireCounter(t, body, `sgw_shared_bucket_import_import_count{database="travel-sample"} 10031`)
		requireCounter(t, body, `sgw_shared_bucket_import_import_error_count{database="travel-sample"} 0`)
		requireGauge(t, body, `sgw_shared_bucket_import_import_processing_time{database="travel-sample"} 1.44880768e+11`)
		require.Regexp(t, "sgw_scrape_duration_seconds \\d+\\.\\d+", body)
		requireGauge(t, body, `sgw_up 1`)
	})
}

func TestMetricsDelta(t *testing.T) {
	var collector = NewCollector(newFakeClient(t, "testdata/metrics3.json", nil))
	testCollector(t, collector, func(t *testing.T, status int, body string) {
		require.Equal(t, 200, status)

		requireCounter(t, body, `sgw_delta_sync_delta_cache_hit{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_delta_sync_delta_cache_miss{database="travel-sample"} 20`)
		requireCounter(t, body, `sgw_delta_sync_delta_pull_replication_count{database="travel-sample"} 2`)
		requireCounter(t, body, `sgw_delta_sync_delta_push_doc_count{database="travel-sample"} 0`)
		requireCounter(t, body, `sgw_delta_sync_deltas_requested{database="travel-sample"} 20`)
		requireCounter(t, body, `sgw_delta_sync_deltas_sent{database="travel-sample"} 20`)

		require.Regexp(t, "sgw_scrape_duration_seconds \\d+\\.\\d+", body)
		requireGauge(t, body, `sgw_up 1`)
	})
}

func TestMetricsError(t *testing.T) {
	var collector = NewCollector(newFakeClient(t, "", fmt.Errorf("fake error")))
	testCollector(t, collector, func(t *testing.T, status int, body string) {
		require.Equal(t, 200, status)
		requireGauge(t, body, `sgw_up 0`)
		require.Regexp(t, "sgw_scrape_duration_seconds \\d+\\.\\d+", body)
	})
}

func testCollector(t *testing.T, collector prometheus.Collector, checker func(t *testing.T, status int, body string)) {
	var registry = prometheus.NewRegistry()
	registry.MustRegister(collector)

	var srv = httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	require.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	checker(t, resp.StatusCode, string(body))
}

type fakeClient struct {
	metrics client.Metrics
	err     error
}

func (c *fakeClient) Expvar() (client.Metrics, error) {
	return c.metrics, c.err
}

func newFakeClient(t *testing.T, path string, err error) client.Client {
	if err != nil {
		return &fakeClient{err: err}
	}
	bts, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	var metrics client.Metrics
	require.NoError(t, json.Unmarshal(bts, &metrics))
	return &fakeClient{metrics: metrics}
}

func requireCounter(t *testing.T, body, metric string) {
	requireMetric(t, body, metric, "counter")
}

func requireGauge(t *testing.T, body, metric string) {
	requireMetric(t, body, metric, "gauge")
}

var metricNameExp = regexp.MustCompile("^(\\w+)\\{?.*$")

func requireMetric(t *testing.T, body, metric, typee string) {
	var name = metricNameExp.FindAllStringSubmatch(metric, -1)[0][1]
	require.Contains(t, body, fmt.Sprintf("# TYPE %s %s", name, typee))
	require.Contains(t, body, metric)
}
