package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newPushCollector() *pushCollector {
	const subsystem = "replication_push"
	return &pushCollector{
		attachmentPushBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "attachment_push_bytes"),
			"attachment_push_bytes",
			perDbLabels,
			nil,
		),
		attachmentPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "attachment_push_count"),
			"attachment_push_count",
			perDbLabels,
			nil,
		),
		conflictWriteCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "conflict_write_count"),
			"conflict_write_count",
			perDbLabels,
			nil,
		),
		docPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_push_count"),
			"doc_push_count",
			perDbLabels,
			nil,
		),
		proposeChangeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "propose_change_count"),
			"propose_change_count",
			perDbLabels,
			nil,
		),
		proposeChangeTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "propose_change_time"),
			"propose_change_time",
			perDbLabels,
			nil,
		),
		syncFunctionCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sync_function_count"),
			"sync_function_count",
			perDbLabels,
			nil,
		),
		syncFunctionTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sync_function_time"),
			"sync_function_time",
			perDbLabels,
			nil,
		),
		writeProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_processing_time"),
			"write_processing_time",
			perDbLabels,
			nil,
		),
	}
}

type pushCollector struct {
	attachmentPushBytes *prometheus.Desc
	attachmentPushCount *prometheus.Desc
	conflictWriteCount  *prometheus.Desc
	docPushCount        *prometheus.Desc
	proposeChangeCount  *prometheus.Desc
	proposeChangeTime   *prometheus.Desc
	syncFunctionCount   *prometheus.Desc
	syncFunctionTime    *prometheus.Desc
	writeProcessingTime *prometheus.Desc
}

func (c *pushCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.attachmentPushBytes
	ch <- c.attachmentPushCount
	ch <- c.conflictWriteCount
	ch <- c.docPushCount
	ch <- c.proposeChangeCount
	ch <- c.proposeChangeTime
	ch <- c.syncFunctionCount
	ch <- c.syncFunctionTime
	ch <- c.writeProcessingTime
}

// nolint: lll
func (c *pushCollector) Collect(ch chan<- prometheus.Metric, name string, push client.CblReplicationPush) {
	ch <- prometheus.MustNewConstMetric(c.attachmentPushBytes, prometheus.GaugeValue, float64(push.AttachmentPushBytes), name)
	ch <- prometheus.MustNewConstMetric(c.attachmentPushCount, prometheus.GaugeValue, float64(push.AttachmentPushCount), name)
	ch <- prometheus.MustNewConstMetric(c.conflictWriteCount, prometheus.GaugeValue, float64(push.ConflictWriteCount), name)
	ch <- prometheus.MustNewConstMetric(c.docPushCount, prometheus.GaugeValue, float64(push.DocPushCount), name)
	ch <- prometheus.MustNewConstMetric(c.proposeChangeCount, prometheus.GaugeValue, float64(push.ProposeChangeCount), name)
	ch <- prometheus.MustNewConstMetric(c.proposeChangeTime, prometheus.GaugeValue, float64(push.ProposeChangeTime), name)
	ch <- prometheus.MustNewConstMetric(c.syncFunctionCount, prometheus.GaugeValue, float64(push.SyncFunctionCount), name)
	ch <- prometheus.MustNewConstMetric(c.syncFunctionTime, prometheus.GaugeValue, float64(push.SyncFunctionTime), name)
	ch <- prometheus.MustNewConstMetric(c.writeProcessingTime, prometheus.GaugeValue, float64(push.WriteProcessingTime), name)
}
