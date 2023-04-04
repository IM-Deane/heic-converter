package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	config := cors.DefaultConfig()
	// TODO: enable CORS for swiftconvert.io or whatever when released
	config.AllowAllOrigins = true
  	// config.AllowOrigins = []string{"http://localhost:3000"}

  	router.Use(cors.New(config))
	router.Static("/static", "resources/static")

	router.GET("/", index)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  	router.MaxMultipartMemory = 5 << 20  // 5 MiB
	router.POST("/api/convert", ConvertImagePOST)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run("localhost:" + port); err != nil {
        log.Panicf("error: %s", err)
	}
}