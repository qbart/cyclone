package cyclone

import (
	"fmt"
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

// Cyclone wraps radix client.
type Cyclone struct {
	Raw *radix.Pool
}

// DefafultPool creates default connection to redis or exists when failed.
func DefaultPool(n int) *radix.Pool {
	raw, err := radix.NewPool("tcp", "127.0.0.1:6379", n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return raw
}

// NewPool creates Cyclone wrapper around radix.Pool
func NewPool(conn *radix.Pool) *Cyclone {
	return &Cyclone{Raw: conn}
}

// List returns list wrapper.
func (c *Cyclone) List(key string) *List {
	list := List{cyclone: c, key: key}
	return &list
}

// Hash returns Hash wrapper.
func (c *Cyclone) Hash(key string) *Hash {
	return &Hash{cyclone: c, key: key}
}

// Hashf returns Hash wrapper. Key is built from fmt.Sprintf(format, any...).
func (c *Cyclone) Hashf(format string, any ...interface{}) *Hash {
	return &Hash{cyclone: c, key: fmt.Sprintf(key, any...)}
}

// Close closes current connection.
func (c *Cyclone) Close() {
	c.Raw.Close()
}
