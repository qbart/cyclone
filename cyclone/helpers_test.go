package cyclone

import (
	"fmt"
	"os"

	"github.com/mediocregopher/radix/v3"
)

func withConn(with func(*Cyclone)) {
	raw, err := radix.NewPool("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), 20)
	if err != nil {
		panic(err)
	}
	c := NewPool(raw)
	defer c.Close()
	with(c)

	raw.Do(radix.Cmd(nil, "FLUSHALL"))
}
