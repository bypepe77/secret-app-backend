package secret

import (
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (ctrl *secretController) GetSecrets(c *gin.Context) {
	secret, err := ctrl.repository.GetSecrets()
	if err != nil {
		utils.SendErrorResponse(c, secretNotFound)
		return
	}
	utils.SendResponse(c, http.StatusOK, toSecretResponseArray(secret))
}
