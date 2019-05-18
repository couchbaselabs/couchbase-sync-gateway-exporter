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

	// global resource utilization
	adminNetBytesRecv            *prometheus.Desc
	adminNetBytesSent            *prometheus.Desc
	errorCount                   *prometheus.Desc
	goMemstatsHeapalloc          *prometheus.Desc
	goMemstatsHeapidle           *prometheus.Desc
	goMemstatsHeapinuse          *prometheus.Desc
	goMemstatsHeapreleased       *prometheus.Desc
	goMemstatsPausetotalns       *prometheus.Desc
	goMemstatsStackinuse         *prometheus.Desc
	goMemstatsStacksys           *prometheus.Desc
	goMemstatsSys                *prometheus.Desc
	goroutinesHighWatermark      *prometheus.Desc
	numGoroutines                *prometheus.Desc
	processCPUPercentUtilization *prometheus.Desc
	processMemoryResident        *prometheus.Desc
	pubNetBytesRecv              *prometheus.Desc
	pubNetBytesSent              *prometheus.Desc
	systemMemoryTotal            *prometheus.Desc
	warnCount                    *prometheus.Desc

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

const (
	namespace               = "sgw"
	subsystem               = ""
	globalResourceSubsystem = "resource_utilization"
	cacheSubsystem          = "cache"
	repPullSubsystem        = "replication_pull"
	repPushSubsystem        = "replication_push"
)

// NewCollector tasks collector
func NewCollector(client client.Client) prometheus.Collector {
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

		//
		//
		// global resource utilization metrics
		//
		//
		adminNetBytesRecv: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "admin_net_bytes_recv"),
			"admin_net_bytes_recv",
			nil,
			nil,
		),
		adminNetBytesSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "admin_net_bytes_sent"),
			"admin_net_bytes_sent",
			nil,
			nil,
		),
		errorCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "error_count"),
			"error_count",
			nil,
			nil,
		),
		goMemstatsHeapalloc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_heapalloc"),
			"go_memstats_heapalloc",
			nil,
			nil,
		),
		goMemstatsHeapidle: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_heapidle"),
			"go_memstats_heapidle",
			nil,
			nil,
		),
		goMemstatsHeapinuse: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_heapinuse"),
			"go_memstats_heapinuse",
			nil,
			nil,
		),
		goMemstatsHeapreleased: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_heapreleased"),
			"go_memstats_heapreleased",
			nil,
			nil,
		),
		goMemstatsPausetotalns: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_pausetotalns"),
			"go_memstats_pausetotalns",
			nil,
			nil,
		),
		goMemstatsStackinuse: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_stackinuse"),
			"go_memstats_stackinuse",
			nil,
			nil,
		),
		goMemstatsStacksys: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_stacksys"),
			"go_memstats_stacksys",
			nil,
			nil,
		),
		goMemstatsSys: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "go_memstats_sys"),
			"go_memstats_sys",
			nil,
			nil,
		),
		goroutinesHighWatermark: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "goroutines_high_watermark"),
			"goroutines_high_watermark",
			nil,
			nil,
		),
		numGoroutines: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "num_goroutines"),
			"num_goroutines",
			nil,
			nil,
		),
		processCPUPercentUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "process_cpu_percent_utilization"),
			"process_cpu_percent_utilization",
			nil,
			nil,
		),
		processMemoryResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "process_memory_resident"),
			"process_memory_resident",
			nil,
			nil,
		),
		pubNetBytesRecv: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "pub_net_bytes_recv"),
			"pub_net_bytes_recv",
			nil,
			nil,
		),
		pubNetBytesSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "pub_net_bytes_sent"),
			"pub_net_bytes_sent",
			nil,
			nil,
		),
		systemMemoryTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "system_memory_total"),
			"system_memory_total",
			nil,
			nil,
		),
		warnCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, globalResourceSubsystem, "warn_count"),
			"warn_count",
			nil,
			nil,
		),

		//
		//
		// db cache metrics
		//
		//
		chanCacheActiveRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_active_revs"),
			"chan_cache_active_revs,",
			[]string{"database"},
			nil,
		),
		chanCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_hits"),
			"chan_cache_hits,",
			[]string{"database"},
			nil,
		),
		chanCacheMaxEntries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_max_entries"),
			"chan_cache_max_entries,",
			[]string{"database"},
			nil,
		),
		chanCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_misses"),
			"chan_cache_misses,",
			[]string{"database"},
			nil,
		),
		chanCacheNumChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_num_channels"),
			"chan_cache_num_channels,",
			[]string{"database"},
			nil,
		),
		chanCacheRemovalRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_removal_revs"),
			"chan_cache_removal_revs,",
			[]string{"database"},
			nil,
		),
		chanCacheTombstoneRevs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "chan_cache_tombstone_revs"),
			"chan_cache_tombstone_revs,",
			[]string{"database"},
			nil,
		),
		numSkippedSeqs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "num_skipped_seqs"),
			"num_skipped_seqs,",
			[]string{"database"},
			nil,
		),
		revCacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "rev_cache_hits"),
			"rev_cache_hits,",
			[]string{"database"},
			nil,
		),
		revCacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cacheSubsystem, "rev_cache_misses"),
			"rev_cache_misses,",
			[]string{"database"},
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
			[]string{"database"},
			nil,
		),
		attachmentPullCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "attachment_pull_count"),
			"attachment_pull_count",
			[]string{"database"},
			nil,
		),
		maxPending: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "max_pending"),
			"max_pending",
			[]string{"database"},
			nil,
		),
		numPullReplActiveContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_active_continuous"),
			"num_pull_repl_active_continuous",
			[]string{"database"},
			nil,
		),
		numPullReplActiveOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_active_one_shot"),
			"num_pull_repl_active_one_shot",
			[]string{"database"},
			nil,
		),
		numPullReplCaughtUp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_caught_up"),
			"num_pull_repl_caught_up",
			[]string{"database"},
			nil,
		),
		numPullReplSinceZero: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_since_zero"),
			"num_pull_repl_since_zero",
			[]string{"database"},
			nil,
		),
		numPullReplTotalContinuous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_total_continuous"),
			"num_pull_repl_total_continuous",
			[]string{"database"},
			nil,
		),
		numPullReplTotalOneShot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "num_pull_repl_total_one_shot"),
			"num_pull_repl_total_one_shot",
			[]string{"database"},
			nil,
		),
		requestChangesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "request_changes_count"),
			"request_changes_count",
			[]string{"database"},
			nil,
		),
		requestChangesTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "request_changes_time"),
			"request_changes_time",
			[]string{"database"},
			nil,
		),
		revProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_processing_time"),
			"rev_processing_time",
			[]string{"database"},
			nil,
		),
		revSendCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_send_count"),
			"rev_send_count",
			[]string{"database"},
			nil,
		),
		revSendLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPullSubsystem, "rev_send_latency"),
			"rev_send_latency",
			[]string{"database"},
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
			[]string{"database"},
			nil,
		),
		attachmentPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "attachment_push_count"),
			"attachment_push_count",
			[]string{"database"},
			nil,
		),
		conflictWriteCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "conflict_write_count"),
			"conflict_write_count",
			[]string{"database"},
			nil,
		),
		docPushCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "doc_push_count"),
			"doc_push_count",
			[]string{"database"},
			nil,
		),
		proposeChangeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "propose_change_count"),
			"propose_change_count",
			[]string{"database"},
			nil,
		),
		proposeChangeTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "propose_change_time"),
			"propose_change_time",
			[]string{"database"},
			nil,
		),
		syncFunctionCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "sync_function_count"),
			"sync_function_count",
			[]string{"database"},
			nil,
		),
		syncFunctionTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "sync_function_time"),
			"sync_function_time",
			[]string{"database"},
			nil,
		),
		writeProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repPushSubsystem, "write_processing_time"),
			"write_processing_time",
			[]string{"database"},
			nil,
		),
	}
}

// Describe all metrics
func (c *sgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration

	// global resource utilization
	ch <- c.adminNetBytesRecv
	ch <- c.adminNetBytesSent
	ch <- c.errorCount
	ch <- c.goMemstatsHeapalloc
	ch <- c.goMemstatsHeapidle
	ch <- c.goMemstatsHeapinuse
	ch <- c.goMemstatsHeapreleased
	ch <- c.goMemstatsPausetotalns
	ch <- c.goMemstatsStackinuse
	ch <- c.goMemstatsStacksys
	ch <- c.goMemstatsSys
	ch <- c.goroutinesHighWatermark
	ch <- c.numGoroutines
	ch <- c.processCPUPercentUtilization
	ch <- c.processMemoryResident
	ch <- c.pubNetBytesRecv
	ch <- c.pubNetBytesSent
	ch <- c.systemMemoryTotal
	ch <- c.warnCount

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
	var utilization = metrics.Global.ResourceUtilization
	ch <- prometheus.MustNewConstMetric(c.adminNetBytesRecv, prometheus.GaugeValue, float64(utilization.AdminNetBytesRecv))
	ch <- prometheus.MustNewConstMetric(c.adminNetBytesSent, prometheus.GaugeValue, float64(utilization.AdminNetBytesSent))
	ch <- prometheus.MustNewConstMetric(c.errorCount, prometheus.GaugeValue, float64(utilization.ErrorCount))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapalloc, prometheus.GaugeValue, float64(utilization.GoMemstatsHeapalloc))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapidle, prometheus.GaugeValue, float64(utilization.GoMemstatsHeapidle))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapinuse, prometheus.GaugeValue, float64(utilization.GoMemstatsHeapinuse))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapreleased, prometheus.GaugeValue, float64(utilization.GoMemstatsHeapreleased))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsPausetotalns, prometheus.GaugeValue, float64(utilization.GoMemstatsPausetotalns))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsStackinuse, prometheus.GaugeValue, float64(utilization.GoMemstatsStackinuse))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsStacksys, prometheus.GaugeValue, float64(utilization.GoMemstatsStacksys))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsSys, prometheus.GaugeValue, float64(utilization.GoMemstatsSys))
	ch <- prometheus.MustNewConstMetric(c.goroutinesHighWatermark, prometheus.GaugeValue, float64(utilization.GoroutinesHighWatermark))
	ch <- prometheus.MustNewConstMetric(c.numGoroutines, prometheus.GaugeValue, float64(utilization.NumGoroutines))
	ch <- prometheus.MustNewConstMetric(c.processCPUPercentUtilization, prometheus.GaugeValue, float64(utilization.ProcessCPUPercentUtilization))
	ch <- prometheus.MustNewConstMetric(c.processMemoryResident, prometheus.GaugeValue, float64(utilization.ProcessMemoryResident))
	ch <- prometheus.MustNewConstMetric(c.pubNetBytesRecv, prometheus.GaugeValue, float64(utilization.PubNetBytesRecv))
	ch <- prometheus.MustNewConstMetric(c.pubNetBytesSent, prometheus.GaugeValue, float64(utilization.PubNetBytesSent))
	ch <- prometheus.MustNewConstMetric(c.systemMemoryTotal, prometheus.GaugeValue, float64(utilization.SystemMemoryTotal))
	ch <- prometheus.MustNewConstMetric(c.warnCount, prometheus.GaugeValue, float64(utilization.WarnCount))

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
