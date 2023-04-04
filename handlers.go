package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func index(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "ok",
    })
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

    imageBytes, err := ioutil.ReadAll(file)
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

		elapsedTime := time.Since(startTime)
		c.Writer.Header().Set("Server-Timing", elapsedTime.String())
		c.Data(http.StatusOK, imgContentType, imgBytes)

    default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  fmt.Sprintf("Unsupported image format: %s", contentType),
		})
    }

}