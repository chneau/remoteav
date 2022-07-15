package common

import "github.com/chneau/remoteav/camera"

type Resolver struct {
	cameras        []*camera.Camera
	selectedCamera *SelectedCamera
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}

type SelectedCamera struct {
	Id         int32
	FormatName string
	FrameSize  string
}

func (r *Resolver) SetCamera(args *SelectedCamera) bool {
	r.selectedCamera = args
	return true
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras}
}
