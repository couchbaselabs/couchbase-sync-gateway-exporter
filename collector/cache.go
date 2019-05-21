package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newCacheCollector() *cacheCollector {
	const subsystem = "cache"
	return &cacheCollector{
		chanCacheActiveRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_active_revs"),
			"chan_cache_active_revs,",
			perDbLabels,
			nil,
		),
		chanCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_hits"),
			"chan_cache_hits,",
			perDbLabels,
			nil,
		),
		chanCacheMaxEntries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_max_entries"),
			"chan_cache_max_entries,",
			perDbLabels,
			nil,
		),
		chanCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_misses"),
			"chan_cache_misses,",
			perDbLabels,
			nil,
		),
		chanCacheNumChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_num_channels"),
			"chan_cache_num_channels,",
			perDbLabels,
			nil,
		),
		chanCachePendingQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_pending_queries"),
			"chan_cache_pending_queries,",
			perDbLabels,
			nil,
		),
		chanCacheRemovalRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_removal_revs"),
			"chan_cache_removal_revs,",
			perDbLabels,
			nil,
		),
		chanCacheTombstoneRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_tombstone_revs"),
			"chan_cache_tombstone_revs,",
			perDbLabels,
			nil,
		),
		numSkippedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_skipped_seqs"),
			"num_skipped_seqs,",
			perDbLabels,
			nil,
		),
		revCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_cache_hits"),
			"rev_cache_hits,",
			perDbLabels,
			nil,
		),
		revCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_cache_misses"),
			"rev_cache_misses,",
			perDbLabels,
			nil,
		),
	}
}

type cacheCollector struct {
	chanCacheActiveRevs     *prometheus.Desc
	chanCacheHits           *prometheus.Desc
	chanCacheMaxEntries     *prometheus.Desc
	chanCacheMisses         *prometheus.Desc
	chanCacheNumChannels    *prometheus.Desc
	chanCachePendingQueries *prometheus.Desc
	chanCacheRemovalRevs    *prometheus.Desc
	chanCacheTombstoneRevs  *prometheus.Desc
	numSkippedSeqs          *prometheus.Desc
	revCacheHits            *prometheus.Desc
	revCacheMisses          *prometheus.Desc
}

func (c *cacheCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.chanCacheActiveRevs
	ch <- c.chanCacheHits
	ch <- c.chanCacheMaxEntries
	ch <- c.chanCacheMisses
	ch <- c.chanCacheNumChannels
	ch <- c.chanCachePendingQueries
	ch <- c.chanCacheRemovalRevs
	ch <- c.chanCacheTombstoneRevs
	ch <- c.numSkippedSeqs
	ch <- c.revCacheHits
	ch <- c.revCacheMisses
}

// nolint: lll
func (c *cacheCollector) Collect(ch chan<- prometheus.Metric, name string, cache client.Cache) {
	ch <- prometheus.MustNewConstMetric(c.chanCacheActiveRevs, prometheus.GaugeValue, float64(cache.ChanCacheActiveRevs), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheHits, prometheus.GaugeValue, float64(cache.ChanCacheHits), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheMaxEntries, prometheus.GaugeValue, float64(cache.ChanCacheMaxEntries), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheMisses, prometheus.GaugeValue, float64(cache.ChanCacheMisses), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheNumChannels, prometheus.GaugeValue, float64(cache.ChanCacheNumChannels), name)
	ch <- prometheus.MustNewConstMetric(c.chanCachePendingQueries, prometheus.GaugeValue, float64(cache.ChanCachePendingQueries), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheRemovalRevs, prometheus.GaugeValue, float64(cache.ChanCacheRemovalRevs), name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheTombstoneRevs, prometheus.GaugeValue, float64(cache.ChanCacheTombstoneRevs), name)
	ch <- prometheus.MustNewConstMetric(c.numSkippedSeqs, prometheus.GaugeValue, float64(cache.NumSkippedSeqs), name)
	ch <- prometheus.MustNewConstMetric(c.revCacheHits, prometheus.GaugeValue, float64(cache.RevCacheHits), name)
	ch <- prometheus.MustNewConstMetric(c.revCacheMisses, prometheus.GaugeValue, float64(cache.RevCacheMisses), name)
}
