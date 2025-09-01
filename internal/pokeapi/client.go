package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/sjadczak/gokedex/internal/pokecache"
)

const (
	baseUrl        = "https://pokeapi.co/api/v2/"
	defaultTimeout = time.Second * 10
)

type Client struct {
	url     string
	timeout time.Duration
	cache   *pokecache.Cache
}

func NewClient(options ...func(*Client)) *Client {
	cache := pokecache.NewCache(10 * time.Second)
	client := &Client{
		url:     baseUrl,
		timeout: defaultTimeout,
		cache:   cache,
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

func WithCustomTimeout(to time.Duration) func(*Client) {
	return func(c *Client) {
		c.timeout = to
	}
}

func WithCustomCacheTimeout(de time.Duration) func(*Client) {
	cache := pokecache.NewCache(de)
	return func(c *Client) {
		c.cache = cache
	}
}

func (c *Client) do(endpoint string) ([]byte, error) {
	if ce, ok := c.cache.Get(endpoint); ok {
		return ce, nil
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

	c.cache.Set(endpoint, body)

	return body, nil
}
