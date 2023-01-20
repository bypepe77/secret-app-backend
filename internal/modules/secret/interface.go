package secret

import "github.com/gin-gonic/gin"

type SecretControllerInterface interface {
	Create(c *gin.Context)
	GetOne(c *gin.Context)
	GetMySecrets(c *gin.Context)
}
