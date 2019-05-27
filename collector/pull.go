package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newPullCollector() *pullCollector {
	const subsystem = "replication_pull"
	return &pullCollector{
		attachmentPullBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "attachment_pull_bytes"),
			"attachment_pull_bytes",
			perDbLabels,
			nil,
		),
		attachmentPullCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "attachment_pull_count"),
			"attachment_pull_count",
			perDbLabels,
			nil,
		),
		maxPending: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "max_pending"),
			"max_pending",
			perDbLabels,
			nil,
		),
		numPullReplActiveContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_active_continuous"),
			"num_pull_repl_active_continuous",
			perDbLabels,
			nil,
		),
		numPullReplActiveOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_active_one_shot"),
			"num_pull_repl_active_one_shot",
			perDbLabels,
			nil,
		),
		numPullReplCaughtUp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_caught_up"),
			"num_pull_repl_caught_up",
			perDbLabels,
			nil,
		),
		numPullReplSinceZero: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_since_zero"),
			"num_pull_repl_since_zero",
			perDbLabels,
			nil,
		),
		numPullReplTotalContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_total_continuous"),
			"num_pull_repl_total_continuous",
			perDbLabels,
			nil,
		),
		numPullReplTotalOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pull_repl_total_one_shot"),
			"num_pull_repl_total_one_shot",
			perDbLabels,
			nil,
		),
		requestChangesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "request_changes_count"),
			"request_changes_count",
			perDbLabels,
			nil,
		),
		requestChangesTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "request_changes_time"),
			"request_changes_time",
			perDbLabels,
			nil,
		),
		revProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_processing_time"),
			"rev_processing_time",
			perDbLabels,
			nil,
		),
		revSendCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_send_count"),
			"rev_send_count",
			perDbLabels,
			nil,
		),
		revSendLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_send_latency"),
			"rev_send_latency",
			perDbLabels,
			nil,
		),
	}
}

type pullCollector struct {
	attachmentPullBytes         *prometheus.Desc
	attachmentPullCount         *prometheus.Desc
	maxPending                  *prometheus.Desc
	numPullReplActiveContinuous *prometheus.Desc
	numPullReplActiveOneShot    *prometheus.Desc
	numPullReplCaughtUp         *prometheus.Desc
	numPullReplSinceZero        *prometheus.Desc
	numPullReplTotalContinuous  *prometheus.Desc
	numPullReplTotalOneShot     *prometheus.Desc
	requestChangesCount         *prometheus.Desc
	requestChangesTime          *prometheus.Desc
	revProcessingTime           *prometheus.Desc
	revSendCount                *prometheus.Desc
	revSendLatency              *prometheus.Desc
}

func (c *pullCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.attachmentPullBytes
	ch <- c.attachmentPullCount
	ch <- c.maxPending
	ch <- c.numPullReplActiveContinuous
	ch <- c.numPullReplActiveOneShot
	ch <- c.numPullReplCaughtUp
	ch <- c.numPullReplSinceZero
	ch <- c.numPullReplTotalContinuous
	ch <- c.numPullReplTotalOneShot
	ch <- c.requestChangesCount
	ch <- c.requestChangesTime
	ch <- c.revProcessingTime
	ch <- c.revSendCount
	ch <- c.revSendLatency
}

// nolint: lll
func (c *pullCollector) Collect(ch chan<- prometheus.Metric, name string, pull client.CblReplicationPull) {
	ch <- prometheus.MustNewConstMetric(c.attachmentPullBytes, prometheus.CounterValue, float64(pull.AttachmentPullBytes), name)
	ch <- prometheus.MustNewConstMetric(c.attachmentPullCount, prometheus.CounterValue, float64(pull.AttachmentPullCount), name)
	ch <- prometheus.MustNewConstMetric(c.maxPending, prometheus.GaugeValue, float64(pull.MaxPending), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplActiveContinuous, prometheus.GaugeValue, float64(pull.NumPullReplActiveContinuous), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplActiveOneShot, prometheus.GaugeValue, float64(pull.NumPullReplActiveOneShot), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplCaughtUp, prometheus.GaugeValue, float64(pull.NumPullReplCaughtUp), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplSinceZero, prometheus.CounterValue, float64(pull.NumPullReplSinceZero), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplTotalContinuous, prometheus.GaugeValue, float64(pull.NumPullReplTotalContinuous), name)
	ch <- prometheus.MustNewConstMetric(c.numPullReplTotalOneShot, prometheus.GaugeValue, float64(pull.NumPullReplTotalOneShot), name)
	ch <- prometheus.MustNewConstMetric(c.requestChangesCount, prometheus.CounterValue, float64(pull.RequestChangesCount), name)
	ch <- prometheus.MustNewConstMetric(c.requestChangesTime, prometheus.CounterValue, float64(pull.RequestChangesTime), name)
	ch <- prometheus.MustNewConstMetric(c.revProcessingTime, prometheus.GaugeValue, float64(pull.RevProcessingTime), name)
	ch <- prometheus.MustNewConstMetric(c.revSendCount, prometheus.CounterValue, float64(pull.RevSendCount), name)
	ch <- prometheus.MustNewConstMetric(c.revSendLatency, prometheus.CounterValue, float64(pull.RevSendLatency), name)
}
