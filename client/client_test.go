package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()
	metrics, err := New(ts.URL).Expvar()
	require.NoError(t, err)
	require.Equal(t, Metrics{}, metrics)
}

func TestClientInvalidJson(t *testing.T) {
	var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	}))
	defer ts.Close()
	metrics, err := New(ts.URL).Expvar()
	require.EqualError(t, err, "failed to parse metrics: unexpected end of JSON input")
	require.Equal(t, Metrics{}, metrics)
}

func TestClientServerDown(t *testing.T) {
	metrics, err := New("http://fakehost:2132").Expvar()
	// nolint: lll
	require.EqualError(t, err, "failed to get metrics: Get http://fakehost:2132/_expvar: dial tcp: lookup fakehost: no such host")
	require.Equal(t, Metrics{}, metrics)
}

func TestClientNonSuccessfulResponse(t *testing.T) {
	var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()
	metrics, err := New(ts.URL).Expvar()
	require.EqualError(t, err, "failed to get metrics: 400 Bad Request")
	require.Equal(t, Metrics{}, metrics)
}
