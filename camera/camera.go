package camera

import (
	"io/ioutil"
	"strconv"

	"github.com/blackjack/webcam"
)

type Camera struct {
	*webcam.Webcam
	id int32
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
			name:       formatName,
			frameSizes: frameSizes,
		})
	}
	return result
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
