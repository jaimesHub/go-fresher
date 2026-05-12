package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", PongV1)            // Define a GET endpoint at /v1/ping that calls the PongV1 handler
		v1.GET("/user/:name", UserHandler) // Define a GET endpoint at /v1/user/:name that calls the UserHandler
		v1.GET("/query", QueryHandler)     // Define a GET endpoint at /v1/query that calls the QueryHandler
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", PongV2)                       // Define a GET endpoint at /v2/ping that calls the PongV2 handler
		v2.GET("/default-query", DefaultQueryHandler) // Define a GET endpoint at /v2/default-query that calls the DefaultQueryHandler
	}
	return r
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

func UserHandler(c *gin.Context) {
	name := c.Param("name") // Get the path parameter "name"
	c.JSON(http.StatusOK, gin.H{
		"user_name": name, // Return the value of "name" in the JSON response
	})
}

func QueryHandler(c *gin.Context) {
	name := c.Query("name") // Get the query parameter "name"
	c.JSON(http.StatusOK, gin.H{
		"name": name, // Return the value of "name" in the JSON response
	})
}

func DefaultQueryHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") // Get the query parameter "name" with a default value of "Guest"
	c.JSON(http.StatusOK, gin.H{
		"name": name, // Return the value of "name" in the JSON response
	})
}
