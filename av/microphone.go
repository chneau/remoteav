package av

import (
	"log"
	"sync"

	"github.com/gordonklaus/portaudio"
)

type Microphone struct {
	deviceInfo *portaudio.DeviceInfo
	stream     *portaudio.Stream
}

func (m *Microphone) Close() error {
	if m.stream == nil {
		return nil
	}
	return m.stream.Close()
}

func (m *Microphone) Stream(stream chan<- []float32) {
	parameters := portaudio.LowLatencyParameters(m.deviceInfo, nil)
	var err error
	m.stream, err = portaudio.OpenStream(parameters, func(in []float32) {
		stream <- in
	})
	if err != nil {
		log.Println(err)
		m.stream = nil
	}
	err = m.stream.Start()
	if err != nil {
		log.Println(err)
		m.stream = nil
	}
}

var micMutex = &sync.Mutex{}
var initialized = false

func GetMicrophones() ([]*Microphone, error) {
	micMutex.Lock()
	defer micMutex.Unlock()
	if !initialized {
		err := portaudio.Initialize()
		if err != nil {
			return nil, err
		}
		initialized = true
	}
	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}
	microphones := []*Microphone{}
	for _, device := range devices {
		if device.MaxInputChannels > 0 {
			microphones = append(microphones, &Microphone{deviceInfo: device})
		}
	}
	return microphones, nil
}
