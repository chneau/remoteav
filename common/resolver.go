package common

import (
	"image"
	"log"

	"github.com/chneau/remoteav/av"
)

type Resolver struct {
	cameras     []*av.Camera
	camera      *av.Camera
	imageStream chan image.Image
}

func (r *Resolver) ImageStream() <-chan image.Image {
	return r.imageStream
}

func (r *Resolver) Cameras() []*av.Camera {
	return r.cameras
}

func (r *Resolver) SetSelectedCamera(args *av.SelectedCamera) bool {
	if r.camera != nil {
		err := r.camera.Close()
		if err != nil {
			log.Println(err)
		}
		r.camera = nil
		r.cameras, err = av.GetCameras()
		if err != nil {
			log.Println(err)
		}
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

func NewResolver(cameras []*av.Camera) *Resolver {
	return &Resolver{cameras: cameras, imageStream: make(chan image.Image)}
}
