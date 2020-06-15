package cyclone

import (
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
				c.Raw.Do(radix.Cmd(c.Raw, "HSET", "HashExists", "a", "1"))

				aExists := c.Hash("HashExists").Exists("a")
				bExists := c.Hash("HashExists").Exists("b")

				g.Assert(aExists).Equal(true)
				g.Assert(bExists).Equal(false)
			})
		})
	})
}
