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
	}
}

type bucketImportCollector struct {
	importCount          *prometheus.Desc
	importErrorCount     *prometheus.Desc
	importProcessingTime *prometheus.Desc
}

func (c *bucketImportCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.importCount
	ch <- c.importErrorCount
	ch <- c.importProcessingTime
}

// nolint: lll
func (c *bucketImportCollector) Collect(ch chan<- prometheus.Metric, name string, metrics client.SharedBucketImport) {
	ch <- prometheus.MustNewConstMetric(c.importCount, prometheus.CounterValue, float64(metrics.ImportCount), name)
	ch <- prometheus.MustNewConstMetric(c.importErrorCount, prometheus.CounterValue, float64(metrics.ImportErrorCount), name)
	ch <- prometheus.MustNewConstMetric(c.importProcessingTime, prometheus.GaugeValue, float64(metrics.ImportProcessingTime), name)
}
