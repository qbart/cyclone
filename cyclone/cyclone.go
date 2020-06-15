package cyclone

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

type Cyclone struct {
	Raw *radix.Pool
}

func DefaultPool(n int) *radix.Pool {
	raw, err := radix.NewPool("tcp", "127.0.0.1:6379", n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return raw
}

func NewPool(conn *radix.Pool) *Cyclone {
	return &Cyclone{Raw: conn}
}

func (c *Cyclone) List(key string) *List {
	list := List{cyclone: c, key: key}
	return &list
}

func (c *Cyclone) Hash(key string) *Hash {
	return &Hash{cyclone: c, key: key}
}

func (c *Cyclone) Close() {
	c.Raw.Close()
}
