package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newDeltaSyncCollector() *deltaSyncCollector {
	const subsystem = "delta_sync"
	return &deltaSyncCollector{
		deltaCacheHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "delta_cache_hit"),
			"delta_cache_hit",
			perDbLabels,
			nil,
		),
		deltaCacheMiss: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "delta_cache_miss"),
			"delta_cache_miss",
			perDbLabels,
			nil,
		),
		deltaPullReplicationCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "delta_pull_replication_count"),
			"delta_pull_replication_count",
			perDbLabels,
			nil,
		),
		deltaPushDocCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "delta_push_doc_count"),
			"delta_push_doc_count",
			perDbLabels,
			nil,
		),
		deltasRequested: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "deltas_requested"),
			"deltas_requested",
			perDbLabels,
			nil,
		),
		deltasSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "deltas_sent"),
			"deltas_sent",
			perDbLabels,
			nil,
		),
	}
}

type deltaSyncCollector struct {
	deltaCacheHit             *prometheus.Desc
	deltaCacheMiss            *prometheus.Desc
	deltaPullReplicationCount *prometheus.Desc
	deltaPushDocCount         *prometheus.Desc
	deltasRequested           *prometheus.Desc
	deltasSent                *prometheus.Desc
}

func (c *deltaSyncCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.deltaCacheHit
	ch <- c.deltaCacheMiss
	ch <- c.deltaPullReplicationCount
	ch <- c.deltaPushDocCount
	ch <- c.deltasRequested
	ch <- c.deltasSent
}

// nolint: lll
func (c *deltaSyncCollector) Collect(ch chan<- prometheus.Metric, name string, delta client.DeltaSync) {
	ch <- prometheus.MustNewConstMetric(c.deltaCacheHit, prometheus.CounterValue, float64(delta.DeltaCacheHit), name)
	ch <- prometheus.MustNewConstMetric(c.deltaCacheMiss, prometheus.CounterValue, float64(delta.DeltaCacheMiss), name)
	ch <- prometheus.MustNewConstMetric(c.deltaPullReplicationCount, prometheus.CounterValue, float64(delta.DeltaPullReplicationCount), name)
	ch <- prometheus.MustNewConstMetric(c.deltaPushDocCount, prometheus.CounterValue, float64(delta.DeltaPushDocCount), name)
	ch <- prometheus.MustNewConstMetric(c.deltasRequested, prometheus.CounterValue, float64(delta.DeltasRequested), name)
	ch <- prometheus.MustNewConstMetric(c.deltasSent, prometheus.CounterValue, float64(delta.DeltasSent), name)
}
