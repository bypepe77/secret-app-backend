package secret

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bypepe77/secret-app-backend/internal/common"
	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	likeExist        = "Already liked this confession"
	errorSettingLike = "Error creating like"
	created          = "Like created"
)

func (ctrl *secretController) LikeSecret(c *gin.Context) {
	id := c.Param("SecretID")

	secretID, err := strconv.Atoi(id)
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}

	claims, err := common.GetClaims(c)
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}

	secret, err := ctrl.repository.HasLiked(secretID, claims.UserID)
	if err != nil {
		utils.SendErrorResponse(c, likeExist)
		return
	}

	if secret {
		utils.SendErrorResponse(c, likeExist)
		return
	}

	err = ctrl.repository.AddLikeToConfession(secretID, claims.UserID)
	if err != nil {
		fmt.Println("err", err)
		utils.SendErrorResponse(c, errorSettingLike)
		return
	}

	utils.SendResponse(c, http.StatusOK, created)
}
