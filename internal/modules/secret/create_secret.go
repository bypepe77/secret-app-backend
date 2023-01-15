package secret

import (
	"fmt"
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/common"
	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	errorParsingData     = "Internal error parsing data"
	internalError        = "Internal Error"
	contentEmpty         = "Content cannot be empty"
	errorCreatingContent = "Error creating content"
	confesionCreated     = "Confesion created"
)

func (ctrl *secretController) Create(c *gin.Context) {
	var payload *SecretPayload
	err := c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorResponse(c, errorParsingData)
		return
	}

	claims, err := common.GetClaims(c)
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}

	if payload.Content == "" {
		utils.SendErrorResponse(c, contentEmpty)
		return
	}

	payload.userID = claims.UserID
	_, err = ctrl.repository.Create(payload)
	if err != nil {
		fmt.Println("err", err)
		utils.SendErrorResponse(c, errorCreatingContent)
		return
	}
	utils.SendResponse(c, http.StatusOK, confesionCreated)
}
