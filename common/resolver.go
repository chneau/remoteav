package common

import (
	"image"
	"log"
	"sync"

	"github.com/chneau/remoteav/av"
)

type Resolver struct {
	cameras     []*av.Camera
	camera      *av.Camera
	imageStream chan image.Image
	mutex       sync.Mutex
}

func (r *Resolver) ImageStream() <-chan image.Image {
	return r.imageStream
}

func (r *Resolver) Cameras() []*av.Camera {
	return r.cameras
}

func (r *Resolver) SetSelectedCamera(args *av.SelectedCamera) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
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
	retries := 10
	for retries > 0 {
		retries--
		err := r.camera.StartStreamingFromSelectedCamera(args)
		if err != nil {
			log.Println("retry", err)
			continue
		}
		go r.camera.Stream(r.imageStream)
		return true
	}
	return false
}

func NewResolver(cameras []*av.Camera) *Resolver {
	return &Resolver{cameras: cameras, imageStream: make(chan image.Image, 10)}
}
