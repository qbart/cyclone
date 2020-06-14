package cyclone

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

type Data struct {
	conn *radix.Pool
}

func DefaultPool(n int) *radix.Pool {
	c, err := radix.NewPool("tcp", "127.0.0.1:6379", n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return c
}

func NewPool(conn *radix.Pool) *Data {
	return &Data{conn: conn}
}

func (d *Data) List(key string) *List {
	list := List{data: d, key: key}
	return &list
}

func (d *Data) Hash(key string) *Hash {
	return &Hash{data: d, key: key}
}

func (d *Data) Close() {
	d.conn.Close()
}
