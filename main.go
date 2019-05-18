package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/client"
	"github.com/couchbaselabs/couchbase-sync-gateway-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// nolint: gochecknoglobals,lll
var (
	version       = "dev"
	app           = kingpin.New("couchbase-sync-gateway-exporter", "exports couchbase's sync-gateway metrics in the prometheus format")
	listenAddress = app.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default("127.0.0.1:9421").String()
	metricsPath   = app.Flag("web.telemetry-path", "Path under which to expose metrics").Default("/metrics").String()
	sgwURL        = app.Flag("sgw.url", "Couchbase URL to scrape").Default("http://localhost:4985").String()
	debug         = app.Flag("debug", "Show debug logs").Bool()
)

func main() {
	app.Version(version)
	app.HelpFlag.Short('h')
	log.AddFlags(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.Infof("starting couchbase-sync-gateway-exporter %s...\n", version)

	var client = client.New(*sgwURL)

	prometheus.MustRegister(collector.NewCollector(client))
	http.Handle(*metricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`
			<html>
			<head><title>Couchbase Sync Gateway Exporter</title></head>
			<body>
				<h1>Couchbase Sync Gateway Exporter</h1>
				<p><a href="`+*metricsPath+`">Metrics</a></p>
			</body>
			</html>
			`)
	})

	log.Infof("server listening on %s", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
