# Adding new metrics

The exporter simply parses the JSON output of the `/_expvar` endpoint and
returns it in the [Prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/).

So, if a new metrics is exposed under `/_expvar`, we need to:

1. parse it from JSON to a Go `struct`;
2. expose it on the exporter's `/metrics` endpoint.

The parsing from JSON to a `struct` is made on the `client` package. You
can simply add your new field to the appropriate `struct` definition - or create
a new one if necessary - on the `model.go` and that should be it.

We use all numbers as `float64` because that's the format that needs to be
exposed anyway. If a number is returned in `/_expvar` as an `int`, we would
still need to convert it to `float` at some point. Doing it when unmarshalling
the JSON saves the hassle and later usage is simpler.

Next, we need to expose it on `/metrics`. This step is done on the `collector`
package.

In the `collector` package you can see there is one `.go` file for each _"kind"_
of metric.

Once you find the appropriated one, we need to do add 4 things:

1. the metric field in the `thingCollector` `struct` (e.g.: `gsiViewsCollector`);
2. the metric field instantiation on the `newThingCollector`
implementation (e.g.: `newGsiViewsCollector`);
3. add the metric to the `Describe` implementation;
4. finally, add the `Collect` implementation.

The step 4 requires attention regarding which kind of metric you are exporting
(if it's a `Counter` or `Gauge`). Other than that, this is all pretty simple.

Its also good to add a test case for your new metric. All collector tests
are written on `collector_test.go`. The same file has some test helper functions
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

If instead of adding a new metric, you need to add a new set of metrics (a
new object inside the `per_database` object for example), not much changes:

1. instead of just adding the a new field, you need to add the new `struct` to
the `client/model.go` file;
2. instead of just adding the metric on an existing `*_collector.go` file,
you'll need to create a new file. You may duplicate any of the existing ones,
change its name, add your metrics there (as you would either way) and wrapping
it on the main `collector.go` file.

## Wrapping up

The main files you need to look at are:

- `client/model.go`
- `collector/collector.go`
- `collector/collector_test.go`
- `collector/testdata/*.json`

You should be able to follow the flow from there and just mimic what is already
there to export new metrics.
