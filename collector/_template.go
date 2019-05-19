package collector

import (
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

func newFooCollector() *fooCollector {
	const subsystem = "foo"
	return &fooCollector{}
}

type fooCollector struct {
}

func (c *fooCollector) Describe(ch chan<- *prometheus.Desc) {

}

func (c *fooCollector) Collect(ch chan<- prometheus.Metric, metrics client.ResourceUtilization) {

}
