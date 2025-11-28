package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStream(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}
