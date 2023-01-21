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
	likeNotExist = "Like doesn't exist"
	likeDeleted  = "Unliked this confession"
)

func (ctrl *secretController) UnlikeSecret(c *gin.Context) {
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
	fmt.Print("secret", secret)
	if err != nil {
		utils.SendErrorResponse(c, "test")
		return
	}

	if secret == false {
		utils.SendErrorResponse(c, likeNotExist)
		return
	}

	err = ctrl.repository.DeleteLikeFromConfession(secretID, claims.UserID)
	if err != nil {
		utils.SendErrorResponse(c, errorSettingLike)
		return
	}

	utils.SendResponse(c, http.StatusOK, likeDeleted)
}
