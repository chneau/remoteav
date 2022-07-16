package common

import (
	"bytes"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
)

func StreamHandlerfunc(ch chan *image.YCbCr) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const boundary = `frame`
		w.Header().Set("Content-Type", `multipart/x-mixed-replace;boundary=`+boundary)
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
