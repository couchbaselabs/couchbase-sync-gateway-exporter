## Extending exporter for new metrics

The exporter parses the JSON output of Sync Gateway [`/_expvar`](https://docs.couchbase.com/sync-gateway/2.5/admin-rest-api.html#/server/get__expvar) admin API endpoint and
transforms it to [Prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/).

To support a new metric that is exposed under `/_expvar`, we follow these steps :

1. Transform it from JSON to a native Go `struct` type

The parsing from JSON to a `struct` is made on the `client` package. You
can simply add your new field to the appropriate `struct` definition or create
a new one if necessary in `model.go` and that should be it.

We define all numbers as `float64` because that's the format that needs to be
exposed anyway. If a number is returned in `/_expvar` as an `int`, we would
still need to convert it to `float` at some point. Doing it when unmarshalling
the JSON saves the hassle and later usage is simpler.

2. Expose it via the exporter's `/metrics` endpoint.

Next, we need to expose it on `/metrics`. This step is done on the `collector`
package.
In the `collector` package you will find one `category.go` file for each "category"
of metric.
Once you find the appropriate one, we need to do the following:
1. Add the metric field in the `categoryCollector` `struct` (e.g.: `gsiViewsCollector`);
2. Add the metric field instantiation on the `newCategoryCollector`
implementation (e.g.: `newGsiViewsCollector`);
3. Add the metric to the `Describe` implementation;
4. finally, add the `Collect` implementation.

Step 4 requires attention regarding which kind of metric you are exporting
(if it's a `Counter` or `Gauge`). Other than that, this is all pretty simple.

3. Write test for new metric. 

All collector tests are written on `collector_test.go`. The same file has some test helper functions
to reduce code repetition, so it all boils out to something like:

```go
func TestMyNewMetric(t *testing.T) {
	var collector = NewCollector(newFakeClient(t, "testdata/mymetricfile.json", nil))
	testCollector(t, collector, func(t *testing.T, status int, body string) {
		requireSuccess(t, status, body)

		requireCounter(t, body, `sgw_my_expected_metric 1.3`)
	})
}
```

So, all you need to do is add, in this example, a
`collector/testdata/mymetricfile.json` file (which is an example output of
`/_expvar`) and add the `require` calls as you expect them to be.

If instead of adding a new metric, you need to add a new set of metrics, a
new object inside the `per_database` object for example, not much changes:

1. Instead of just adding the a new field, you need to add the new `struct` to
the `client/model.go` file;
2. Instead of  adding the metric to an existing `*_collector.go` file,
create a new file. You may duplicate any of the existing ones,
change it's name, add your metrics there (as you would either way) and wrapping
it on the main `collector.go` file.

## Wrapping up

The main files you need to look into are:

- `client/model.go`
- `collector/collector.go`
- `collector/collector_test.go`
- `collector/testdata/*.json`

You should be able to follow the flow from there and just duplicate what is already
there to export new metrics.
