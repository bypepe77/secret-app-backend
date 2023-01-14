package auth

import (
	"github.com/bypepe77/secret-app-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Exists(username string) (bool, error)
	Create(payload *UserPayload) (models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		DB: db,
	}
}

func (repository *userRepository) Create(payload *UserPayload) (models.User, error) {
	var user models.User
	user.Username = payload.Username
	user.Email = payload.Email
	user.Password = payload.Password

	if err := repository.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *userRepository) Exists(username string) (bool, error) {
	var user models.User
	if err := repository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
