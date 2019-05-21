package client

import (
	"encoding/json"
	"fmt"
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
	if resp.StatusCode != 200 {
		return metrics, fmt.Errorf("failed to get metrics: %s", resp.Status)
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return metrics, errors.Wrap(err, "failed to parse metrics")
	}
	if err := json.Unmarshal(bts, &metrics); err != nil {
		return metrics, errors.Wrap(err, "failed to parse metrics")
	}
	return metrics, nil
}
