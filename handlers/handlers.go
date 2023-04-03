package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IM-Deane/heic-converter/util"
)

func enableCors(w *http.ResponseWriter) {
    // allow all during development
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func Health(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    fmt.Printf("got / health request\n")
    
    if (*r).Method == "OPTIONS" {
        return
   } else if (*r).Method == http.MethodGet {
        status := "OK"
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": status})
        return
    }
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}

func ConvertImage(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    if (*r).Method == "OPTIONS" {
        return
   } else if (*r).Method == http.MethodPost {
        fmt.Printf("got / convert request: \n")

        // for now we support png and jpeg
        convertToPng := r.URL.Query().Get("format") == "png"

        file, _, err := r.FormFile("image")
        if err != nil {
            http.Error(w, "Missing image file", http.StatusBadRequest)
            return
        }

        fmt.Printf("got / file: \n")

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
                fmt.Printf("got / converted image to jpeg: \n")
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
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}