package auth

import (
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	userNotFound  = "User not found"
	incorrectData = "Wrong credentials"
	userBanned    = "This user is banned"
	internalError = "Internal error while trying to generate response"
)

func (ctrl *authController) Login(c *gin.Context) {
	var payload *UserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorResponse(c, errorParsingUserData)
		return
	}

	if payload.Username == "" || payload.Password == "" {
		utils.SendErrorResponse(c, errorIncorrectData)
		return
	}

	user, err := ctrl.repository.GetByUsername(payload.Username)
	if err != nil {
		utils.SendErrorResponse(c, userNotFound)
		return
	}

	if payload.Password != user.Password || payload.Username != user.Username {
		utils.SendErrorResponse(c, incorrectData)
		return
	}

	if user.IsBanned {
		utils.SendErrorResponse(c, userBanned)
		return
	}

	token, err := ctrl.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		utils.SendErrorResponse(c, internalError)
		return
	}

	response := &UserCretedResponse{
		Username: user.Username,
		Email:    user.Email,
		IsBanned: user.IsBanned,
		IsActive: user.IsActive,
		Token:    token,
	}

	utils.SendResponse(c, http.StatusOK, response)
}
