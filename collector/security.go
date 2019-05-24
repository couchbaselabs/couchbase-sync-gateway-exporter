package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newSecurityCollector() *securityCollector {
	const subsystem = "security"
	return &securityCollector{
		authFailedCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "auth_failed_count"),
			"auth_failed_count",
			perDbLabels,
			nil,
		),
		authSuccessCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "auth_success_count"),
			"auth_success_count",
			perDbLabels,
			nil,
		),
		numAccessErrors: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_access_errors"),
			"num_access_errors",
			perDbLabels,
			nil,
		),
		numDocsRejected: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_docs_rejected"),
			"num_docs_rejected",
			perDbLabels,
			nil,
		),
		totalAuthTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_auth_time"),
			"total_auth_time",
			perDbLabels,
			nil,
		),
	}
}

type securityCollector struct {
	authFailedCount  *prometheus.Desc
	authSuccessCount *prometheus.Desc
	numAccessErrors  *prometheus.Desc
	numDocsRejected  *prometheus.Desc
	totalAuthTime    *prometheus.Desc
}

func (c *securityCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.authFailedCount
	ch <- c.authSuccessCount
	ch <- c.numAccessErrors
	ch <- c.numDocsRejected
	ch <- c.totalAuthTime
}

// nolint: lll
func (c *securityCollector) Collect(ch chan<- prometheus.Metric, name string, sec client.Security) {
	ch <- prometheus.MustNewConstMetric(c.authFailedCount, prometheus.CounterValue, float64(sec.AuthFailedCount), name)
	ch <- prometheus.MustNewConstMetric(c.authSuccessCount, prometheus.CounterValue, float64(sec.AuthSuccessCount), name)
	ch <- prometheus.MustNewConstMetric(c.numAccessErrors, prometheus.CounterValue, float64(sec.NumAccessErrors), name)
	ch <- prometheus.MustNewConstMetric(c.numDocsRejected, prometheus.CounterValue, float64(sec.NumDocsRejected), name)
	ch <- prometheus.MustNewConstMetric(c.totalAuthTime, prometheus.GaugeValue, float64(sec.TotalAuthTime), name)
}
