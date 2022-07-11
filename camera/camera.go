package camera

import (
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

func New(id int) (*Camera, error) {
	cam, err := webcam.Open("/dev/video" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	return &Camera{id: int32(id), Webcam: cam}, nil
}
