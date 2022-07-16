package common

import (
	"log"

	"github.com/chneau/remoteav/camera"
)

type Resolver struct {
	cameras        []*camera.Camera
	selectedCamera *SelectedCamera
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}

type SelectedCamera struct {
	Id        int32
	Format    string
	FrameSize string
}

func (r *Resolver) SetSelectedCamera(args *SelectedCamera) bool {
	r.selectedCamera = args
	log.Printf("args: %#+v\n", args)
	return true
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras}
}
