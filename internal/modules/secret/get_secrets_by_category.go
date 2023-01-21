package secret

import (
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (ctrl *secretController) GetSecretsByCategory(c *gin.Context) {
	categoryName := c.Param("CategoryName")

	secret, err := ctrl.repository.GetSecretsByCategory(categoryName)
	if err != nil {
		utils.SendErrorResponse(c, secretNotFound)
		return
	}

	utils.SendResponse(c, http.StatusOK, toSecretResponseArray(secret))
}
