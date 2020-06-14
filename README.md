# cyclone - WIP - not intened for production use

Wrapper around [radix](https://github.com/mediocregopher/radix) with some additional TODO-features.

## Connection

```go
redis := cyclone.NewPool(cyclone.DefaultPool(20)) // or pass radix client
defer redis.Close()
```

## Hash api

```go
redis.Hash("stats").Incr("reqs", 1)
redis.Hash("stats").Set("uptime", "0s")
// ...
ch := redis.Hash("big").Scan().Match("*").Count(10).ChannelKV(50)
for kv := range ch {
  log.Println(kv.Key, "=>", kv.Val)
}
```
