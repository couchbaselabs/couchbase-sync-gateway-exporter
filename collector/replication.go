package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newReplicationCollector() *replicationCollector {
	const subsystem = "replication"
	var labels = []string{"replication"}
	return &replicationCollector{
		sgrDocsCheckedSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sgr_docs_checked_sent"),
			"sgr_docs_checked_sent",
			labels,
			nil,
		),
		sgrNumAttachmentBytesTransferred: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sgr_num_attachment_bytes_transferred"),
			"sgr_num_attachment_bytes_transferred",
			labels,
			nil,
		),
		sgrNumAttachmentsTransferred: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sgr_num_attachments_transferred"),
			"sgr_num_attachments_transferred",
			labels,
			nil,
		),
		sgrNumDocsFailedToPush: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sgr_num_docs_failed_to_push"),
			"sgr_num_docs_failed_to_push",
			labels,
			nil,
		),
		sgrNumDocsPushed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sgr_num_docs_pushed"),
			"sgr_num_docs_pushed",
			labels,
			nil,
		),
	}
}

type replicationCollector struct {
	sgrDocsCheckedSent               *prometheus.Desc
	sgrNumAttachmentBytesTransferred *prometheus.Desc
	sgrNumAttachmentsTransferred     *prometheus.Desc
	sgrNumDocsFailedToPush           *prometheus.Desc
	sgrNumDocsPushed                 *prometheus.Desc
}

func (c *replicationCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.sgrDocsCheckedSent
	ch <- c.sgrNumAttachmentBytesTransferred
	ch <- c.sgrNumAttachmentsTransferred
	ch <- c.sgrNumDocsFailedToPush
	ch <- c.sgrNumDocsPushed
}

// nolint: lll
func (c *replicationCollector) Collect(ch chan<- prometheus.Metric, name string, replication client.Replication) {
	ch <- prometheus.MustNewConstMetric(c.sgrDocsCheckedSent, prometheus.GaugeValue, float64(replication.SgrDocsCheckedSent), name)
	ch <- prometheus.MustNewConstMetric(c.sgrNumAttachmentBytesTransferred, prometheus.GaugeValue, float64(replication.SgrNumAttachmentBytesTransferred), name)
	ch <- prometheus.MustNewConstMetric(c.sgrNumAttachmentsTransferred, prometheus.GaugeValue, float64(replication.SgrNumAttachmentsTransferred), name)
	ch <- prometheus.MustNewConstMetric(c.sgrNumDocsFailedToPush, prometheus.GaugeValue, float64(replication.SgrNumDocsFailedToPush), name)
	ch <- prometheus.MustNewConstMetric(c.sgrNumDocsPushed, prometheus.GaugeValue, float64(replication.SgrNumDocsPushed), name)
}
