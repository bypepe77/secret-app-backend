package auth

import (
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	errorParsingUserData = "Internal error parsing user data."
	errorIncorrectData   = "Email, password or username cannot be empty"
	errorCreatingUser    = "Error creating user, please try again"
	userExist            = "This username already exist"
)

func (ctrl *authController) Register(c *gin.Context) {
	var payload *UserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorResponse(c, errorParsingUserData)
		return
	}

	if payload.Email == "" || payload.Username == "" || payload.Password == "" {
		utils.SendErrorResponse(c, errorIncorrectData)
	}

	ifExistUser, err := ctrl.repository.Exists(payload.Username)
	if err != nil {
		utils.SendErrorResponse(c, errorCreatingUser)
		return
	}

	if ifExistUser {
		utils.SendErrorResponse(c, userExist)
		return
	}

	user, err := ctrl.repository.Create(payload)
	if err != nil {
		utils.SendErrorResponse(c, errorCreatingUser)
		return
	}

	token, err := ctrl.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		utils.SendErrorResponse(c, errorCreatingUser)
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
