package cyclone

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/mediocregopher/radix/v3"
)

func TestList(t *testing.T) {
	g := Goblin(t)
	withConn(func(c *Cyclone) {
		g.Describe(".BPop", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".BRPop", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".BRPopLPush", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".Index", func() {
			g.It("Returns element at index", func() {
				list := c.List("ListLIndex")

				// empty list
				elem := list.Index(0)
				g.Assert(elem).Eql("")

				// with elments
				g.Assert(list.Push("a", "b")).Eql(2)
				elem = list.Index(1)
				g.Assert(elem).Eql("a")
			})
		})

		g.Describe(".Insert", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".Len", func() {
			g.It("Returns len of the list", func() {
				c.List("ListLen").Push("a", "b", "c")

				length := c.List("ListLen").Len()
				g.Assert(length).Eql(3)
			})
		})

		g.Describe(".Pop", func() {
			g.It("Pops element from HEAD", func() {
				list := c.List("ListLPop")

				// empty list
				elem := list.Pop()
				g.Assert(elem).Eql("")

				list.Push("a", "b")
				elem = list.Pop()
				g.Assert(elem).Eql("b")
				g.Assert(list.Range(0, -1)).Eql([]string{"a"})
			})
		})

		g.Describe(".Pos", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".Push", func() {
			g.It("Pushes elements into list HEAD", func() {
				c.List("ListLPush").Push("a", "b", "c")
				length := c.List("ListLPush").Push("d")

				g.Assert(length).Eql(4)

				elems := c.List("ListLPush").Range(0, -1)
				g.Assert(elems[0]).Eql("d")
				g.Assert(elems[1]).Eql("c")
				g.Assert(elems[2]).Eql("b")
				g.Assert(elems[3]).Eql("a")
			})
		})

		g.Describe(".PushX", func() {
			g.It("Pushes elements into list HEAD only if list exists", func() {
				c.List("ListLPushX").Push("a")
				lenExisting := c.List("ListLPushX").PushX("b")
				lenNonExisting := c.List("ListLPushXother").PushX("c")

				g.Assert(lenExisting).Eql(2)
				g.Assert(lenNonExisting).Eql(0)

				elems := c.List("ListLPushX").Range(0, -1)
				g.Assert(elems[0]).Eql("b")
				g.Assert(elems[1]).Eql("a")

				elems = c.List("ListLPushXother").Range(0, -1)
				g.Assert(len(elems)).Eql(0)
			})
		})

		g.Describe(".Range", func() {
			g.It("Returns range from given offsets", func() {
				c.Raw.Do(radix.Cmd(nil, "RPUSH", "ListLRange", "a", "b", "c"))

				// full range
				elems := c.List("ListLRange").Range(0, -1)

				g.Assert(len(elems)).Equal(3)
				g.Assert(elems[0]).Eql("a")
				g.Assert(elems[1]).Eql("b")
				g.Assert(elems[2]).Eql("c")

				// single element
				elems = c.List("ListLRange").Range(1, 1)

				g.Assert(len(elems)).Equal(1)
				g.Assert(elems[0]).Eql("b")

				// out of bounds
				elems = c.List("ListLRange").Range(4, 6)

				g.Assert(len(elems)).Equal(0)
			})
		})

		g.Describe(".Rem", func() {
			g.It("Removes elments from list", func() {
				list := c.List("ListLRem")

				// empty list
				g.Assert(
					list.Rem(0, "a"),
				).Eql(0)

				// count == 0
				list.Push("a", "b", "a", "b")
				g.Assert(
					list.Rem(0, "a"),
				).Eql(2)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"b", "b"})

				// count > 0
				list.Push("c", "b", "b")
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"b", "b", "c", "b", "b"})
				g.Assert(
					list.Rem(2, "b"),
				).Eql(2)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"c", "b", "b"})

				// count < 0
				list.Push("b")
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"b", "c", "b", "b"})
				g.Assert(
					list.Rem(-2, "b"),
				).Eql(2)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"b", "c"})
			})
		})

		g.Describe(".Set", func() {
			g.It("Sets element at index", func() {
				list := c.List("ListLSet")

				// empty list
				g.Assert(
					list.Set(0, "a"),
				).Eql(false)

				// add some data
				list.Push("b", "a")
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"a", "b"})

				// with existing index
				g.Assert(
					list.Set(1, "c"),
				).Eql(true)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"a", "c"})
			})
		})

		g.Describe(".Trim", func() {
			g.It("Trims list to given range", func() {
				list := c.List("ListLTrim")
				list.Push("f", "e", "d", "c", "b", "a")

				g.Assert(
					list.Trim(0, 3),
				).Eql(true)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"a", "b", "c", "d"})

				g.Assert(
					list.Trim(1, -3),
				).Eql(true)
				g.Assert(
					list.Range(0, -1),
				).Eql([]string{"b"})
			})
		})

		g.Describe(".RPop", func() {
			g.It("Pops element from TAIL", func() {
				list := c.List("ListRPop")

				// empty list
				elem := list.RPop()
				g.Assert(elem).Eql("")

				list.RPush("a", "b")
				elem = list.RPop()
				g.Assert(elem).Eql("b")
				g.Assert(list.Range(0, -1)).Eql([]string{"a"})
			})
		})

		g.Describe(".RPopLPush", func() {
			g.Xit("", func() {
			})
		})

		g.Describe(".RPush", func() {
			g.It("Pushes elements into list TAIL", func() {
				c.List("ListRPush").Push("a", "b", "c")
				length := c.List("ListRPush").RPush("d")

				g.Assert(length).Eql(4)

				elems := c.List("ListRPush").Range(0, -1)
				g.Assert(elems[0]).Eql("c")
				g.Assert(elems[1]).Eql("b")
				g.Assert(elems[2]).Eql("a")
				g.Assert(elems[3]).Eql("d")
			})
		})

		g.Describe(".RPushX", func() {
			g.It("Pushes elements into list TAIL only if list exists", func() {
				c.List("ListRPushX").RPush("a")
				lenExisting := c.List("ListRPushX").RPushX("b")
				lenNonExisting := c.List("ListRPushXother").RPushX("c")

				g.Assert(lenExisting).Eql(2)
				g.Assert(lenNonExisting).Eql(0)

				elems := c.List("ListRPushX").Range(0, -1)
				g.Assert(elems[0]).Eql("a")
				g.Assert(elems[1]).Eql("b")

				elems = c.List("ListRPushXother").Range(0, -1)
				g.Assert(len(elems)).Eql(0)
			})
		})
	})
}
