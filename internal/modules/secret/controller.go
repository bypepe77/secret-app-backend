package secret

import (
	"gorm.io/gorm"
)

type secretController struct {
	db         *gorm.DB
	repository SecretRepositoryInterface
}

func NewSecretController(db *gorm.DB) SecretControllerInterface {
	return &secretController{
		db:         db,
		repository: NewSecretRepository(db),
	}
}
