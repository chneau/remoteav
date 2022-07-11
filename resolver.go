package main

import "github.com/chneau/remoteav/camera"

type Resolver struct {
	cameras []*camera.Camera
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}
