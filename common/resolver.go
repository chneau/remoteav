package common

import (
	"github.com/blackjack/webcam"
	"github.com/chneau/remoteav/camera"
	"github.com/samber/lo"
)

type Resolver struct {
	cameras        []*camera.Camera
	camera         *camera.Camera
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
	pixelFormat, found := lo.FindKey(r.camera.GetSupportedFormats(), args.Format)
	if !found {
		return false
	}
	frameSize, found := lo.Find(r.camera.GetSupportedFrameSizes(pixelFormat), func(frameSize webcam.FrameSize) bool {
		return frameSize.GetString() == args.FrameSize
	})
	if !found {
		return false
	}
	_, _, _, err := r.camera.SetImageFormat(pixelFormat, frameSize.MaxWidth, frameSize.MaxHeight)
	return err == nil
}

func NewResolver(cameras []*camera.Camera) *Resolver {
	return &Resolver{cameras: cameras}
}
