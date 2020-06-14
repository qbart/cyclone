[![LICENSE](https://img.shields.io/github/license/qbart/cyclone)](https://github.com/qbart/cyclone/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbart/cyclone)](https://goreportcard.com/report/github.com/qbart/cyclone)
[![Last commit](https://img.shields.io/github/last-commit/qbart/cyclone)](https://github.com/qbart/cyclone/commits/master)
# cyclone - WIP - not intened for production use

Wrapper around [radix](https://github.com/mediocregopher/radix) with some additional TODO-features. [GoDoc here](https://pkg.go.dev/github.com/qbart/cyclone/cyclone)

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
