package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// go get -u github.com/gin-gonic/gin
	// Your first Gin Application: https://github.com/gin-gonic/gin#your-first-gin-application
	// To run the server, use:
	// go run cmd/server/main.go
	// Then, you can test the endpoint by visiting http://localhost:8080/ping in your browser or using curl:
	// curl http://localhost:8080/ping

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", PongV1) // Define a GET endpoint at /v1/ping that calls the PongV1 handler
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", PongV2) // Define a GET endpoint at /v2/ping that calls the PongV2 handler
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func Pong(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func PongV1(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong v1",
	})
}

func PongV2(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong v2",
	})
}
