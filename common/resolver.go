package common

import "github.com/chneau/remoteav/camera"

type Resolver struct {
	Cameras_ []*camera.Camera
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.Cameras_
}
