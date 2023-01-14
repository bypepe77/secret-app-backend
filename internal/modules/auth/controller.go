package auth

import (
	"github.com/bypepe77/secret-app-backend/internal/common"
	"gorm.io/gorm"
)

type authController struct {
	db         *gorm.DB
	jwtService common.JWTService
	repository UserRepositoryInterface
}

func NewAuthController(db *gorm.DB) AuthControllerInterface {
	return &authController{
		db:         db,
		jwtService: common.JWTAuthService(),
		repository: NewUserRepository(db),
	}
}
