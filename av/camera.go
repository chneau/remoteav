package av

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/blackjack/webcam"
	"github.com/samber/lo"
)

type SelectedCamera struct {
	Id        int32
	Format    string
	FrameSize string
}

type Camera struct {
	*webcam.Webcam
	id          int32
	frameSize   webcam.FrameSize
	pixelFormat webcam.PixelFormat
}

func (c *Camera) Id() int32 {
	return c.id
}

func (c *Camera) SupportedFormats() []SupportedFormat {
	result := []SupportedFormat{}
	for format, formatName := range c.GetSupportedFormats() {
		frameSizes := []string{}
		for _, frameSize := range c.GetSupportedFrameSizes(format) {
			frameSizes = append(frameSizes, frameSize.GetString())
		}
		result = append(result, SupportedFormat{
			format:     formatName,
			frameSizes: frameSizes,
		})
	}
	return result
}

func (c *Camera) StartStreamingFromSelectedCamera(settings *SelectedCamera) error {
	if settings == nil {
		return errors.New("settings is nil")
	}
	if c.id != settings.Id {
		return errors.New("camera id mismatch")
	}
	pixelFormat, found := lo.FindKey(c.GetSupportedFormats(), settings.Format)
	if !found {
		return errors.New("unsupported format")
	}
	frameSize, found := lo.Find(c.GetSupportedFrameSizes(pixelFormat), func(frameSize webcam.FrameSize) bool {
		return frameSize.GetString() == settings.FrameSize
	})
	if !found {
		return errors.New("unsupported frame size")
	}
	_, _, _, err := c.SetImageFormat(pixelFormat, frameSize.MaxWidth, frameSize.MaxHeight)
	if err != nil {
		return err
	}
	err = c.StartStreaming()
	if err != nil {
		return err
	}
	c.frameSize = frameSize
	c.pixelFormat = pixelFormat
	return nil
}

func (c *Camera) Stream(imageStream chan image.Image) {
	for {
		err := c.Webcam.WaitForFrame(100)
		if err != nil {
			log.Println(err)
		}
		frame, err := c.Webcam.ReadFrame()
		if err != nil {
			return
		}
		if c.pixelFormat == 1448695129 { // YUYV 4:2:2
			yuyv := image.NewYCbCr(image.Rect(0, 0, int(c.frameSize.MaxWidth), int(c.frameSize.MaxHeight)), image.YCbCrSubsampleRatio422)
			for i := range yuyv.Cb {
				ii := i * 4
				yuyv.Y[i*2] = frame[ii]
				yuyv.Y[i*2+1] = frame[ii+2]
				yuyv.Cb[i] = frame[ii+1]
				yuyv.Cr[i] = frame[ii+3]
			}
			imageStream <- yuyv
		}
		if c.pixelFormat == 1196444237 { // Motion-JPEG
			img, err := jpeg.Decode(bytes.NewReader(frame))
			if err != nil {
				log.Println(err)
			}
			imageStream <- img
		}
	}
}

func new(id int) (*Camera, error) {
	cam, err := webcam.Open("/dev/video" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	return &Camera{id: int32(id), Webcam: cam}, nil
}

func GetCameras() ([]*Camera, error) {
	files, err := ioutil.ReadDir("/dev")
	if err != nil {
		return nil, err
	}
	result := []*Camera{}
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) <= 5 || fileName[:5] != "video" {
			continue
		}
		cameraNumber, err := strconv.Atoi(fileName[5:])
		if err != nil {
			return nil, err
		}
		cam, err := new(cameraNumber)
		if err != nil {
			return nil, err
		}
		result = append(result, cam)
	}
	return result, nil
}

type SupportedFormat struct {
	format     string
	frameSizes []string
}

func (s SupportedFormat) Format() string {
	return s.format
}

func (s SupportedFormat) FrameSizes() []string {
	return s.frameSizes
}
