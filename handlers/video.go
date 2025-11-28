package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnqueueVideo(c *gin.Context) {

}

func GetStreamStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "streaming"})
}

func StopStream(c *gin.Context) {

}

func SkipStream(c *gin.Context) {

}
