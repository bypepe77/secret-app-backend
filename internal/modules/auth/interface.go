package auth

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
