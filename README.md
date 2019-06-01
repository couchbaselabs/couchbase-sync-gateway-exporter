# couchbase-sync-gateway-exporter

A [Prometheus][] exporter for [Couchbase Sync Gateway][sgw].

---

An usual deployment would look similar to the following:

![deploymeny](docs/deployment.png)

In that image we have:

- a couchbase server cluster
- several sync gateway instances
- one couchbase-sync-gateway-exporter running for each sync gateway instance
- a load balancer
- clients talking to sync gateway instances through the load balancer

The exporter uses Sync Gateway's admin port to gather metrics, and that port
binds only to `localhost`. On the other hand, this exporter binds by default
to `0.0.0.0`. You can change that using `--web.listen-address` flag. You can
also secure it using firewall/VPC.

That being said, the Prometheus instance should then gather metrics from all
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

## Building, testing...

### Running the exporter locally

**Requirement**: [Go](https://golang.org) 1.12+.

```sh
go run main.go --help
```

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

### Generating the Grafana dashboard

**Requirement**: [jsonnet](https://jsonnet.org/)

This generates a `dashboard.json` file from a `dashboard.jsonnet` file, using
jsonnet and [grafonnet-lib](https://github.com/grafana/grafonnet-lib).

```sh
make grafana
```

or

```sh
# grafana-dev generates the grafana dashboard and also setup a fresh
# grafana instance on localhost:3000, expecting default username and password.
make grafana-dev
```

---

## Running with docker-compose

**Requirement**: Docker 18.09.2+ and docker-compose 1.23.2+

This project has a `docker-compose.yaml` file with a full test environment,
which was used during the development, but can also be used to see locally
how everything works.

All tasks are available as `make` tasks:

```sh
make start
make stop
```

---

## Running on Kubernetes

**Requirement**: Kubernetes 1.12+ and on k3s v0.5.0+

There is a full example using Couchbase Operator and Prometheus Operator on
the [docs/kubebernetes](/docs/kubernetes) folder.

Note that this is just an example configuration, and you would probably
want to customize it.
