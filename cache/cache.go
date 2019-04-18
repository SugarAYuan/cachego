package cache

import (
	"sort"
	"sync"
	"time"
)

type Cache struct {
	datas    map[string]*Data
	mu       sync.RWMutex
	hmu      sync.RWMutex
	dataHeap []*Data
}

type Data struct {
	Key        string
	Value      interface{}
	cleanTime  time.Time
	Expiration int64
}

var LocalCache *Cache

func InitLocalCache() {
	LocalCache = &Cache{
		datas: make(map[string]*Data),
		mu:    sync.RWMutex{},
	}
	go func(c *Cache) {
		for {
			c.hmu.RLock()
			if c.Len() < 1 {
				c.hmu.RUnlock()
				continue
			}
			if c.dataHeap[c.Len()-1].cleanTime.Before(time.Now()) {
				c.hmu.RUnlock()
				d := c.pop()
				c.ttl(d.Key)
			} else {
				c.hmu.RUnlock()
			}

		}
	}(LocalCache)
}

func (c *Cache) Len() int {
	return len(c.dataHeap)
}

func (c *Cache) Swap(i, j int) {
	c.dataHeap[j], c.dataHeap[i] = c.dataHeap[i], c.dataHeap[j]
}

func (c *Cache) Less(i, j int) bool {
	return c.dataHeap[i].cleanTime.After(c.dataHeap[j].cleanTime)
}

func (c *Cache) push(data *Data) {
	c.hmu.Lock()
	c.dataHeap = append(c.dataHeap, data)
	sort.Sort(c)
	c.hmu.Unlock()
}

func (c *Cache) pop() *Data {
	d := &Data{}
	c.hmu.Lock()
	l := len(c.dataHeap)
	c.dataHeap, d = c.dataHeap[:l-1], c.dataHeap[l-1]
	c.hmu.Unlock()
	return d
}

func (c *Cache) ttl(key string) {
	c.mu.RLock()
	if _, ok := c.datas[key]; ok {
		if c.datas[key].cleanTime.Before(time.Now()) {
			c.mu.RUnlock()
			c.Remove(key)
		} else {
			c.mu.RUnlock()
		}
	} else {
		c.mu.RUnlock()
	}

}

func (c *Cache) Get(key string) *Data {
	data := new(Data)
	ok := false
	c.mu.RLock()
	if data, ok = c.datas[key]; !ok {
		data = nil
	}
	c.mu.RUnlock()
	return data
}

func (c *Cache) Set(data *Data) bool {
	data.cleanTime = time.Now().Add(time.Duration(data.Expiration) * time.Second)
	c.mu.Lock()
	c.datas[data.Key] = data
	c.mu.Unlock()
	c.push(data)
	return true
}

func (c *Cache) Remove(key string) bool {
	b := c.Get(key)
	if b == nil {
		return false
	}
	c.mu.Lock()
	delete(c.datas, key)
	c.mu.Unlock()
	return true
}
