package pokecache

import (
	"runtime"
	"sync"
	"time"
)

const (
	defaultExpiration = 10 * time.Second
)

type Cache struct {
	c *cache
}

type cache struct {
	c  map[string]entry
	mu sync.RWMutex
	ev *evicter
}

type entry struct {
	createdAt time.Time
	val       []byte
}

func (c *cache) set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.c[key] = entry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *cache) get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, ok := c.c[key]

	return e.val, ok
}

func (c *cache) evict() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()
	for k, v := range c.c {
		itemExp := v.createdAt.Add(c.ev.Interval).UnixNano()
		if itemExp < now {
			delete(c.c, k)
		}
	}
}

func NewCache(exp time.Duration) *Cache {
	if exp < 0 {
		exp = defaultExpiration
	}

	cm := make(map[string]entry)
	c := &cache{
		c: cm,
	}

	stop := make(chan struct{})
	ev := &evicter{
		Interval: exp,
		stop:     stop,
	}
	ev.run(c)
	c.ev = ev

	return &Cache{c}
}

func (c *Cache) Set(key string, val []byte) {
	c.c.set(key, val)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	return c.c.get(key)
}

type evicter struct {
	Interval time.Duration
	stop     chan struct{}
}

func (e *evicter) run(c *cache) {
	runtime.SetFinalizer(c, stopEvicter)
	go e.reap(c)
}

func (e *evicter) reap(c *cache) {
	ticker := time.NewTicker(e.Interval)
	for {
		select {
		case <-ticker.C:
			c.evict()
		case <-e.stop:
			ticker.Stop()
			return
		}
	}
}

func stopEvicter(c *cache) {
	c.ev.stop <- struct{}{}
}
