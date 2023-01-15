package secret

import (
	"net/http"
	"strconv"

	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	secretNotFound = "Secret not found"
)

func (ctrl *secretController) GetOne(c *gin.Context) {
	id := c.Param("SecretID")

	secretID, err := strconv.Atoi(id)
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}

	secret, err := ctrl.repository.GetByIDWithUser(secretID)
	if err != nil {
		utils.SendErrorResponse(c, secretNotFound)
		return
	}

	utils.SendResponse(c, http.StatusOK, toSecretResponse(secret))
}
