package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newDatabaseCollector() *databaseCollector {
	const subsystem = "database"
	return &databaseCollector{
		abandonedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "abandoned_seqs"),
			"abandoned_seqs",
			perDbLabels,
			nil,
		),
		crc32CMatchCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "crc32c_match_count"),
			"crc32c_match_count",
			perDbLabels,
			nil,
		),
		dcpCachingCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "dcp_caching_count"),
			"dcp_caching_count",
			perDbLabels,
			nil,
		),
		dcpCachingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "dcp_caching_time"),
			"dcp_caching_time",
			perDbLabels,
			nil,
		),
		dcpReceivedCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "dcp_received_count"),
			"dcp_received_count",
			perDbLabels,
			nil,
		),
		dcpReceivedTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "dcp_received_time"),
			"dcp_received_time",
			perDbLabels,
			nil,
		),
		docReadsBytesBlip: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_reads_bytes_blip"),
			"doc_reads_bytes_blip",
			perDbLabels,
			nil,
		),
		docWritesBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_writes_bytes"),
			"doc_writes_bytes",
			perDbLabels,
			nil,
		),
		docWritesBytesBlip: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_writes_bytes_blip"),
			"doc_writes_bytes_blip",
			perDbLabels,
			nil,
		),
		numDocReadsBlip: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_doc_reads_blip"),
			"num_doc_reads_blip",
			perDbLabels,
			nil,
		),
		numDocReadsRest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_doc_reads_rest"),
			"num_doc_reads_rest",
			perDbLabels,
			nil,
		),
		numDocWrites: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_doc_writes"),
			"num_doc_writes",
			perDbLabels,
			nil,
		),
		numReplicationsActive: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_replications_active"),
			"num_replications_active",
			perDbLabels,
			nil,
		),
		numReplicationsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_replications_total"),
			"num_replications_total",
			perDbLabels,
			nil,
		),
		sequenceGetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_get_count"),
			"sequence_get_count",
			perDbLabels,
			nil,
		),
		sequenceReleasedCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_released_count"),
			"sequence_released_count",
			perDbLabels,
			nil,
		),
		sequenceReservedCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_reserved_count"),
			"sequence_reserved_count",
			perDbLabels,
			nil,
		),
		warnChannelsPerDocCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "warn_channels_per_doc_count"),
			"warn_channels_per_doc_count",
			perDbLabels,
			nil,
		),
		warnGrantsPerDocCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "warn_grants_per_doc_count"),
			"warn_grants_per_doc_count",
			perDbLabels,
			nil,
		),
		warnXattrSizeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "warn_xattr_size_count"),
			"warn_xattr_size_count",
			perDbLabels,
			nil,
		),
	}
}

type databaseCollector struct {
	abandonedSeqs           *prometheus.Desc
	crc32CMatchCount        *prometheus.Desc
	dcpCachingCount         *prometheus.Desc
	dcpCachingTime          *prometheus.Desc
	dcpReceivedCount        *prometheus.Desc
	dcpReceivedTime         *prometheus.Desc
	docReadsBytesBlip       *prometheus.Desc
	docWritesBytes          *prometheus.Desc
	docWritesBytesBlip      *prometheus.Desc
	numDocReadsBlip         *prometheus.Desc
	numDocReadsRest         *prometheus.Desc
	numDocWrites            *prometheus.Desc
	numReplicationsActive   *prometheus.Desc
	numReplicationsTotal    *prometheus.Desc
	sequenceGetCount        *prometheus.Desc
	sequenceReleasedCount   *prometheus.Desc
	sequenceReservedCount   *prometheus.Desc
	warnChannelsPerDocCount *prometheus.Desc
	warnGrantsPerDocCount   *prometheus.Desc
	warnXattrSizeCount      *prometheus.Desc
}

func (c *databaseCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.abandonedSeqs
	ch <- c.crc32CMatchCount
	ch <- c.dcpCachingCount
	ch <- c.dcpCachingTime
	ch <- c.dcpReceivedCount
	ch <- c.dcpReceivedTime
	ch <- c.docReadsBytesBlip
	ch <- c.docWritesBytes
	ch <- c.docWritesBytesBlip
	ch <- c.numDocReadsBlip
	ch <- c.numDocReadsRest
	ch <- c.numDocWrites
	ch <- c.numReplicationsActive
	ch <- c.numReplicationsTotal
	ch <- c.sequenceGetCount
	ch <- c.sequenceReleasedCount
	ch <- c.sequenceReservedCount
	ch <- c.warnChannelsPerDocCount
	ch <- c.warnGrantsPerDocCount
	ch <- c.warnXattrSizeCount
}

func (c *databaseCollector) Collect(ch chan<- prometheus.Metric, name string, db client.Database) {
	ch <- prometheus.MustNewConstMetric(c.abandonedSeqs, prometheus.GaugeValue, float64(db.AbandonedSeqs), name)
	ch <- prometheus.MustNewConstMetric(c.crc32CMatchCount, prometheus.GaugeValue, float64(db.Crc32CMatchCount), name)
	ch <- prometheus.MustNewConstMetric(c.dcpCachingCount, prometheus.GaugeValue, float64(db.DcpCachingCount), name)
	ch <- prometheus.MustNewConstMetric(c.dcpCachingTime, prometheus.GaugeValue, float64(db.DcpCachingTime), name)
	ch <- prometheus.MustNewConstMetric(c.dcpReceivedCount, prometheus.GaugeValue, float64(db.DcpReceivedCount), name)
	ch <- prometheus.MustNewConstMetric(c.dcpReceivedTime, prometheus.GaugeValue, float64(db.DcpReceivedTime), name)
	ch <- prometheus.MustNewConstMetric(c.docReadsBytesBlip, prometheus.GaugeValue, float64(db.DocReadsBytesBlip), name)
	ch <- prometheus.MustNewConstMetric(c.docWritesBytes, prometheus.GaugeValue, float64(db.DocWritesBytes), name)
	ch <- prometheus.MustNewConstMetric(c.docWritesBytesBlip, prometheus.GaugeValue, float64(db.DocWritesBytesBlip), name)
	ch <- prometheus.MustNewConstMetric(c.numDocReadsBlip, prometheus.GaugeValue, float64(db.NumDocReadsBlip), name)
	ch <- prometheus.MustNewConstMetric(c.numDocReadsRest, prometheus.GaugeValue, float64(db.NumDocReadsRest), name)
	ch <- prometheus.MustNewConstMetric(c.numDocWrites, prometheus.GaugeValue, float64(db.NumDocWrites), name)
	ch <- prometheus.MustNewConstMetric(c.numReplicationsActive, prometheus.GaugeValue, float64(db.NumReplicationsActive), name)
	ch <- prometheus.MustNewConstMetric(c.numReplicationsTotal, prometheus.GaugeValue, float64(db.NumReplicationsTotal), name)
	ch <- prometheus.MustNewConstMetric(c.sequenceGetCount, prometheus.GaugeValue, float64(db.SequenceGetCount), name)
	ch <- prometheus.MustNewConstMetric(c.sequenceReleasedCount, prometheus.GaugeValue, float64(db.SequenceReleasedCount), name)
	ch <- prometheus.MustNewConstMetric(c.sequenceReservedCount, prometheus.GaugeValue, float64(db.SequenceReservedCount), name)
	ch <- prometheus.MustNewConstMetric(c.warnChannelsPerDocCount, prometheus.GaugeValue, float64(db.WarnChannelsPerDocCount), name)
	ch <- prometheus.MustNewConstMetric(c.warnGrantsPerDocCount, prometheus.GaugeValue, float64(db.WarnGrantsPerDocCount), name)
	ch <- prometheus.MustNewConstMetric(c.warnXattrSizeCount, prometheus.GaugeValue, float64(db.WarnXattrSizeCount), name)
}
