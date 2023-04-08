package main

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)


func index(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "ok",
    })
}


type ImageResult struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data,omitempty"`
	Error    string `json:"error,omitempty"`
}

func ConvertImagePOST(c *gin.Context) {
    // for now we support png and jpeg
    convertToPng := c.Query("format") == "png"

    file, _, err := c.Request.FormFile("file")
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Missing image file",
		})
        return
    }

    imageBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Unable to read image file",
		})
        return
    }

	startTime := time.Now()
    contentType := http.DetectContentType(imageBytes)

    switch contentType {
    case "image/heic":
    case "application/octet-stream":
        var imgBytes []byte
        var imgContentType string
        var err error

        if convertToPng {
            imgBytes, err = ToPng(imageBytes)
            imgContentType = "image/png"
        } else {
            imgBytes, err = ToJpeg(imageBytes)
            imgContentType = "image/jpeg"
        }

        if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
				"error":  err.Error(),
			})
            return
        }
        // round to 10 milliseconds (two decimal places)
		elapsedTime := time.Since(startTime).Truncate(time.Millisecond * 10)
		c.Writer.Header().Set("Server-Timing", elapsedTime.String())
		c.Data(http.StatusOK, imgContentType, imgBytes)

    default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  fmt.Sprintf("Unsupported image format: %s", contentType),
		})
    }
}

func ConvertPOSTSSE(c *gin.Context) {
	// for now we support png and jpeg
	convertToPng := c.Query("format") == "png"

	// Get the array of files from the request
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Missing image files",
		})
		return
	}

	// Create a WaitGroup to wait for all conversions to complete
	var wg sync.WaitGroup

	// Limit the number of concurrent goroutines
	concurrency := 8
	sem := make(chan struct{}, concurrency)

	// Set the response header for SSE
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// Helper function to send an ImageResult as an SSE event
	sendImageResult := func(result ImageResult) {
		data, _ := json.Marshal(result)
		c.String(http.StatusOK, "event: image\n")
		c.String(http.StatusOK, "data: %s\n\n", data)
		c.Writer.Flush()
	}

	// Iterate through all the files
	for _, file := range form.File["file"] {
		wg.Add(1) // Increment the WaitGroup counter
		sem <- struct{}{}

		// Start a goroutine to handle each file
		go func(file *multipart.FileHeader) {
			defer func() {
				<-sem
				wg.Done() // Decrement the WaitGroup counter when the goroutine is done
			}()

		// Open the file and defer its closing
		fileContents, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "Unable to open image file",
			})
			return
		}
		defer fileContents.Close()

		// Process the image and send the result as an SSE event
		imageBytes, err := io.ReadAll(fileContents)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "Unable to read image file",
			})
			return
		}

		startTime := time.Now()
		contentType := http.DetectContentType(imageBytes)

		switch contentType {
		case "image/heic":
		case "application/octet-stream":
			var imgBytes []byte
			var imgContentType string
			var err error

			if convertToPng {
				imgBytes, err = ToPng(imageBytes)
				imgContentType = "image/png"
			} else {
				imgBytes, err = ToJpeg(imageBytes)
				imgContentType = "image/jpeg"
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "failed",
					"error":  err.Error(),
				})
				return
			}
			// round to 10 milliseconds (two decimal places)
			elapsedTime := time.Since(startTime).Truncate(time.Millisecond * 10)
			c.Writer.Header().Set("Server-Timing", elapsedTime.String())
			c.Data(http.StatusOK, imgContentType, imgBytes)

		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  fmt.Sprintf("Unsupported image format: %s", contentType),
			})
		}
				result := processImage(file, convertToPng)
				sendImageResult(result)
			}(file)
		}

	// Wait for all conversions to complete
	wg.Wait()

	// Send a final event to signal that all images have been processed
	c.String(http.StatusOK, "event: complete\n")
	c.String(http.StatusOK, "data: {}\n\n")
	c.Writer.Flush()
}
