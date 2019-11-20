package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newCacheCollector() *cacheCollector {
	const subsystem = "cache"
	return &cacheCollector{
		abandonedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "abandoned_seqs"),
			"abandoned_seqs,",
			perDbLabels,
			nil,
		),
		chanCacheActiveRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_active_revs"),
			"chan_cache_active_revs,",
			perDbLabels,
			nil,
		),
		chanCacheBypassCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_bypass_count"),
			"chan_cache_bypass_count,",
			perDbLabels,
			nil,
		),
		chanCacheChannelsAdded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_channels_added"),
			"chan_cache_channels_added,",
			perDbLabels,
			nil,
		),
		chanCacheChannelsEvictedInactive: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_channels_evicted_inactive"),
			"chan_cache_channels_evicted_inactive,",
			perDbLabels,
			nil,
		),
		chanCacheChannelsEvictedNRU: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_channels_evicted_nru"),
			"chan_cache_channels_evicted_nru,",
			perDbLabels,
			nil,
		),
		chanCacheCompactCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_compact_count"),
			"chan_cache_compact_count,",
			perDbLabels,
			nil,
		),
		chanCacheCompactTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "chan_cache_compact_time"),
			"chan_cache_compact_time,",
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
		highSeqCached: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "high_seq_cached"),
			"high_seq_cached,",
			perDbLabels,
			nil,
		),
		highSeqStable: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "high_seq_stable"),
			"high_seq_stable,",
			perDbLabels,
			nil,
		),
		numActiveChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_active_channels"),
			"num_active_channels,",
			perDbLabels,
			nil,
		),
		numSkippedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_skipped_seqs"),
			"num_skipped_seqs,",
			perDbLabels,
			nil,
		),
		pendingSeqLen: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "pending_seq_len"),
			"pending_seq_len,",
			perDbLabels,
			nil,
		),
		revCacheBypass: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rev_cache_bypass"),
			"rev_cache_bypass,",
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
		skippedSeqLen: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "skipped_seq_len"),
			"skipped_seq_len,",
			perDbLabels,
			nil,
		),
	}
}

type cacheCollector struct {
	abandonedSeqs                    *prometheus.Desc
	chanCacheActiveRevs              *prometheus.Desc
	chanCacheBypassCount             *prometheus.Desc
	chanCacheChannelsAdded           *prometheus.Desc
	chanCacheChannelsEvictedInactive *prometheus.Desc
	chanCacheChannelsEvictedNRU      *prometheus.Desc
	chanCacheCompactCount            *prometheus.Desc
	chanCacheCompactTime             *prometheus.Desc
	chanCacheHits                    *prometheus.Desc
	chanCacheMaxEntries              *prometheus.Desc
	chanCacheMisses                  *prometheus.Desc
	chanCacheNumChannels             *prometheus.Desc
	chanCachePendingQueries          *prometheus.Desc
	chanCacheRemovalRevs             *prometheus.Desc
	chanCacheTombstoneRevs           *prometheus.Desc
	highSeqCached                    *prometheus.Desc
	highSeqStable                    *prometheus.Desc
	numActiveChannels                *prometheus.Desc
	numSkippedSeqs                   *prometheus.Desc
	pendingSeqLen                    *prometheus.Desc
	revCacheBypass                   *prometheus.Desc
	revCacheHits                     *prometheus.Desc
	revCacheMisses                   *prometheus.Desc
	skippedSeqLen                    *prometheus.Desc
}

func (c *cacheCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.abandonedSeqs
	ch <- c.chanCacheActiveRevs
	ch <- c.chanCacheBypassCount
	ch <- c.chanCacheChannelsAdded
	ch <- c.chanCacheChannelsEvictedInactive
	ch <- c.chanCacheChannelsEvictedNRU
	ch <- c.chanCacheCompactCount
	ch <- c.chanCacheCompactTime
	ch <- c.chanCacheHits
	ch <- c.chanCacheMaxEntries
	ch <- c.chanCacheMisses
	ch <- c.chanCacheNumChannels
	ch <- c.chanCachePendingQueries
	ch <- c.chanCacheRemovalRevs
	ch <- c.chanCacheTombstoneRevs
	ch <- c.highSeqCached
	ch <- c.highSeqStable
	ch <- c.numActiveChannels
	ch <- c.numSkippedSeqs
	ch <- c.pendingSeqLen
	ch <- c.revCacheBypass
	ch <- c.revCacheHits
	ch <- c.revCacheMisses
	ch <- c.skippedSeqLen
}

// nolint: lll
func (c *cacheCollector) Collect(ch chan<- prometheus.Metric, name string, cache client.Cache) {
	ch <- prometheus.MustNewConstMetric(c.abandonedSeqs, prometheus.CounterValue, cache.AbandonedSeqs, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheActiveRevs, prometheus.GaugeValue, cache.ChanCacheActiveRevs, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheBypassCount, prometheus.CounterValue, cache.ChanCacheBypassCount, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheChannelsAdded, prometheus.CounterValue, cache.ChanCacheChannelsAdded, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheChannelsEvictedInactive, prometheus.CounterValue, cache.ChanCacheChannelsEvictedInactive, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheChannelsEvictedNRU, prometheus.CounterValue, cache.ChanCacheChannelsEvictedNRU, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheCompactCount, prometheus.CounterValue, cache.ChanCacheCompactCount, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheCompactTime, prometheus.CounterValue, cache.ChanCacheCompactTime, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheHits, prometheus.CounterValue, cache.ChanCacheHits, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheMaxEntries, prometheus.GaugeValue, cache.ChanCacheMaxEntries, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheMisses, prometheus.CounterValue, cache.ChanCacheMisses, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheNumChannels, prometheus.GaugeValue, cache.ChanCacheNumChannels, name)
	ch <- prometheus.MustNewConstMetric(c.chanCachePendingQueries, prometheus.GaugeValue, cache.ChanCachePendingQueries, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheRemovalRevs, prometheus.GaugeValue, cache.ChanCacheRemovalRevs, name)
	ch <- prometheus.MustNewConstMetric(c.chanCacheTombstoneRevs, prometheus.GaugeValue, cache.ChanCacheTombstoneRevs, name)
	ch <- prometheus.MustNewConstMetric(c.highSeqCached, prometheus.CounterValue, cache.HighSeqCached, name)
	ch <- prometheus.MustNewConstMetric(c.highSeqStable, prometheus.CounterValue, cache.HighSeqStable, name)
	ch <- prometheus.MustNewConstMetric(c.numActiveChannels, prometheus.GaugeValue, cache.NumActiveChannels, name)
	ch <- prometheus.MustNewConstMetric(c.numSkippedSeqs, prometheus.CounterValue, cache.NumSkippedSeqs, name)
	ch <- prometheus.MustNewConstMetric(c.pendingSeqLen, prometheus.GaugeValue, cache.PendingSeqLen, name)
	ch <- prometheus.MustNewConstMetric(c.revCacheBypass, prometheus.GaugeValue, cache.RevCacheBypass, name)
	ch <- prometheus.MustNewConstMetric(c.revCacheHits, prometheus.CounterValue, cache.RevCacheHits, name)
	ch <- prometheus.MustNewConstMetric(c.revCacheMisses, prometheus.CounterValue, cache.RevCacheMisses, name)
	ch <- prometheus.MustNewConstMetric(c.skippedSeqLen, prometheus.GaugeValue, cache.SkippedSeqLen, name)
}
