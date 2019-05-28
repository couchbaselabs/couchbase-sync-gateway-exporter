# Deploying to Kubernetes

## Install Helm

```sh
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
helm init --service-account tiller --upgrade
```

## Install Couchbase Operator

```sh
helm repo add couchbase https://couchbase-partners.github.io/helm-charts/

helm install --namespace couchbase --name couchbase-operator couchbase/couchbase-operator
helm install --namespace couchbase --name couchbase couchbase/couchbase-cluster -f kubernetes/couchbase/values.yaml
```

This should install the operator and launch a new Couchbase cluster.

It should create a Couchbase "cluster" with a single node, a `default`
bucket and credentials being `Administrator`/`password`.

You can port forward to it to see if everything is good (it may take a while):

```sh
kubectl port-forward -n couchbase svc/couchbase-couchbase-cluster-ui 8091:8091
open http://localhost:8091
```

## Install Prometheus Operator

```sh
helm install --namespace prometheus --name prom stable/prometheus-operator -f kubernetes/prometheus/values.yaml
```

This should show you all pods running:

```sh
kubectl -n prometheus get pods
```

## Setup Sync Gateway + Exporter

Create the config secret:

```sh
kubectl create -n couchbase secret generic sgw-config --from-file=./kubernetes/sgw-config.json
kubectl apply -f kubernetes/sgw
```

This should launch 2 SGW instances each one with the exporter as a sidecar:

```sh
kubectl get pods -l app sync-gateway
```

