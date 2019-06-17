# couchbase-sync-gateway-exporter

A [Prometheus][] exporter for monitoring [Couchbase Sync Gateway][sgw]. The Sync Gateway stats are reported in JSON format through a [REST endpoint](https://docs.couchbase.com/sync-gateway/2.5/admin-rest-api.html#/server/get__expvar). The stats that are exported to Prometheus can be visualized using [Grafana](https://grafana.com).

---

The figure below shows a typical deployment of a Couchbase Mobile cluster with the sync gateway exporter:

![deployment](docs/deployment.png)

It includes -

- A couchbase server cluster of 3 or more nodes
- Two or more sync gateway instances fronted by a load balancer
- A *couchbase-sync-gateway-exporter* running alongside each sync gateway instance. The exporter polls for stats over the Sync Gateway admin port that is only exposed on localhost. While not shown in the figure, The exporter is polled by Prometheus server and optionally, the stats can be visualized using Grafana. 
- Couchbase Lite enabled clients talking to sync gateway instances through the load balancer

The exporter uses Sync Gateway's admin port to gather metrics, and that port
binds only to `localhost`. On the other hand, this exporter binds by default
to `0.0.0.0`. You can change that using `--web.listen-address` flag. You can
also secure it using firewall/VPC.

That being said, the Prometheus instance should be configured to gather metrics from all
the exporter instances. You can achieve that by just listing all endpoints
there or by using [service discovery][sd-config]:

```yaml
rule_files:
- '/etc/prometheus/rules/*'

scrape_configs:
  - job_name: swg
    static_configs:
    - targets:
      - sgw01.foo.local:9421
      - sgw02.foo.local:9421
      - sgw03.foo.local:9421
```

That's pretty much it.

[Prometheus]: https://prometheus.io
[sgw]: https://www.couchbase.com/products/sync-gateway
[sd-config]: https://prometheus.io/docs/prometheus/latest/configuration/configuration/

---

## Building, testing

### Running the exporter binary

**Requirement**: [Go](https://golang.org) 1.12+.

```sh
go run main.go --help
```
---


### Running unit tests

```sh
make test
```

or

```sh
# cover runs the test task and also opens the default browser with the
# coverage report.
make cover
```

### Linting Go code

**Requirement**: [golangci-lint](https://github.com/golangci/golangci-lint).

```sh
make lint
```


---

## Deploying with docker

### Building a docker image of exporter

**Requirement**: Docker 18.09.2+ and docker-compose 1.23.2+
If you are working with containers, you can create a docker image of the Sync Gateway exporter using the `build.dockerfile`

```sh
docker build -f build.dockerfile  .
```

---
### Deploying with docker-compose

**Requirement**: Docker 18.09.2+ and docker-compose 1.23.2+

This project has a `docker-compose.yaml` file with a full test environment,
which was used during the development, but can also be used to deploy entire Couchbase Mobile stack on your local machine. 
The  Couchbase Server container is pre-installed with the "travel sample" bucket and is configured with a RBAC user corresponding to the Sync Gateway. 

```sh
docker-compose down

docker-compose up
```

You can also build the exporter image  and deploy in a single step using the following commands

```sh
make start

make stop
```

*NOTE* : Depending on how long it takes for Couchbase server to get deployed and initialized, it is possible that the server is not available when the the Sync Gateway attempts to connect. In this case, the Sync Gateway tries to reconnect to the server for a certain number of times and if its too long, it  gives up. In this case you will have to launch the sync gateway separately again once you have confirmed that the Couchbase Server is up

```sh
docker-compose up sgw
```

---

## Deploying on Kubernetes

**Requirement**: Kubernetes 1.12+ and on k3s v0.5.0+

There is a full example on how to set up a Couchbase Mobile cluster with monitoring in the
the [kubernetes/](/kubernetes) folder.

Note that this is just an example configuration, and you would probably
want to customize it.

---


## Generating the Grafana dashboard

**Requirement**: [jsonnet](https://jsonnet.org/)

This generates a `dashboard.json` file from a `dashboard.jsonnet` file, using
jsonnet and [grafonnet-lib](https://github.com/grafana/grafonnet-lib). 

First pull relevant submodule for jsonnet
```sh
git submodule update --init --rebase --remote --recursive
```

```sh
# grafana script generates the grafana dashboard and pushes the dashboard to
# a running grafana instance on localhost:3000. So be sure that your
# grafana instance is running
make grafana
```

or

```sh
# grafana-dev generates the grafana dashboard and also setup a fresh
# grafana instance on localhost:3000, expecting default username and password.
make grafana-dev
```
---

## Extending the exporter
Follow instructions in our [docs/develop](/docs/develop) folder.

## Cutting a new release

**Requirement**: [GoReleaser](https://goreleaser.com) v0.107.0+

- `git tag` the new release in the [SemVer](https://semver.org/) format;
- push the tag;
- run `goreleaser`.
