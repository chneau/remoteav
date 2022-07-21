package common

import (
	"image"
	"log"
	"sync"

	"github.com/chneau/remoteav/av"
)

type Resolver struct {
	cameras        []*av.Camera
	selectedCamera *av.Camera
	imageStream    chan image.Image
	imageMutex     sync.Mutex

	microphones        []*av.Microphone
	selectedMicrophone *av.Microphone
	audioStream        chan []float32
	audioMutex         sync.Mutex
}

func (r *Resolver) SetSelectedMicrophone(args *av.SelectedMicrophone) bool {
	return false
}

func (r *Resolver) AudioMutex() <-chan []float32 {
	return r.audioStream
}

func (r *Resolver) Microphones() []*av.Microphone {
	return r.microphones
}

func (r *Resolver) ImageStream() <-chan image.Image {
	return r.imageStream
}

func (r *Resolver) Cameras() []*av.Camera {
	return r.cameras
}

func (r *Resolver) SetSelectedCamera(args *av.SelectedCamera) bool {
	r.imageMutex.Lock()
	defer r.imageMutex.Unlock()
	if r.selectedCamera != nil {
		err := r.selectedCamera.Close()
		if err != nil {
			log.Println(err)
		}
		r.selectedCamera = nil
		r.cameras, err = av.GetCameras()
		if err != nil {
			log.Println(err)
		}
	}
	for _, camera := range r.cameras {
		if camera.Id() == args.Id {
			r.selectedCamera = camera
		}
	}
	if r.selectedCamera == nil {
		log.Println("camera not found")
		return false
	}
	retries := 10
	for retries > 0 {
		retries--
		err := r.selectedCamera.StartStreamingFromSelectedCamera(args)
		if err != nil {
			log.Println("retry", err)
			continue
		}
		go r.selectedCamera.Stream(r.imageStream)
		return true
	}
	return false
}

func NewResolver() *Resolver {
	cameras, err := av.GetCameras()
	if err != nil {
		log.Println(err)
	}
	microphones, err := av.GetMicrophones()
	if err != nil {
		log.Println(err)
	}
	return &Resolver{
		cameras:     cameras,
		imageStream: make(chan image.Image, 10),
		microphones: microphones,
		audioStream: make(chan []float32, 10),
	}
}
