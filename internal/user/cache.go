package user

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var _defaultUser *User = &User{
	ID:        11111,
	Password:  "password",
	History:   &History{},
	Variables: VMap{},
}

var once sync.Once

var defaultCache *Cache

func GetCache(defaultExpiration, cleanupInterval time.Duration) *Cache {

	once.Do(func() {
		data := make(map[string]Item)

		defaultCache = &Cache{
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

func GetTestCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	p := GetCache(defaultExpiration, cleanupInterval)
	p.Set(fmt.Sprintf("%v", _defaultUser.ID), _defaultUser, time.Hour)
	return p
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {

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

func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}

	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("Key not found")
	}

	delete(c.items, key)

	return nil
}

func (c *Cache) Count() int {

	c.Lock()

	defer c.Unlock()

	l := len(c.items)

	return l
}

func (c *Cache) Rename(prewKey, newKey string) error {

	c.Lock()

	defer c.Unlock()

	i, found := c.items[prewKey]
	if !found {
		return errors.New("Key not found")
	}

	c.items[newKey] = i
	delete(c.items, prewKey)

	return nil
}

func (c *Cache) Exist(key string) bool {
	c.RLock()

	defer c.RUnlock()

	_, found := c.items[key]

	return found
}

func (c *Cache) FlushAll() int {

	c.Lock()

	defer c.Unlock()

	am := len(c.items)
	c.items = make(map[string]Item, 0)

	return am
}

func (c *Cache) startGC() {
	go c.gc()
}

func (c *Cache) gc() {

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

func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
