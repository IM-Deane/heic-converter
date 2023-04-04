package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ConfigRuntime()
	StartGin()
}

func envVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
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
	envVariable("ENV")

	router := gin.Default()
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Server-Timing"}
	if envVariable("ENV") == "production" {
		config.AllowOrigins = []string{envVariable("CLIENT_ORIGIN")}
	} else {
		config.AllowAllOrigins = true
	}
  	
  	router.Use(cors.New(config))
	router.Static("/static", "resources/static")

	router.GET("/", index)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  	router.MaxMultipartMemory = 5 << 20  // 5 MiB
	router.POST("/api/convert", ConvertImagePOST)

	port := envVariable("PORT") 
	if port == "" {
		port = "8080"
	}
	if err := router.Run("localhost:" + port); err != nil {
        log.Panicf("error: %s", err)
	}
}