package common

import (
	"image"

	"github.com/chneau/remoteav/camera"
)

type Resolver struct {
	cameras     []*camera.Camera
	camera      *camera.Camera
	imageStream chan *image.YCbCr
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}

func (r *Resolver) SetSelectedCamera(args *camera.SelectedCamera) bool {
	if r.camera != nil {
		_ = r.camera.StopStreaming()
		r.camera = nil
	}
	for _, camera := range r.cameras {
		if camera.Id() == args.Id {
			r.camera = camera
		}
	}
	if r.camera == nil {
		return false
	}
	err := r.camera.StartStreamingFromSelectedCamera(args)
	if err != nil {
		return false
	}
	return err == nil
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras, imageStream: make(chan *image.YCbCr)}
}
