package cyclone

import "github.com/mediocregopher/radix/v3"

type List struct {
	data *Data
	key  string
}

func (l *List) Push(elem string) int {
	var count int
	l.data.conn.Do(radix.Cmd(&count, "LPUSH", l.key, elem))
	return count
}
