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

	// Define a simple GET endpoint
	r.GET("/ping", Pong)

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
