package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newBucketImportCollector() *bucketImportCollector {
	const subsystem = "shared_bucket_import"
	return &bucketImportCollector{
		importCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_count"),
			"import_count",
			perDbLabels,
			nil,
		),
		importCancelCAS: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_cancel_cas"),
			"import_cancel_cas",
			perDbLabels,
			nil,
		),
		importErrorCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_error_count"),
			"import_error_count",
			perDbLabels,
			nil,
		),
		importProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_processing_time"),
			"import_processing_time",
			perDbLabels,
			nil,
		),
		importHighSeq: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_high_seq"),
			"import_high_seq",
			perDbLabels,
			nil,
		),
		importPartitions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "import_partitions"),
			"import_partitions",
			perDbLabels,
			nil,
		),
	}
}

type bucketImportCollector struct {
	importCount          *prometheus.Desc
	importCancelCAS      *prometheus.Desc
	importErrorCount     *prometheus.Desc
	importProcessingTime *prometheus.Desc
	importHighSeq        *prometheus.Desc
	importPartitions     *prometheus.Desc
}

func (c *bucketImportCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.importCount
	ch <- c.importCancelCAS
	ch <- c.importErrorCount
	ch <- c.importProcessingTime
	ch <- c.importHighSeq
	ch <- c.importPartitions
}

// nolint: lll
func (c *bucketImportCollector) Collect(ch chan<- prometheus.Metric, name string, metrics client.SharedBucketImport) {
	ch <- prometheus.MustNewConstMetric(c.importCount, prometheus.CounterValue, metrics.ImportCount, name)
	ch <- prometheus.MustNewConstMetric(c.importCancelCAS, prometheus.CounterValue, metrics.ImportCancelCAS, name)
	ch <- prometheus.MustNewConstMetric(c.importErrorCount, prometheus.CounterValue, metrics.ImportErrorCount, name)
	ch <- prometheus.MustNewConstMetric(c.importProcessingTime, prometheus.GaugeValue, metrics.ImportProcessingTime, name)
	ch <- prometheus.MustNewConstMetric(c.importHighSeq, prometheus.GaugeValue, metrics.ImportHighSeq, name)
	ch <- prometheus.MustNewConstMetric(c.importPartitions, prometheus.GaugeValue, metrics.ImportPartitions, name)
}
