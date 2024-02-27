package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Кэш с использованием мьютекса для безопасного доступа к данным
type cacheImpl struct {
	mu    sync.Mutex
	store map[string]string
}

func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		store: make(map[string]string),
	}
}

func (c *cacheImpl) Get(k string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.store[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
