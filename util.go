package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"image/jpeg"
	"image/png"

	"github.com/adrium/goheif"
	"github.com/h2non/bimg"
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


func processHEICImage(imageBytes []byte) (image.Image, error) {
	img, err := goheif.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func processNonHEICImage(imageBytes []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func processImage(
	file *multipart.FileHeader,
	convertToPng bool,
	fileId string) ImageResult {
	src, err := file.Open()
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    "Unable to read image file",
		}
	}
	defer src.Close()

	progressMap[fileId] = 10

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    "Unable to read image file",
		}
	}

	progressMap[fileId] = 20

	contentType := http.DetectContentType(imageBytes)

	var img image.Image

	if contentType == "image/heic" || contentType == "application/octet-stream" {
		img, err = processHEICImage(imageBytes)
	} else {
		img, err = processNonHEICImage(imageBytes)
	}

	progressMap[fileId] = 50

	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	// Encode image.Image into a byte slice
	var buf bytes.Buffer
	if convertToPng {
		err = png.Encode(&buf, img)
	} else {
		err = jpeg.Encode(&buf, img, nil)
	}
	progressMap[fileId] = 80

	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	encodedImageBytes := buf.Bytes()

	// Create options for the output image format
	var outputOptions bimg.Options
	if convertToPng {
		outputOptions = bimg.Options{Type: bimg.PNG}
	} else {
		outputOptions = bimg.Options{Type: bimg.JPEG}
	}

	// Process the image
	imgBytes, err := bimg.NewImage(encodedImageBytes).Process(outputOptions)
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	filename := removeFileExtension(file.Filename)
	if convertToPng {
		filename = fmt.Sprintf("%s.png", filename)
    } else {
		filename = fmt.Sprintf("%s.jpeg", filename)
	}

	progressMap[fileId] = 100

	return ImageResult{
		Filename: filename,
		Data:     imgBytes,
	}
}

func removeFileExtension(filename string) string {
    return filename[0 : len(filename)-len(filepath.Ext(filename))]
}

func updateProgress(fileId string, progress int) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	progressMap[fileId] = progress
}

func addClient(fileId string, client chan string) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	progressMap[fileId] = 0
	go func() {
		for progress := range progressMap {
			client <- fmt.Sprintf("fileId: %s, progress: %s%%", fileId, progress)
		}
	}()
}

func removeClient(fileId string, client chan string) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	delete(progressMap, fileId)
	close(client)
}
