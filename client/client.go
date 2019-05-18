package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Client is the sgw client
type Client interface {
	Expvar() (Metrics, error)
}

type client struct {
	baseURL string
}

// New creates a new couchbase client
func New(url string) Client {
	return &client{
		baseURL: url,
	}
}

func (c *client) Expvar() (Metrics, error) {
	var metrics Metrics
	resp, err := http.Get(c.baseURL + "/_expvar")
	if err != nil {
		return metrics, errors.Wrap(err, "failed to get metrics")
	}
	bts, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return metrics, errors.Wrap(err, "failed to parse metrics")
	}
	return metrics, json.Unmarshal(bts, &metrics)
}
