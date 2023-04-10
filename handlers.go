package main

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "ok",
    })
}

type ImageResult struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data,omitempty"`
	Error    string `json:"error,omitempty"`
}

func EventProgressGET(c *gin.Context) {
	fileId := c.Query("fileId")

    c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")
    c.Header("Access-Control-Allow-Origin", "*")

    // Create a new channel to receive updates
    updateChan := make(chan int)
    defer close(updateChan)

    // Register the channel with the fileId
    progressMutex.Lock()
    progressChannels[fileId] = updateChan
    progressMutex.Unlock()

    // Unregister the channel when the function exits
    defer func() {
        progressMutex.Lock()
        delete(progressChannels, fileId)
        progressMutex.Unlock()
    }()

    for {
        select {
        case progress, ok := <-updateChan:
            if !ok {
                return
            }
            c.SSEvent("message", progress) // emit progress event
            c.Writer.Flush()
        case <-c.Request.Context().Done():
            return
        }
    }
}

func convertImagePOST(c *gin.Context) {
	convertToPng := c.DefaultPostForm("convertToFormat", "jpeg")
	files := c.Request.MultipartForm.File["images"]
	fileIds := c.PostFormArray("fileIds")

	if len(files) != len(fileIds) {
		c.String(http.StatusBadRequest, "Mismatch between files and fileIds")
		return
	}
	startTime := time.Now()
	convertedImages := make([]ImageResult, len(files))
	for i, file := range files {
		fileId := fileIds[i]
		updateProgress(fileId, 0)
		convertedImages[i] = processImage(file, convertToPng, fileId)
	}

	// round to 10 milliseconds (two decimal places)
	elapsedTime := time.Since(startTime).Truncate(time.Millisecond * 10)
	c.Writer.Header().Set("Server-Timing", elapsedTime.String())

	// Prepare the JSON response
	responseData := make([]map[string]interface{}, len(convertedImages))
	for i, imgResult := range convertedImages {
		if imgResult.Error != "" {
			responseData[i] = map[string]interface{}{
				"error":    true,
				"errorMsg": imgResult.Error,
			}
		} else {
			responseData[i] = map[string]interface{}{
				"error":    false,
				"filename": imgResult.Filename,
				"data":     base64.StdEncoding.EncodeToString(imgResult.Data),
			}
		}
	}
	c.JSON(http.StatusOK, responseData)
}