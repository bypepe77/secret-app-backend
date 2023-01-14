package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoute struct {
	db         *gorm.DB
	routeGroup gin.RouterGroup
	ctrl       AuthControllerInterface
}

func NewUserRoute(db *gorm.DB, r gin.RouterGroup) *UserRoute {
	return &UserRoute{
		db:         db,
		routeGroup: r,
		ctrl:       NewAuthController(db),
	}
}

func (routeController *UserRoute) RegisterUserRoutes() {
	routeController.routeGroup.POST("/register", routeController.ctrl.Register)
}
