package util

import (
	"bytes"
	"fmt"
	"net/http"

	"image/jpeg"
	"image/png"

	"github.com/adrium/goheif"
	"github.com/pkg/errors"
)



func ToJpeg(imageBytes []byte) ([]byte, error) {
    contentType := http.DetectContentType(imageBytes)

    switch contentType {
    case "image/heic":
    case "application/octet-stream":
        img, err := goheif.Decode(bytes.NewReader(imageBytes))
        if err != nil {
            return nil, errors.Wrap(err, "unable to decode heic")
        }

        buf := new(bytes.Buffer)
        if err := jpeg.Encode(buf, img, nil); err != nil {
            return nil, errors.Wrap(err, "unable to encode jpeg")
        }
        return buf.Bytes(), nil
    }

    return nil, fmt.Errorf("unable to convert %#v to jpeg", contentType)
}

func ToPng(imageBytes []byte) ([]byte, error) {
    contentType := http.DetectContentType(imageBytes)

    switch contentType {
    case "image/heic":
    case "application/octet-stream":
        img, err := goheif.Decode(bytes.NewReader(imageBytes))
        if err != nil {
            return nil, errors.Wrap(err, "unable to decode heic")
        }

        buf := new(bytes.Buffer)
        if err := png.Encode(buf, img); err != nil {
            return nil, errors.Wrap(err, "unable to encode png")
        }
        return buf.Bytes(), nil
    }

    return nil, fmt.Errorf("unable to convert %#v to png", contentType)
}





