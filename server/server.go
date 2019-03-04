package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Root path
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server started")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Post WebHook
	r.POST("/webhook", func(c *gin.Context) {
		c.String(http.StatusOK, "post webhook")
	})

	// Onboarding
	r.POST("/onboard", func(c *gin.Context) {
		c.String(http.StatusOK, "post onboard")
	})

	// Offboarding
	r.POST("/offboard", func(c *gin.Context) {
		c.String(http.StatusOK, "post offboard")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
