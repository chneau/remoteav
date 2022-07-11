package main

type Resolver struct {
	cameras []*Camera
}

func (r *Resolver) Cameras() []*Camera {
	return r.cameras
}

type Camera struct {
	id int32
}

func (c *Camera) Id() int32 {
	return c.id
}
