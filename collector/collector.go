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
}

const (
	namespace               = "couchbase"
	subsystem               = "sgw"
	globalResourceSubsystem = "resource_utilization"
	cacheSubsystem          = "cache"
)

// NewCollector tasks collector
func NewCollector(client client.Client) prometheus.Collector {
	// nolint: lll
	return &sgwCollector{
		client: client,
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
	}
}

// Describe all metrics
func (c *sgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
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
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	// nolint: lll
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
