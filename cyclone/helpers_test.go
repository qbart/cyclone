package cyclone

func withConn(with func(*Cyclone)) {
	c := NewPool(DefaultPool(20))
	defer c.Close()
	with(c)
}
