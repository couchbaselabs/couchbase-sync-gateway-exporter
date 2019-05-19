package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newGlobalCollector() *globalCollector {
	const subsystem = "resource_utilization"
	return &globalCollector{
		adminNetBytesRecv: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "admin_net_bytes_recv"),
			"admin_net_bytes_recv",
			nil,
			nil,
		),
		adminNetBytesSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "admin_net_bytes_sent"),
			"admin_net_bytes_sent",
			nil,
			nil,
		),
		errorCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "error_count"),
			"error_count",
			nil,
			nil,
		),
		goMemstatsHeapalloc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_heapalloc"),
			"go_memstats_heapalloc",
			nil,
			nil,
		),
		goMemstatsHeapidle: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_heapidle"),
			"go_memstats_heapidle",
			nil,
			nil,
		),
		goMemstatsHeapinuse: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_heapinuse"),
			"go_memstats_heapinuse",
			nil,
			nil,
		),
		goMemstatsHeapreleased: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_heapreleased"),
			"go_memstats_heapreleased",
			nil,
			nil,
		),
		goMemstatsPausetotalns: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_pausetotalns"),
			"go_memstats_pausetotalns",
			nil,
			nil,
		),
		goMemstatsStackinuse: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_stackinuse"),
			"go_memstats_stackinuse",
			nil,
			nil,
		),
		goMemstatsStacksys: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_stacksys"),
			"go_memstats_stacksys",
			nil,
			nil,
		),
		goMemstatsSys: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "go_memstats_sys"),
			"go_memstats_sys",
			nil,
			nil,
		),
		goroutinesHighWatermark: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "goroutines_high_watermark"),
			"goroutines_high_watermark",
			nil,
			nil,
		),
		numGoroutines: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_goroutines"),
			"num_goroutines",
			nil,
			nil,
		),
		processCPUPercentUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "process_cpu_percent_utilization"),
			"process_cpu_percent_utilization",
			nil,
			nil,
		),
		processMemoryResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "process_memory_resident"),
			"process_memory_resident",
			nil,
			nil,
		),
		pubNetBytesRecv: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "pub_net_bytes_recv"),
			"pub_net_bytes_recv",
			nil,
			nil,
		),
		pubNetBytesSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "pub_net_bytes_sent"),
			"pub_net_bytes_sent",
			nil,
			nil,
		),
		systemMemoryTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "system_memory_total"),
			"system_memory_total",
			nil,
			nil,
		),
		warnCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "warn_count"),
			"warn_count",
			nil,
			nil,
		),
	}
}

type globalCollector struct {
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

func (c *globalCollector) Describe(ch chan<- *prometheus.Desc) {
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
}

func (c *globalCollector) Collect(ch chan<- prometheus.Metric, metrics client.ResourceUtilization) {
	ch <- prometheus.MustNewConstMetric(c.adminNetBytesRecv, prometheus.GaugeValue, float64(metrics.AdminNetBytesRecv))
	ch <- prometheus.MustNewConstMetric(c.adminNetBytesSent, prometheus.GaugeValue, float64(metrics.AdminNetBytesSent))
	ch <- prometheus.MustNewConstMetric(c.errorCount, prometheus.GaugeValue, float64(metrics.ErrorCount))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapalloc, prometheus.GaugeValue, float64(metrics.GoMemstatsHeapalloc))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapidle, prometheus.GaugeValue, float64(metrics.GoMemstatsHeapidle))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapinuse, prometheus.GaugeValue, float64(metrics.GoMemstatsHeapinuse))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsHeapreleased, prometheus.GaugeValue, float64(metrics.GoMemstatsHeapreleased))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsPausetotalns, prometheus.GaugeValue, float64(metrics.GoMemstatsPausetotalns))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsStackinuse, prometheus.GaugeValue, float64(metrics.GoMemstatsStackinuse))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsStacksys, prometheus.GaugeValue, float64(metrics.GoMemstatsStacksys))
	ch <- prometheus.MustNewConstMetric(c.goMemstatsSys, prometheus.GaugeValue, float64(metrics.GoMemstatsSys))
	ch <- prometheus.MustNewConstMetric(c.goroutinesHighWatermark, prometheus.GaugeValue, float64(metrics.GoroutinesHighWatermark))
	ch <- prometheus.MustNewConstMetric(c.numGoroutines, prometheus.GaugeValue, float64(metrics.NumGoroutines))
	ch <- prometheus.MustNewConstMetric(c.processCPUPercentUtilization, prometheus.GaugeValue, float64(metrics.ProcessCPUPercentUtilization))
	ch <- prometheus.MustNewConstMetric(c.processMemoryResident, prometheus.GaugeValue, float64(metrics.ProcessMemoryResident))
	ch <- prometheus.MustNewConstMetric(c.pubNetBytesRecv, prometheus.GaugeValue, float64(metrics.PubNetBytesRecv))
	ch <- prometheus.MustNewConstMetric(c.pubNetBytesSent, prometheus.GaugeValue, float64(metrics.PubNetBytesSent))
	ch <- prometheus.MustNewConstMetric(c.systemMemoryTotal, prometheus.GaugeValue, float64(metrics.SystemMemoryTotal))
	ch <- prometheus.MustNewConstMetric(c.warnCount, prometheus.GaugeValue, float64(metrics.WarnCount))
}
