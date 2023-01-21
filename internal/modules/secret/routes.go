package secret

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SecretRoute struct {
	db         *gorm.DB
	routeGroup gin.RouterGroup
	ctrl       SecretControllerInterface
}

func NewSecretRoute(db *gorm.DB, r gin.RouterGroup) *SecretRoute {
	return &SecretRoute{
		db:         db,
		routeGroup: r,
		ctrl:       NewSecretController(db),
	}
}

func (routeController *SecretRoute) RegisterSecretRoutes() {
	routeController.routeGroup.POST("/new", routeController.ctrl.Create)
	routeController.routeGroup.GET("/get/:SecretID", routeController.ctrl.GetOne)
	routeController.routeGroup.GET("/get/mySecrets", routeController.ctrl.GetMySecrets)
	routeController.routeGroup.GET("/like/:SecretID", routeController.ctrl.LikeSecret)
	routeController.routeGroup.GET("/unlike/:SecretID", routeController.ctrl.UnlikeSecret)
}
