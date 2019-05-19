package collector

import (
	"sync"
	"time"

	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const namespace = "sgw"

var perDbLabels = []string{"database"}

// NewCollector collector
func NewCollector(client client.Client) prometheus.Collector {
	const subsystem = ""

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
		globalCollector:   newGlobalCollector(),
		cacheCollector:    newCacheCollector(),
		pullCollector:     newPullCollector(),
		pushCollector:     newPushCollector(),
		databaseCollector: newDatabaseCollector(),
	}
}

type sgwCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc

	globalCollector   *globalCollector
	cacheCollector    *cacheCollector
	pullCollector     *pullCollector
	pushCollector     *pushCollector
	databaseCollector *databaseCollector
}

// Describe all metrics
func (c *sgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration

	c.globalCollector.Describe(ch)
	c.cacheCollector.Describe(ch)
	c.pullCollector.Describe(ch)
	c.pushCollector.Describe(ch)
	c.databaseCollector.Describe(ch)
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
		c.cacheCollector.Collect(ch, name, db.Cache)

		log.Debugf("collecting replication pull metrics for db %s", name)
		c.pullCollector.Collect(ch, name, db.CblReplicationPull)

		log.Debugf("collecting replication push metrics for db %s", name)
		c.pushCollector.Collect(ch, name, db.CblReplicationPush)

		log.Debugf("collecting database metrics for db %s", name)
		c.databaseCollector.Collect(ch, name, db.Database)
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
