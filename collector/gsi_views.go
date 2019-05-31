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
		principalsCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "principals_count"),
			"principals_count",
			perDbLabels,
			nil,
		),
		resyncCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "resync_count"),
			"resync_count",
			perDbLabels,
			nil,
		),
		sequencesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sequences_count"),
			"sequences_count",
			perDbLabels,
			nil,
		),
		sessionsCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sessions_count"),
			"sessions_count",
			perDbLabels,
			nil,
		),
		tombstonesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tombstones_count"),
			"tombstones_count",
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
	principalsCount *prometheus.Desc
	resyncCount     *prometheus.Desc
	sequencesCount  *prometheus.Desc
	sessionsCount   *prometheus.Desc
	tombstonesCount *prometheus.Desc
}

func (c *gsiViewsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.accessCount
	ch <- c.roleAccessCount
	ch <- c.channelsCount
	ch <- c.allDocsCount
	ch <- c.principalsCount
	ch <- c.resyncCount
	ch <- c.sequencesCount
	ch <- c.sessionsCount
	ch <- c.tombstonesCount
}

// nolint: lll
func (c *gsiViewsCollector) Collect(ch chan<- prometheus.Metric, name string, gsi client.GsiViews) {
	ch <- prometheus.MustNewConstMetric(c.accessCount, prometheus.CounterValue, gsi.AccessCount, name)
	ch <- prometheus.MustNewConstMetric(c.roleAccessCount, prometheus.CounterValue, gsi.RoleAccessCount, name)
	ch <- prometheus.MustNewConstMetric(c.channelsCount, prometheus.CounterValue, gsi.ChannelsCount, name)
	ch <- prometheus.MustNewConstMetric(c.allDocsCount, prometheus.CounterValue, gsi.AllDocsCount, name)
	ch <- prometheus.MustNewConstMetric(c.principalsCount, prometheus.CounterValue, gsi.PrincipalsCount, name)
	ch <- prometheus.MustNewConstMetric(c.resyncCount, prometheus.CounterValue, gsi.ResyncCount, name)
	ch <- prometheus.MustNewConstMetric(c.sequencesCount, prometheus.CounterValue, gsi.SequencesCount, name)
	ch <- prometheus.MustNewConstMetric(c.sessionsCount, prometheus.CounterValue, gsi.SessionsCount, name)
	ch <- prometheus.MustNewConstMetric(c.tombstonesCount, prometheus.CounterValue, gsi.TombstonesCount, name)
}
