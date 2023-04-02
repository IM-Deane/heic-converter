package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IM-Deane/heic-converter/util"
)

func Health(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    status := "OK"
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": status})
}


func ConvertImage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	// for now we support png and jpeg
	convertToPng := r.URL.Query().Get("format") == "png"


    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Missing image file", http.StatusBadRequest)
        return
    }

    defer file.Close()

    imageBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Unable to read image file", http.StatusInternalServerError)
        return
    }

    contentType := http.DetectContentType(imageBytes)

    switch contentType {
    case "image/heic":
        var imgBytes []byte
        var imgContentType string
        var err error

        if convertToPng {
            imgBytes, err = util.ToPng(imageBytes)
            imgContentType = "image/png"
        } else {
            imgBytes, err = util.ToJpeg(imageBytes)
            imgContentType = "image/jpeg"
        }

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", imgContentType)
        w.Write(imgBytes)

    default:
        http.Error(w, fmt.Sprintf("Unsupported image format: %s", contentType), http.StatusBadRequest)
    }
}