package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data, "status": true})
}

func SendErrorResponse(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err, "status": false})
}
