package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newGsiViewsCollector() *gsiViewsCollector {
	const subsystem = "gsi_views"
	return &gsiViewsCollector{
		accessCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "access_count"),
			"access_count",
			perDbLabels,
			nil,
		),
		roleAccessCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "roleAccess_count"),
			"roleAccess_count",
			perDbLabels,
			nil,
		),
		channelsCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "channels_count"),
			"channels_count",
			perDbLabels,
			nil,
		),
		allDocsCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "allDocs_count"),
			"allDocs_count",
			perDbLabels,
			nil,
		),
	}
}

type gsiViewsCollector struct {
	accessCount     *prometheus.Desc
	roleAccessCount *prometheus.Desc
	channelsCount   *prometheus.Desc
	allDocsCount    *prometheus.Desc
}

func (c *gsiViewsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.accessCount
	ch <- c.roleAccessCount
	ch <- c.channelsCount
	ch <- c.allDocsCount
}

// nolint: lll
func (c *gsiViewsCollector) Collect(ch chan<- prometheus.Metric, name string, role client.GsiViews) {
	ch <- prometheus.MustNewConstMetric(c.accessCount, prometheus.CounterValue, float64(role.AccessCount), name)
	ch <- prometheus.MustNewConstMetric(c.roleAccessCount, prometheus.CounterValue, float64(role.RoleAccessCount), name)
	ch <- prometheus.MustNewConstMetric(c.channelsCount, prometheus.CounterValue, float64(role.ChannelsCount), name)
	ch <- prometheus.MustNewConstMetric(c.allDocsCount, prometheus.CounterValue, float64(role.AllDocsCount), name)
}
