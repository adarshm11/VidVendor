package routes

import (
	"github.com/gin-gonic/gin"

	"VidVendor/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/enqueue", handlers.EnqueueVideo)
	r.POST("/stop", handlers.StopStream)
	r.POST("/skip", handlers.SkipStream)
	r.GET("/status", handlers.GetStreamStatus)
	r.GET("/stream", handlers.GetStream)
}
