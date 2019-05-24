package collector

import (
	"sync"
	"time"

	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const namespace = "sgw"

// nolint: gochecknoglobals
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

		globalCollector:       newGlobalCollector(),
		cacheCollector:        newCacheCollector(),
		pullCollector:         newPullCollector(),
		pushCollector:         newPushCollector(),
		databaseCollector:     newDatabaseCollector(),
		gsiViewsCollector:     newGsiViewsCollector(),
		securityCollector:     newSecurityCollector(),
		bucketImportCollector: newBucketImportCollector(),
		replicationCollector:  newReplicationCollector(),
		deltaSyncCollector:    newDeltaSyncCollector(),
	}
}

type sgwCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc

	globalCollector       *globalCollector
	cacheCollector        *cacheCollector
	pullCollector         *pullCollector
	pushCollector         *pushCollector
	databaseCollector     *databaseCollector
	gsiViewsCollector     *gsiViewsCollector
	securityCollector     *securityCollector
	bucketImportCollector *bucketImportCollector
	replicationCollector  *replicationCollector
	deltaSyncCollector    *deltaSyncCollector
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
	c.gsiViewsCollector.Describe(ch)
	c.securityCollector.Describe(ch)
	c.bucketImportCollector.Describe(ch)
	c.replicationCollector.Describe(ch)
	c.deltaSyncCollector.Describe(ch)
}

// Collect all metrics
func (c *sgwCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
	}()
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

		log.Debugf("collecting gsi views metrics for db %s", name)
		c.gsiViewsCollector.Collect(ch, name, db.GsiViews)

		log.Debugf("collecting security metrics for db %s", name)
		c.securityCollector.Collect(ch, name, db.Security)

		log.Debugf("collecting shared bucket import metrics for db %s", name)
		c.bucketImportCollector.Collect(ch, name, db.SharedBucketImport)

		log.Debugf("collecting delta sync metrics for db %s", name)
		c.deltaSyncCollector.Collect(ch, name, db.DeltaSync)
	}

	// per-replication metrics
	for name, replication := range metrics.PerReplication {
		log.Debugf("collecting replication metrics for replication %s", name)
		c.replicationCollector.Collect(ch, name, replication)
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
}
