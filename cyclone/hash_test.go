package cyclone

import (
	"log"
	"strconv"
	"testing"

	. "github.com/franela/goblin"
	"github.com/mediocregopher/radix/v3"
)

func TestHash(t *testing.T) {
	g := Goblin(t)
	withConn(func(c *Cyclone) {
		g.Describe(".Del", func() {
			g.It("Deletes keys and returns deleted count", func() {
				var before, after map[string]string
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashDel", "a", "1", "b", "2", "c", "3"))
				c.Raw.Do(radix.Cmd(&before, "HGETALL", "HashDel"))

				g.Assert(before["a"]).Eql("1")
				g.Assert(before["b"]).Eql("2")
				g.Assert(before["c"]).Eql("3")

				deletedKeys := c.Hash("HashDel").Del("a", "b")

				g.Assert(deletedKeys).Equal(2)
				c.Raw.Do(radix.Cmd(&after, "HGETALL", "HashDel"))
				g.Assert(after["a"]).Eql("")
				g.Assert(after["b"]).Eql("")
				g.Assert(after["c"]).Eql("3")
			})
		})

		g.Describe(".Exists", func() {
			g.It("Returns T for existing keys, F otherwise", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashExists", "a", "1"))

				aExists := c.Hash("HashExists").Exists("a")
				bExists := c.Hash("HashExists").Exists("b")

				g.Assert(aExists).Equal(true)
				g.Assert(bExists).Equal(false)
			})
		})

		g.Describe(".Get", func() {
			g.It("Returns value for a key", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashGet", "a", "1"))

				val := c.Hash("HashGet").Get("a")

				g.Assert(val).Equal("1")
			})
		})

		g.Describe(".GetAll", func() {
			g.It("Returns all value for a key", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashGetAll", "a", "1", "b", "2"))

				val := c.Hash("HashGetAll").GetAll()

				g.Assert(val).Equal(map[string]string{
					"a": "1",
					"b": "2",
				})
			})
		})

		g.Describe(".Incr", func() {
			g.It("Increments a floating field and returns incrmented value", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashIncr", "a", "1", "b", "2"))

				val1 := c.Hash("HashIncr").Incr("a", 5)
				val2 := c.Hash("HashIncr").Get("a")

				g.Assert(val1).Eql(6)
				g.Assert(val2).Equal("6")
			})
		})

		g.Describe(".IncrFloat", func() {
			g.It("Increments an integer field and returns incrmented value", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashIncrFloat", "a", "3.14"))

				val1 := c.Hash("HashIncrFloat").IncrFloat("a", -0.43)
				val2 := c.Hash("HashIncrFloat").Get("a")

				g.Assert(val1).Eql(2.71)
				g.Assert(val2).Equal("2.71")
			})
		})

		g.Describe(".Keys", func() {
			g.It("Returns hash keys", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashKeys", "a", "1", "b", "2"))

				val := c.Hash("HashKeys").Keys()

				g.Assert(val).Eql([]string{"a", "b"})
			})
		})

		g.Describe(".Len", func() {
			g.It("Returns num of keys", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashLen", "a", "1", "b", "2"))

				val := c.Hash("HashLen").Len()

				g.Assert(val).Eql(2)
			})
		})

		g.Describe(".MGet", func() {
			g.It("Returns selected key values", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashMGet", "a", "1", "b", "2", "c", "3"))

				val := c.Hash("HashMGet").MGet("a", "c")

				g.Assert(val).Eql([]string{"1", "3"})
			})
		})

		g.Describe(".Scan", func() {
			g.It("Chan iteration", func() {
				for i := 0; i < 100; i++ {
					c.Raw.Do(radix.Cmd(nil, "HSET", "HashScan", strconv.Itoa(i), strconv.Itoa(i)))
				}

				ch := c.Hash("HashScan").Scan().Match("88").Chan(0)
				result := make([]string, 0)
				for field := range ch {
					result = append(result, field)
				}

				g.Assert(len(result)).Eql(2)
				g.Assert(result[0]).Eql("88")
				g.Assert(result[1]).Eql("88")
			})

			g.It("ChanKV iteration", func() {
				for i := 0; i < 100; i++ {
					c.Raw.Do(radix.Cmd(nil, "HSET", "HashScan_KV", strconv.Itoa(i), strconv.Itoa(i)))
				}

				ch := c.Hash("HashScan_KV").Scan().ChanKV(0)
				result := make(map[string]string, 0)
				for kv := range ch {
					result[kv.Key] = kv.Val
				}

				g.Assert(len(result)).Eql(100)
				for i := 0; i < 100; i++ {
					g.Assert(result[strconv.Itoa(i)]).Eql(strconv.Itoa(i))
				}
			})

			g.It("ChanKV Match iteration", func() {
				for i := 0; i < 100; i++ {
					c.Raw.Do(radix.Cmd(nil, "HSET", "HashScan_KV_Match", strconv.Itoa(i), strconv.Itoa(i)))
				}

				ch := c.Hash("HashScan_KV_Match").Scan().Match("2*").ChanKV(0)
				result := make(map[string]string, 0)
				for kv := range ch {
					log.Println(kv)
					result[kv.Key] = kv.Val
				}

				g.Assert(len(result)).Eql(11)
				g.Assert(result["2"]).Eql("2")
				g.Assert(result["21"]).Eql("21")
				g.Assert(result["22"]).Eql("22")
				g.Assert(result["23"]).Eql("23")
				g.Assert(result["24"]).Eql("24")
				g.Assert(result["25"]).Eql("25")
				g.Assert(result["26"]).Eql("26")
				g.Assert(result["27"]).Eql("27")
				g.Assert(result["28"]).Eql("28")
				g.Assert(result["29"]).Eql("29")
			})
		})

		g.Describe(".Set", func() {
			g.It("Sets fields by K/V pairs", func() {
				addedFields := c.Hash("HashSet").Set(
					"a", "1",
					"b", "2",
					"c", "3",
				)
				val := c.Hash("HashSet").MGet("a", "c")

				g.Assert(addedFields).Eql(3)
				g.Assert(val).Eql([]string{"1", "3"})
			})
		})

		g.Describe(".SetNX", func() {
			g.It("Sets field by K/V only if not exists", func() {
				first := c.Hash("HashSetNX").SetNX("a", "1")
				second := c.Hash("HashSetNX").SetNX("a", "2")
				val := c.Hash("HashSetNX").Get("a")

				g.Assert(first).Eql(true)
				g.Assert(second).Eql(false)
				g.Assert(val).Eql("1")
			})
		})

		g.Describe(".StrLen", func() {
			g.It("Returns field value length", func() {
				c.Raw.Do(radix.Cmd(nil, "HSET", "HashStrLen", "a", "12", "b", "ᴓ"))

				lenOfa := c.Hash("HashStrLen").StrLen("a")
				lenOfb := c.Hash("HashStrLen").StrLen("b")

				g.Assert(lenOfa).Eql(2)
				g.Assert(lenOfb).Eql(3) // byte-length of ᴓ (e1b493)
			})
		})

		g.Describe(".Vals", func() {
			g.It("Return all values", func() {
				c.Hash("HashVals").Set(
					"a", "1",
					"b", "2",
					"c", "3",
				)
				val := c.Hash("HashVals").Vals()

				g.Assert(val).Eql([]string{"1", "2", "3"})
			})
		})
	})
}
