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
}

// NewCollector tasks collector
func NewCollector(client client.Client) prometheus.Collector {
	const namespace = "couchbase"
	const subsystem = "sgw"
	const globalResourceSubsystem = "resource_utilization"
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
	}
}

// Describe all metrics
func (c *sgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
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

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	// nolint: lll
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
