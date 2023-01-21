package secret

import (
	"github.com/bypepe77/secret-app-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type SecretControllerInterface interface {
	Create(c *gin.Context)
	GetOne(c *gin.Context)
	GetMySecrets(c *gin.Context)
	LikeSecret(c *gin.Context)
	UnlikeSecret(c *gin.Context)
	GetSecrets(c *gin.Context)
}

type SecretRepositoryInterface interface {
	Create(p *SecretPayload) (*models.Confession, error)
	GetByIDWithUser(id int) (*models.Confession, error)
	GetMySecrets(userID int, pagination *Pagination) ([]*models.Confession, error)
	HasLiked(confessionID int, userID int) (bool, error)
	AddLikeToConfession(confessionID int, userID int) error
	DeleteLikeFromConfession(confessionID int, userID int) error
	GetSecrets() ([]*models.Confession, error)
}
