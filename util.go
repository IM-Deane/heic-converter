package main

import (
	"bytes"
	"fmt"
	"net/http"

	"image/jpeg"
	"image/png"

	"github.com/adrium/goheif"
	"github.com/pkg/errors"
	// "github.com/gin-gonic/gin"
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

// func rateLimit(c *gin.Context) {
// 	ip := c.ClientIP()
// 	value := int(c.ips.Add(ip, 1))
// 	if value%50 == 0 {
// 		fmt.Printf("ip: %s, count: %d\n", ip, value)
// 	}
// 	if value >= 200 {
// 		if value%200 == 0 {
// 			fmt.Println("ip blocked")
// 		}
// 		c.Abort()
// 		c.String(http.StatusServiceUnavailable, "you were automatically banned :)")
// 	}
// }