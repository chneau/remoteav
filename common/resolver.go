package common

import (
	"image"
	"log"
	"sync"
	"time"

	"github.com/chneau/remoteav/av"
)

const AudioPath = "/audio"
const VideoPath = "/video"

type Resolver struct {
	cameras        []*av.Camera
	selectedCamera *av.Camera
	videoStream    chan image.Image
	videoMutex     sync.Mutex

	microphones        []*av.Microphone
	selectedMicrophone *av.Microphone
	audioStream        chan []float32
	audioMutex         sync.Mutex
}

func (r *Resolver) SetSelectedMicrophone(args *av.SelectedMicrophone) bool {
	r.audioMutex.Lock()
	defer r.audioMutex.Unlock()
	if r.selectedMicrophone != nil {
		err := r.selectedMicrophone.Close()
		if err != nil {
			log.Println(err)
		}
		r.selectedMicrophone = nil
	}
	for _, microphone := range r.microphones {
		if microphone.Name() == args.Name {
			r.selectedMicrophone = microphone
			return true
		}
	}
	if r.selectedMicrophone == nil {
		log.Println("microphone not found")
		return false
	}
	go r.selectedMicrophone.Stream(r.audioStream)
	return true
}

func (r *Resolver) AudioStream() <-chan []float32 {
	return r.audioStream
}

func (r *Resolver) Microphones() []*av.Microphone {
	return r.microphones
}

func (r *Resolver) AudioPath() string {
	return AudioPath
}

func (r *Resolver) VideoStream() <-chan image.Image {
	return r.videoStream
}

func (r *Resolver) Cameras() []*av.Camera {
	return r.cameras
}

func (r *Resolver) VideoPath() string {
	return VideoPath
}

func (r *Resolver) SetSelectedCamera(args *av.SelectedCamera) bool {
	r.videoMutex.Lock()
	defer r.videoMutex.Unlock()
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
			time.Sleep(time.Millisecond * 100)
			continue
		}
		go r.selectedCamera.Stream(r.videoStream)
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
		videoStream: make(chan image.Image, 10),
		microphones: microphones,
		audioStream: make(chan []float32, 10),
	}
}
