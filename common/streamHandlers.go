package common

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
)

const boundary = "frame"

func StreamVideoHandler(ch <-chan image.Image) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/x-mixed-replace;boundary="+boundary)
		multipartWriter := multipart.NewWriter(w)
		err := multipartWriter.SetBoundary(boundary)
		if err != nil {
			panic(err)
		}
		for image := range ch {
			bb := &bytes.Buffer{}
			err := jpeg.Encode(bb, image, &jpeg.Options{Quality: 40})
			if err != nil {
				panic(err)
			}
			b := bb.Bytes()
			iw, err := multipartWriter.CreatePart(textproto.MIMEHeader{
				"Content-type":   []string{"image/jpeg"},
				"Content-length": []string{strconv.Itoa(len(b))},
			})
			if err != nil {
				panic(err)
			}
			_, err = iw.Write(b)
			if err != nil {
				return
			}
		}
	}
}

func StreamAudioHandler(ch <-chan []float32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "audio/wave")
		for frame := range ch {
			err := binary.Write(w, binary.BigEndian, &frame)
			if err != nil {
				panic(err)
			}
		}
	}
}
