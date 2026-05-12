package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaimesHub/go-fresher/internal/routers"
)

func main() {
	// go get -u github.com/gin-gonic/gin
	// Your first Gin Application: https://github.com/gin-gonic/gin#your-first-gin-application
	// To run the server, use:
	// go run cmd/server/main.go
	// Then, you can test the endpoint by visiting http://localhost:8080/ping in your browser or using curl:
	// curl http://localhost:8080/ping

	// // Create a Gin router with default middleware (logger and recovery)
	// r := gin.Default()

	// // Define a simple GET endpoint
	// r.GET("/ping", func(c *gin.Context) {
	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r := routers.NewRouter()

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

// how to get parameters from requests in Gin:
// 1. Query Parameters (URL parameters):
//   - Example: /ping?name=John
//   - Access: c.Query("name") // returns "John"
func QueryHandler(c *gin.Context) {
	name := c.Query("name") // Get the query parameter "name"
	c.JSON(http.StatusOK, gin.H{
		"name": name, // Return the value of "name" in the JSON response
	})
}

// 2. Path Parameters:
//   - Example: /user/:name
//   - Access: c.Param("name") // returns the value of :name
func UserHandler(c *gin.Context) {
	name := c.Param("name") // Get the path parameter "name"
	c.JSON(http.StatusOK, gin.H{
		"user_name": name, // Return the value of "name" in the JSON response
	})
}

// 3. Form Parameters (for POST requests):
//    - Example: form data with key "username"
//    - Access: c.PostForm("username") // returns the value of "username"

// 4. JSON Body:
//    - Example: JSON payload { "name": "John" }
//    - Access: Use c.BindJSON(&struct) to bind the JSON to a struct

// 5. Header Parameters:
//    - Example: Custom header "X-Custom-Header"
//    - Access: c.GetHeader("X-Custom-Header") // returns the value of the header

// DefaultQuery example
func DefaultQueryHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") // Get the query parameter "name" with a default value of "Guest"
	c.JSON(http.StatusOK, gin.H{
		"name": name, // Return the value of "name" in the JSON response
	})
}
