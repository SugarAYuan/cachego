package main

import (
	"cachego/cache"
)

func main() {
	cache.InitLocalCache()
	d := new(cache.Data)
	d.Key = "test1"
	d.Expiration = 5
	d.Value = "1"
	cache.LocalCache.Set(d)
	d = new(cache.Data)
	d.Key = "test2"
	d.Expiration = 10
	d.Value = "2"
	cache.LocalCache.Set(d)
	d = new(cache.Data)
	d.Key = "test3"
	d.Expiration = 15
	d.Value = "3"
	cache.LocalCache.Set(d)
	d = new(cache.Data)
	d.Key = "test4"
	d.Expiration = 20
	d.Value = "4"
	cache.LocalCache.Set(d)
	d = new(cache.Data)
	d.Key = "test5"
	d.Expiration = 5
	d.Value = "5"
	cache.LocalCache.Set(d)

}
