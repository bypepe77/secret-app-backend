package secret

import (
	"net/http"
	"strconv"

	"github.com/bypepe77/secret-app-backend/internal/common"
	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	errorGettingMySecrets = "Error fething your secrets"
)

func (ctrl *secretController) GetMySecrets(c *gin.Context) {
	claims, err := common.GetClaims(c)
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	pagination := &Pagination{Limit: limit, Offset: offset}

	secrets, err := ctrl.repository.GetMySecrets(claims.UserID, pagination)
	if err != nil {
		utils.SendErrorResponse(c, errorGettingMySecrets)
	}

	utils.SendResponse(c, http.StatusOK, toSecretResponseArray(secrets))
}
