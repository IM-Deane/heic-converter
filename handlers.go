package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

	// TODO: can upload the image to server briefly
	// fileExt := filepath.Ext(header.Filename)
	// originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	// now := time.Now()
	// filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	// filePath := "http://localhost:8000/tmp/single/" + filename

	// imageFile, _, err := image.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	// err = imaging.Save(src, fmt.Sprintf("public/single/%v", filename))
	// if err != nil {
	// 	log.Fatalf("failed to save image: %v", err)
	// }

	// c.JSON(http.StatusOK, gin.H{"filepath": filePath})

    // defer file.Close()

    imageBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Unable to read image file",
		})
        return
    }

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

		c.Data(http.StatusOK, imgContentType, imgBytes)

    default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  fmt.Sprintf("Unsupported image format: %s", contentType),
		})
    }

}