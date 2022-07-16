package common

import (
	"image"
	"log"

	"github.com/chneau/remoteav/camera"
)

type Resolver struct {
	cameras     []*camera.Camera
	camera      *camera.Camera
	imageStream chan *image.YCbCr
}

func (r *Resolver) ImageStream() <-chan *image.YCbCr {
	return r.imageStream
}

func (r *Resolver) Cameras() []*camera.Camera {
	return r.cameras
}

func (r *Resolver) SetSelectedCamera(args *camera.SelectedCamera) bool {
	if r.camera != nil {
		_ = r.camera.StopStreaming()
		_ = r.camera.Close()
		r.camera = nil
		r.cameras, _ = camera.GetCameras()
	}
	for _, camera := range r.cameras {
		if camera.Id() == args.Id {
			r.camera = camera
		}
	}
	if r.camera == nil {
		log.Println("camera not found")
		return false
	}
	err := r.camera.StartStreamingFromSelectedCamera(args)
	if err != nil {
		log.Println(err)
		return false
	}
	go func() {
		err := r.camera.Stream(r.imageStream)
		if err != nil {
			log.Println(err)
		}
	}()
	return err == nil
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras, imageStream: make(chan *image.YCbCr)}
}
