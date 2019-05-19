package collector

import (
	"sync"
	"time"

	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type sgwCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc

	globalCollector *globalCollector

	// db cache
	chanCacheActiveRevs    *prometheus.Desc
	chanCacheHits          *prometheus.Desc
	chanCacheMaxEntries    *prometheus.Desc
	chanCacheMisses        *prometheus.Desc
	chanCacheNumChannels   *prometheus.Desc
	chanCacheRemovalRevs   *prometheus.Desc
	chanCacheTombstoneRevs *prometheus.Desc
	numSkippedSeqs         *prometheus.Desc
	revCacheHits           *prometheus.Desc
	revCacheMisses         *prometheus.Desc

	// db replication pull
	attachmentPullBytes         *prometheus.Desc
	attachmentPullCount         *prometheus.Desc
	maxPending                  *prometheus.Desc
	numPullReplActiveContinuous *prometheus.Desc
	numPullReplActiveOneShot    *prometheus.Desc
	numPullReplCaughtUp         *prometheus.Desc
	numPullReplSinceZero        *prometheus.Desc
	numPullReplTotalContinuous  *prometheus.Desc
	numPullReplTotalOneShot     *prometheus.Desc
	requestChangesCount         *prometheus.Desc
	requestChangesTime          *prometheus.Desc
	revProcessingTime           *prometheus.Desc
	revSendCount                *prometheus.Desc
	revSendLatency              *prometheus.Desc

	// db replication push
	attachmentPushBytes *prometheus.Desc
	attachmentPushCount *prometheus.Desc
	conflictWriteCount  *prometheus.Desc
	docPushCount        *prometheus.Desc
	proposeChangeCount  *prometheus.Desc
	proposeChangeTime   *prometheus.Desc
	syncFunctionCount   *prometheus.Desc
	syncFunctionTime    *prometheus.Desc
	writeProcessingTime *prometheus.Desc
}

const namespace = "sgw"

// NewCollector tasks collector
func NewCollector(client client.Client) prometheus.Collector {
	const (
		subsystem        = ""
		cacheSubsystem   = "cache"
		repPullSubsystem = "replication_pull"
		repPushSubsystem = "replication_push"
	)

	var perDbLabels = []string{"database"}

	// nolint: lll
	return &sgwCollector{
		client: client,

		// default metrics
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "up"),
			"SGW admin API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
		globalCollector: newGlobalCollector(),

		//
		//
		// db cache metrics
		//
		//
		chanCacheActiveRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_active_revs"),
			"chan_cache_active_revs,",
			perDbLabels,
			nil,
		),
		chanCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_hits"),
			"chan_cache_hits,",
			perDbLabels,
			nil,
		),
		chanCacheMaxEntries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_max_entries"),
			"chan_cache_max_entries,",
			perDbLabels,
			nil,
		),
		chanCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_misses"),
			"chan_cache_misses,",
			perDbLabels,
			nil,
		),
		chanCacheNumChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_num_channels"),
			"chan_cache_num_channels,",
			perDbLabels,
			nil,
		),
		chanCacheRemovalRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_removal_revs"),
			"chan_cache_removal_revs,",
			perDbLabels,
			nil,
		),
		chanCacheTombstoneRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_tombstone_revs"),
			"chan_cache_tombstone_revs,",
			perDbLabels,
			nil,
		),
		numSkippedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "num_skipped_seqs"),
			"num_skipped_seqs,",
			perDbLabels,
			nil,
		),
		revCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "rev_cache_hits"),
			"rev_cache_hits,",
			perDbLabels,
			nil,
		),
		revCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "rev_cache_misses"),
			"rev_cache_misses,",
			perDbLabels,
			nil,
		),

		//
		//
		// db replication pull metrics
		//
		//
		attachmentPullBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "attachment_pull_bytes"),
			"attachment_pull_bytes",
			perDbLabels,
			nil,
		),
		attachmentPullCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "attachment_pull_count"),
			"attachment_pull_count",
			perDbLabels,
			nil,
		),
		maxPending: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "max_pending"),
			"max_pending",
			perDbLabels,
			nil,
		),
		numPullReplActiveContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_active_continuous"),
			"num_pull_repl_active_continuous",
			perDbLabels,
			nil,
		),
		numPullReplActiveOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_active_one_shot"),
			"num_pull_repl_active_one_shot",
			perDbLabels,
			nil,
		),
		numPullReplCaughtUp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_caught_up"),
			"num_pull_repl_caught_up",
			perDbLabels,
			nil,
		),
		numPullReplSinceZero: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_since_zero"),
			"num_pull_repl_since_zero",
			perDbLabels,
			nil,
		),
		numPullReplTotalContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_total_continuous"),
			"num_pull_repl_total_continuous",
			perDbLabels,
			nil,
		),
		numPullReplTotalOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_total_one_shot"),
			"num_pull_repl_total_one_shot",
			perDbLabels,
			nil,
		),
		requestChangesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "request_changes_count"),
			"request_changes_count",
			perDbLabels,
			nil,
		),
		requestChangesTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "request_changes_time"),
			"request_changes_time",
			perDbLabels,
			nil,
		),
		revProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_processing_time"),
			"rev_processing_time",
			perDbLabels,
			nil,
		),
		revSendCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_send_count"),
			"rev_send_count",
			perDbLabels,
			nil,
		),
		revSendLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_send_latency"),
			"rev_send_latency",
			perDbLabels,
			nil,
		),

		//
		//
		// db replication push metrics
		//
		//
		attachmentPushBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "attachment_push_bytes"),
			"attachment_push_bytes",
			perDbLabels,
			nil,
		),
		attachmentPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "attachment_push_count"),
			"attachment_push_count",
			perDbLabels,
			nil,
		),
		conflictWriteCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "conflict_write_count"),
			"conflict_write_count",
			perDbLabels,
			nil,
		),
		docPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "doc_push_count"),
			"doc_push_count",
			perDbLabels,
			nil,
		),
		proposeChangeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "propose_change_count"),
			"propose_change_count",
			perDbLabels,
			nil,
		),
		proposeChangeTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "propose_change_time"),
			"propose_change_time",
			perDbLabels,
			nil,
		),
		syncFunctionCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "sync_function_count"),
			"sync_function_count",
			perDbLabels,
			nil,
		),
		syncFunctionTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "sync_function_time"),
			"sync_function_time",
			perDbLabels,
			nil,
		),
		writeProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "write_processing_time"),
			"write_processing_time",
			perDbLabels,
			nil,
		),
	}
}

// Describe all metrics
func (c *sgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration

	c.globalCollector.Describe(ch)

	// cache
	ch <- c.chanCacheActiveRevs
	ch <- c.chanCacheHits
	ch <- c.chanCacheMaxEntries
	ch <- c.chanCacheMisses
	ch <- c.chanCacheNumChannels
	ch <- c.chanCacheRemovalRevs
	ch <- c.chanCacheTombstoneRevs
	ch <- c.numSkippedSeqs
	ch <- c.revCacheHits
	ch <- c.revCacheMisses

	// replication pull
	ch <- c.attachmentPullBytes
	ch <- c.attachmentPullCount
	ch <- c.maxPending
	ch <- c.numPullReplActiveContinuous
	ch <- c.numPullReplActiveOneShot
	ch <- c.numPullReplCaughtUp
	ch <- c.numPullReplSinceZero
	ch <- c.numPullReplTotalContinuous
	ch <- c.numPullReplTotalOneShot
	ch <- c.requestChangesCount
	ch <- c.requestChangesTime
	ch <- c.revProcessingTime
	ch <- c.revSendCount
	ch <- c.revSendLatency

	// db replication push
	ch <- c.attachmentPushBytes
	ch <- c.attachmentPushCount
	ch <- c.conflictWriteCount
	ch <- c.docPushCount
	ch <- c.proposeChangeCount
	ch <- c.proposeChangeTime
	ch <- c.syncFunctionCount
	ch <- c.syncFunctionTime
	ch <- c.writeProcessingTime
}

// Collect all metrics
func (c *sgwCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting sgw metrics...")
	result, err := c.client.Expvar()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape sgw")
		return
	}

	var metrics = result.Syncgateway

	// global resource utilization metrics
	log.Debug("collecting global resource utilization metrics")
	c.globalCollector.Collect(ch, metrics.Global.ResourceUtilization)

	// per-db metrics
	for name, db := range metrics.PerDb {
		log.Debugf("collecting cache metrics for db %s", name)
		var cache = db.Cache
		ch <- prometheus.MustNewConstMetric(c.chanCacheActiveRevs, prometheus.GaugeValue, float64(cache.ChanCacheActiveRevs), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheHits, prometheus.GaugeValue, float64(cache.ChanCacheHits), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheMaxEntries, prometheus.GaugeValue, float64(cache.ChanCacheMaxEntries), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheMisses, prometheus.GaugeValue, float64(cache.ChanCacheMisses), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheNumChannels, prometheus.GaugeValue, float64(cache.ChanCacheNumChannels), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheRemovalRevs, prometheus.GaugeValue, float64(cache.ChanCacheRemovalRevs), name)
		ch <- prometheus.MustNewConstMetric(c.chanCacheTombstoneRevs, prometheus.GaugeValue, float64(cache.ChanCacheTombstoneRevs), name)
		ch <- prometheus.MustNewConstMetric(c.numSkippedSeqs, prometheus.GaugeValue, float64(cache.NumSkippedSeqs), name)
		ch <- prometheus.MustNewConstMetric(c.revCacheHits, prometheus.GaugeValue, float64(cache.RevCacheHits), name)
		ch <- prometheus.MustNewConstMetric(c.revCacheMisses, prometheus.GaugeValue, float64(cache.RevCacheMisses), name)

		log.Debugf("collecting replication pull metrics for db %s", name)
		var pull = db.CblReplicationPull
		ch <- prometheus.MustNewConstMetric(c.attachmentPullBytes, prometheus.GaugeValue, float64(pull.AttachmentPullBytes), name)
		ch <- prometheus.MustNewConstMetric(c.attachmentPullCount, prometheus.GaugeValue, float64(pull.AttachmentPullCount), name)
		ch <- prometheus.MustNewConstMetric(c.maxPending, prometheus.GaugeValue, float64(pull.MaxPending), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplActiveContinuous, prometheus.GaugeValue, float64(pull.NumPullReplActiveContinuous), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplActiveOneShot, prometheus.GaugeValue, float64(pull.NumPullReplActiveOneShot), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplCaughtUp, prometheus.GaugeValue, float64(pull.NumPullReplCaughtUp), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplSinceZero, prometheus.GaugeValue, float64(pull.NumPullReplSinceZero), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplTotalContinuous, prometheus.GaugeValue, float64(pull.NumPullReplTotalContinuous), name)
		ch <- prometheus.MustNewConstMetric(c.numPullReplTotalOneShot, prometheus.GaugeValue, float64(pull.NumPullReplTotalOneShot), name)
		ch <- prometheus.MustNewConstMetric(c.requestChangesCount, prometheus.GaugeValue, float64(pull.RequestChangesCount), name)
		ch <- prometheus.MustNewConstMetric(c.requestChangesTime, prometheus.GaugeValue, float64(pull.RequestChangesTime), name)
		ch <- prometheus.MustNewConstMetric(c.revProcessingTime, prometheus.GaugeValue, float64(pull.RevProcessingTime), name)
		ch <- prometheus.MustNewConstMetric(c.revSendCount, prometheus.GaugeValue, float64(pull.RevSendCount), name)
		ch <- prometheus.MustNewConstMetric(c.revSendLatency, prometheus.GaugeValue, float64(pull.RevSendLatency), name)

		log.Debugf("collecting replication push metrics for db %s", name)
		var push = db.CblReplicationPush
		ch <- prometheus.MustNewConstMetric(c.attachmentPushBytes, prometheus.GaugeValue, float64(push.AttachmentPushBytes), name)
		ch <- prometheus.MustNewConstMetric(c.attachmentPushCount, prometheus.GaugeValue, float64(push.AttachmentPushCount), name)
		ch <- prometheus.MustNewConstMetric(c.conflictWriteCount, prometheus.GaugeValue, float64(push.ConflictWriteCount), name)
		ch <- prometheus.MustNewConstMetric(c.docPushCount, prometheus.GaugeValue, float64(push.DocPushCount), name)
		ch <- prometheus.MustNewConstMetric(c.proposeChangeCount, prometheus.GaugeValue, float64(push.ProposeChangeCount), name)
		ch <- prometheus.MustNewConstMetric(c.proposeChangeTime, prometheus.GaugeValue, float64(push.ProposeChangeTime), name)
		ch <- prometheus.MustNewConstMetric(c.syncFunctionCount, prometheus.GaugeValue, float64(push.SyncFunctionCount), name)
		ch <- prometheus.MustNewConstMetric(c.syncFunctionTime, prometheus.GaugeValue, float64(push.SyncFunctionTime), name)
		ch <- prometheus.MustNewConstMetric(c.writeProcessingTime, prometheus.GaugeValue, float64(push.WriteProcessingTime), name)
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	// nolint: lll
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
