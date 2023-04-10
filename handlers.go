package main

import (
	"encoding/base64"
	"io"
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

	if fileId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing fileId parameter"})
		return
	}

	channel := make(chan string)
	addClient(fileId, channel)
	defer removeClient(fileId, channel)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-channel)
		return true
	})
}

func convertImagePOST(c *gin.Context) {
	convertToPng := c.DefaultPostForm("convertToPng", "false") == "true"
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