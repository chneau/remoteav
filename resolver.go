package main

type Resolver struct {
}

func (_ *Resolver) Cameras() []*Camera {
	return []*Camera{{id: 1}, {id: 2}, {id: 3}}
}

type Camera struct {
	id int32
}

func (c *Camera) Id() int32 {
	return c.id
}
