package client

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

func (*client) Expvar() (Metrics, error) {
	return Metrics{}, nil
}

// Metrics JSON representation from /_expvar
type Metrics struct {
}
