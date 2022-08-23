package user

import (
	"errors"
	"sync"
	"time"
)

var once sync.Once

var defaultCache *inMemoryCache

type inMemoryCache struct {
	defaultExpiration time.Duration
	cleanupInterval   time.Duration

	sync.RWMutex
	items map[string]Item
}

func GetCache(defaultExpiration, cleanupInterval time.Duration) *inMemoryCache {

	once.Do(func() {
		data := make(map[string]Item)

		defaultCache = &inMemoryCache{
			items:             data,
			defaultExpiration: defaultExpiration,
			cleanupInterval:   cleanupInterval,
		}

		if cleanupInterval > 0 {
			defaultCache.startGC()
		}
	})

	return defaultCache
}

func (c *inMemoryCache) Set(key string, value interface{}, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()

	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}

func (c *inMemoryCache) Get(key string) (interface{}, error) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, &ItemNotFoundError{}
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, errors.New("Item is outdated")
		}

	}

	return item.Value, nil
}

func (c *inMemoryCache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return &ItemNotFoundError{}
	}

	delete(c.items, key)

	return nil
}

func (c *inMemoryCache) Count() int {

	c.Lock()

	defer c.Unlock()

	l := len(c.items)

	return l
}

func (c *inMemoryCache) Rename(prewKey, newKey string) error {

	c.Lock()

	defer c.Unlock()

	i, found := c.items[prewKey]
	if !found {
		return &ItemNotFoundError{}
	}

	c.items[newKey] = i
	delete(c.items, prewKey)

	return nil
}

func (c *inMemoryCache) Exist(key string) bool {
	c.RLock()

	defer c.RUnlock()

	_, found := c.items[key]

	return found
}

func (c *inMemoryCache) FlushAll() int {

	c.Lock()

	defer c.Unlock()

	am := len(c.items)
	c.items = make(map[string]Item, 0)

	return am
}

func (c *inMemoryCache) startGC() {
	go c.gc()
}

func (c *inMemoryCache) gc() {

	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

func (c *inMemoryCache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *inMemoryCache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}

type stabCache struct{}

func GetDummyCache(defaultExpiration, cleanupInterval time.Duration) *stabCache {
	return &stabCache{}
}

func (c *stabCache) Set(key string, value interface{}, duration time.Duration) {}

func (c *stabCache) Get(key string) (interface{}, error) {
	return nil, &ItemNotFoundError{}
}

func (c *stabCache) Delete(key string) error {
	return nil
}

func (c *stabCache) Count() int {
	return 0
}

func (c *stabCache) Rename(prewKey, newKey string) error {
	return nil
}

func (c *stabCache) Exist(key string) bool {
	return false
}

func (c *stabCache) FlushAll() int {
	return 0
}
