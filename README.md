[![GoDoc](https://godoc.org/github.com/qbart/cyclone?status.svg)](https://pkg.go.dev/github.com/qbart/cyclone/cyclone)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/qbart/cyclone.svg)
[![CI](https://github.com/qbart/cyclone/workflows/CI/badge.svg?branch=master)](https://github.com/qbart/cyclone/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbart/cyclone)](https://goreportcard.com/report/github.com/qbart/cyclone)
[![Last commit](https://img.shields.io/github/last-commit/qbart/cyclone)](https://github.com/qbart/cyclone/commits/master)

# cyclone - WIP

Wrapper around [radix](https://github.com/mediocregopher/radix) with some additional TODO-features.
It contains comments from [redis.io/commands](https://redis.io/commands) for a convenience but you should always refer to offical Redis docs.

## Connection

```go
redis := cyclone.NewPool(cyclone.DefaultPool(20)) // or pass radix client
defer redis.Close()
```

## Hash

```go
redis.Hash("stats").Incr("reqs", 1)
redis.Hash("stats").Set("uptime", "0s")
// ...
ch := redis.Hash("big").Scan().Match("*").Count(10).ChanKV(50)
for kv := range ch {
  log.Println(kv.Key, "=>", kv.Val)
}
```

## List

```go
redis.List("list").Push("a", "b", "c")
redis.List("list").RPop()
// ...
```
