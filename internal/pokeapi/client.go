package pokeapi

import (
	"io"
	"net/http"
	"time"
)

const (
	baseUrl        = "https://pokeapi.co/api/v2/"
	defaultTimeout = time.Second * 10
)

type Client struct {
	c        *apiCache
	timeout  time.Duration
	url      string
	useCache bool
}

func NewClient(options ...func(*Client)) *Client {
	client := &Client{
		url:      baseUrl,
		timeout:  defaultTimeout,
		useCache: true,
		c:        NewCache(),
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

func WithCustomCacheExpiration(d time.Duration) func(*Client) {
	return func(c *Client) {
		cache := NewCache(
			WithCustomExpiration(d),
		)
		c.c = cache
	}
}

func WithoutCache() func(*Client) {
	return func(c *Client) {
		c.c = nil
		c.useCache = false
	}
}

func (c *Client) do(endpoint string) ([]byte, error) {
	// Check if endpoint has been cached
	cached, found := c.c.Get(endpoint)
	if found && c.useCache {
		return cached, nil
	}

	req, err := http.NewRequest(http.MethodGet, c.url+endpoint, nil)
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{Timeout: c.timeout}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	c.c.Set(endpoint, body)
	return body, nil
}
