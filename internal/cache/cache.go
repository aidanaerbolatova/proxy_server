package cache

import (
	"encoding/json"
	"fmt"
	"proxy/models"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]models.Response
}

func NewCache() *Cache {
	return &Cache{
		items: map[string]models.Response{},
	}
}

func (c *Cache) Set(key string, response models.Response) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = response
}

func (c *Cache) Get(key string) (models.Response, bool) {
	c.RLock()
	defer c.RUnlock()

	resp, ok := c.items[key]
	if !ok {
		return resp, false
	}
	return resp, true
}

func (c *Cache) ConvertCacheKey(request models.Request) (string, error) {
	header, err := json.Marshal(request.Headers)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s:%s", request.Method, request.URL, string(header)), nil
}
