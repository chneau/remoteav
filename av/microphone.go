package av

import (
	"github.com/gordonklaus/portaudio"
)

type Microphone struct {
	*portaudio.DeviceInfo
}

func GetMicrophones() ([]*Microphone, error) {
	microphones := []*Microphone{}
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if device.MaxInputChannels > 0 {
			microphones = append(microphones, &Microphone{DeviceInfo: device})
		}
	}
	return microphones, nil
}
