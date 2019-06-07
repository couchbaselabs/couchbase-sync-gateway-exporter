# Deploying to Kubernetes
Follow these instructions to deploy sync gateway exporter with Prometheus and Grafana to a Couchbase Mobile Kubernetes cluster. You will have to customize the sample config files that are used in the instructions below to suit your deployment. 

## Install Helm
We will use [Helm](https://helm.sh) for deploying the Prometheus Operator. If you already have your environment discussed in the official docs for Helm, you can skip this step

```sh
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
helm init --service-account tiller --upgrade
```

## Install Couchbase Server Cluster
If you already have a Couchbase server cluster, you can skip this step.

Follow the instructions in the [Couchbase official guide](https://docs.couchbase.com/operator/1.2/overview.html)to deploy Couchbase server cluster using the Couchbase Autonomous Operator. 

Steps below is a quick start guide to getting Couchbase server cluster going using Helm. It is extracted from the [official docs](https://docs.couchbase.com/operator/current/helm-discussed in the official docs-guide.html).

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

## Setup RBAC user on Couchbase Server 
Next, you will have to create an appropriate RBAC user for the Sync Gateway to connect to the Couchbase server cluster.
The default sync gateway config file used in the repo expects the Sync Gateway RBAC user with credentials of "admin" and "password". Follow instructions [here](https://docs.couchbase.com/sync-gateway/2.5/getting-started.html#creating-an-rbac-user) to configure the RBAC user with specified credentials.


## Install Prometheus Operator
If you already have Prometheus Operator deployed in your environment, you can skip this step. Details on Prometheus Operator are available in these [official docs](https://github.com/helm/charts/tree/master/stable/prometheus-operator)

```sh
helm install --namespace prometheus --name prom stable/prometheus-operator -f kubernetes/prometheus/values.yaml
```

This should show you all pods running:

```sh
kubectl -n prometheus get pods
```

## Setup Sync Gateway + Exporter
In this step you will set up a pod with 2 containers, corresponding to the Sync Gateway and the exporter respectively
The steps below are a simplified version of the environment that is discussed in the [official docs](https://docs.couchbase.com/sync-gateway/2.5/kubernetes/deploy-cluster.html).

Create the config secret:

```sh
kubectl create -n couchbase secret generic sgw-config --from-file=./kubernetes/sgw-config.json
kubectl apply -f kubernetes/sgw
```

This should launch a single Sync Gateway instance with the exporter as a sidecar:

```sh
kubectl -n couchbase get pods -l app=sync-gateway
```

## Checking setup

```sh
kubectl -n prometheus port-forward svc/prom-prometheus-operator-prometheus 9090:9090

open http://localhost:9090
```

You should see something like this:

![](/docs/kubernetes/screen-1sgw.png)

Then we can scale the sync gateway with:

```sh
kubectl scale -n couchbase deploy/sync-gateway --replicas 2
```

And refresh that page, so you can see something like this:


![](/docs/kubernetes/screen-2sgw.png)

## Grafana
We use Grafana for stats visualization. 

First, lets port-forward grafana to our local environment:

```sh
kubectl -n prometheus port-forward svc/prom-grafana 3000:80

open http://localhost:3000
```

Username and password are `admin` and `admin`. It will have a
set of dashboards already there.

To install our dashboard, you can run:

```sh
make grafana-dev
```

And then you should be able to find the Sync Gateway dashboard on
http://localhost:3000:

![](/docs/kubernetes/dash.png)

By default it will show metrics for all Sync Gateway instances, but you can
of course filter them:

![](/docs/kubernetes/choose-instances.png)

Those values are queried on dashboard load, so you may need to reload it as
new instances come online.

By default, it will use the pod name to differentiate between Sync Gateway
instances. You can revert back to the prometheus-operator default of using
the pod IP by commenting
[these lines](https://github.com/caarlos0/couchbase-sync-gateway-exporter/blob/master/kubernetes/sgw/servicemonitor.yaml#L18-L25)
and re-applying that file.
