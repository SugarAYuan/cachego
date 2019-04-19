# cachego


## 调用 main函数里有具体的例子

```go
//先初始化
cache.InitLocalCache()
d := new(cache.Data)
d.Key = "test1"
d.Expiration = 5 //过期时间 second
d.Value = "1"
cache.LocalCache.Set(d)
```

> 接下来打算加个多节点同步，然后给一个http接口来进行set和get或者其他方式

> 写这个项目其实完全是造轮子，不过可以学到很多东西。大家有什么意见和建议可以发给我哈。