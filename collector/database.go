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
		docWritesXattrBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_writes_xattr_bytes"),
			"doc_writes_xattr_bytes",
			perDbLabels,
			nil,
		),
		highSeqFeed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "high_seq_feed"),
			"high_seq_feed",
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
		numTombstonesCompacted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_tombstones_compacted"),
			"num_tombstones_compacted",
			perDbLabels,
			nil,
		),
		sequenceAssignedCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_assigned_count"),
			"sequence_assigned_count",
			perDbLabels,
			nil,
		),
		sequenceGetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_get_count"),
			"sequence_get_count",
			perDbLabels,
			nil,
		),
		sequenceIncrCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequence_incr_count"),
			"sequence_incr_count",
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
	docWritesXattrBytes     *prometheus.Desc
	highSeqFeed             *prometheus.Desc
	numDocReadsBlip         *prometheus.Desc
	numDocReadsRest         *prometheus.Desc
	numDocWrites            *prometheus.Desc
	numReplicationsActive   *prometheus.Desc
	numReplicationsTotal    *prometheus.Desc
	numTombstonesCompacted  *prometheus.Desc
	sequenceAssignedCount   *prometheus.Desc
	sequenceGetCount        *prometheus.Desc
	sequenceIncrCount       *prometheus.Desc
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
	ch <- c.docWritesXattrBytes
	ch <- c.highSeqFeed
	ch <- c.numDocReadsBlip
	ch <- c.numDocReadsRest
	ch <- c.numDocWrites
	ch <- c.numReplicationsActive
	ch <- c.numReplicationsTotal
	ch <- c.numTombstonesCompacted
	ch <- c.sequenceAssignedCount
	ch <- c.sequenceGetCount
	ch <- c.sequenceIncrCount
	ch <- c.sequenceReleasedCount
	ch <- c.sequenceReservedCount
	ch <- c.warnChannelsPerDocCount
	ch <- c.warnGrantsPerDocCount
	ch <- c.warnXattrSizeCount
}

// nolint: lll
func (c *databaseCollector) Collect(ch chan<- prometheus.Metric, name string, db client.Database) {
	ch <- prometheus.MustNewConstMetric(c.abandonedSeqs, prometheus.CounterValue, db.AbandonedSeqs, name)
	ch <- prometheus.MustNewConstMetric(c.crc32CMatchCount, prometheus.GaugeValue, db.Crc32CMatchCount, name)
	ch <- prometheus.MustNewConstMetric(c.dcpCachingCount, prometheus.GaugeValue, db.DcpCachingCount, name)
	ch <- prometheus.MustNewConstMetric(c.dcpCachingTime, prometheus.GaugeValue, db.DcpCachingTime, name)
	ch <- prometheus.MustNewConstMetric(c.dcpReceivedCount, prometheus.GaugeValue, db.DcpReceivedCount, name)
	ch <- prometheus.MustNewConstMetric(c.dcpReceivedTime, prometheus.GaugeValue, db.DcpReceivedTime, name)
	ch <- prometheus.MustNewConstMetric(c.docReadsBytesBlip, prometheus.CounterValue, db.DocReadsBytesBlip, name)
	ch <- prometheus.MustNewConstMetric(c.docWritesBytes, prometheus.CounterValue, db.DocWritesBytes, name)
	ch <- prometheus.MustNewConstMetric(c.docWritesBytesBlip, prometheus.CounterValue, db.DocWritesBytesBlip, name)
	ch <- prometheus.MustNewConstMetric(c.docWritesXattrBytes, prometheus.CounterValue, db.DocWritesXattrBytes, name)
	ch <- prometheus.MustNewConstMetric(c.highSeqFeed, prometheus.CounterValue, db.HighSeqFeed, name)
	ch <- prometheus.MustNewConstMetric(c.numDocReadsBlip, prometheus.CounterValue, db.NumDocReadsBlip, name)
	ch <- prometheus.MustNewConstMetric(c.numDocReadsRest, prometheus.CounterValue, db.NumDocReadsRest, name)
	ch <- prometheus.MustNewConstMetric(c.numDocWrites, prometheus.CounterValue, db.NumDocWrites, name)
	ch <- prometheus.MustNewConstMetric(c.numReplicationsActive, prometheus.GaugeValue, db.NumReplicationsActive, name)
	ch <- prometheus.MustNewConstMetric(c.numReplicationsTotal, prometheus.CounterValue, db.NumReplicationsTotal, name)
	ch <- prometheus.MustNewConstMetric(c.numTombstonesCompacted, prometheus.CounterValue, db.NumTombstonesCompacted, name)
	ch <- prometheus.MustNewConstMetric(c.sequenceAssignedCount, prometheus.CounterValue, db.SequenceAssignedCount, name)
	ch <- prometheus.MustNewConstMetric(c.sequenceGetCount, prometheus.CounterValue, db.SequenceGetCount, name)
	ch <- prometheus.MustNewConstMetric(c.sequenceIncrCount, prometheus.CounterValue, db.SequenceIncrCount, name)
	ch <- prometheus.MustNewConstMetric(c.sequenceReleasedCount, prometheus.CounterValue, db.SequenceReleasedCount, name)
	ch <- prometheus.MustNewConstMetric(c.sequenceReservedCount, prometheus.CounterValue, db.SequenceReservedCount, name)
	ch <- prometheus.MustNewConstMetric(c.warnChannelsPerDocCount, prometheus.CounterValue, db.WarnChannelsPerDocCount, name)
	ch <- prometheus.MustNewConstMetric(c.warnGrantsPerDocCount, prometheus.CounterValue, db.WarnGrantsPerDocCount, name)
	ch <- prometheus.MustNewConstMetric(c.warnXattrSizeCount, prometheus.CounterValue, db.WarnXattrSizeCount, name)
}
