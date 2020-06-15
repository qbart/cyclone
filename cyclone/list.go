package cyclone

import "github.com/mediocregopher/radix/v3"

type List struct {
	cyclone *Cyclone
	key     string
}

func (l *List) Push(elem string) int {
	var count int
	l.cyclone.Raw.Do(radix.Cmd(&count, "LPUSH", l.key, elem))
	return count
}
