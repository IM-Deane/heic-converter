package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"image/jpeg"
	"image/png"

	"github.com/adrium/goheif"
	"github.com/h2non/bimg"
	"github.com/rwcarlsen/goexif/exif"
)

func processHEICImage(imageBytes []byte) (image.Image, error) {
	reader := bytes.NewReader(imageBytes)

	heifImg, err := goheif.Decode(reader)
	if err != nil {
		return nil, err
	}

	// Extract EXIF data and image orientation
	exifData, _ := exif.Decode(reader)
	if exifData != nil {
		orientationVal, err := exifData.Get(exif.Orientation)
		
		if err == nil && orientationVal != nil {
			orientation, _ := orientationVal.Int(0)
			heifImg = applyOrientation(heifImg, orientation) // keep OG image orientation
		}
	}
	return heifImg, nil
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
	convertToFormat string,
	imageQuality string,
	fileId string) ImageResult {
	src, err := file.Open()
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    "Unable to read image file",
		}
	}
	defer src.Close()

	updateProgress(fileId, 10)

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    "Unable to read image file",
		}
	}

	updateProgress(fileId, 20)

	contentType := http.DetectContentType(imageBytes)

	var img image.Image

	if contentType == "image/heic" || contentType == "application/octet-stream" {
		img, err = processHEICImage(imageBytes)
	} else {
		img, err = processNonHEICImage(imageBytes)
	}

	updateProgress(fileId, 50)

	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	// Encode image.Image into a byte slice
	var buf bytes.Buffer
	if convertToFormat == "png" {
		err = png.Encode(&buf, img)
	} else {
		err = jpeg.Encode(&buf, img, nil)
	}

	updateProgress(fileId, 80)

	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	encodedImageBytes := buf.Bytes()
	imageQualityInt, _ := strconv.Atoi(imageQuality)

	// Create options for the output image format
	var outputOptions bimg.Options
	if  convertToFormat == "png" {
		outputOptions = bimg.Options{Type: bimg.PNG, Quality: imageQualityInt, Lossless: true}
	} else {
		// if image quality is less than 75 we can use lossless compression
		if imageQualityInt <= 75 {
			outputOptions = bimg.Options{Type: bimg.JPEG, Quality: imageQualityInt, Lossless: true}
		} else {
			outputOptions = bimg.Options{Type: bimg.JPEG, Quality: imageQualityInt}
		}
	}

	// Process the image
	imgBytes, err := bimg.NewImage(encodedImageBytes).Process(outputOptions)
	if err != nil {
		return ImageResult{
			Filename: file.Filename,
			Error:    err.Error(),
		}
	}

	updateProgress(fileId, 90)

	filename := removeFileExtension(file.Filename)

	if  convertToFormat == "png" {
		filename = fmt.Sprintf("%s.png", filename)
    } else {
		filename = fmt.Sprintf("%s.jpeg", filename)
	}

	updateProgress(fileId, 100)

	return ImageResult{
		Filename: filename,
		Data:     imgBytes,
	}
}

func applyOrientation(img image.Image, orientation int) image.Image {
	bimgOptions := bimg.Options{}

	switch orientation {
	case 2:
		bimgOptions.Flip = true
	case 3:
		bimgOptions.Rotate = 180
	case 4:
		bimgOptions.Rotate = 180
		bimgOptions.Flip = true
	case 5:
		bimgOptions.Rotate = 270
		bimgOptions.Flip = true
	case 6:
		bimgOptions.Rotate = 270
	case 7:
		bimgOptions.Rotate = 90
		bimgOptions.Flip = true
	case 8:
		bimgOptions.Rotate = 90
	}

	if orientation > 1 {
		bimgBuf, _ := bimg.NewImage(imageToBytes(img)).Process(bimgOptions)
		img, _, _ = image.Decode(bytes.NewReader(bimgBuf))
	}

	return img
}

func imageToBytes(img image.Image) []byte {
	var buf bytes.Buffer
	_ = jpeg.Encode(&buf, img, nil)
	return buf.Bytes()
}

func removeFileExtension(filename string) string {
    return filename[0 : len(filename)-len(filepath.Ext(filename))]
}