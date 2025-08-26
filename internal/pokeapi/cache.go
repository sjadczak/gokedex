package pokeapi

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	DefaultExpire   = time.Minute * 5
	CleanupInterval = time.Minute * 10
)

type apiCache struct {
	c            *cache.Cache
	customExpire time.Duration
}

func NewCache(options ...func(*apiCache)) *apiCache {
	cache := cache.New(DefaultExpire, CleanupInterval)
	c := &apiCache{c: cache}

	for _, o := range options {
		o(c)
	}

	return c
}

func WithCustomExpiration(d time.Duration) func(*apiCache) {
	return func(c *apiCache) {
		c.customExpire = d
	}
}

func (c *apiCache) Set(endpoint string, body []byte) {
	if c.customExpire > 0 {
		c.c.Set(endpoint, body, c.customExpire)
	} else {
		c.c.SetDefault(endpoint, body)
	}
}

func (c *apiCache) Get(endpoint string) ([]byte, bool) {
	res, found := c.c.Get(endpoint)
	if !found {
		return []byte{}, false
	}
	return res.([]byte), true
}
