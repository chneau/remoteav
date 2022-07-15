package common

import "github.com/chneau/remoteav/camera"

type Resolver struct {
	cameras []*camera.Camera
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras}
}
