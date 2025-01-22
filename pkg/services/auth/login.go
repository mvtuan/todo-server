package auth

import (
	"server/pkg/common"
	"server/pkg/helpers"
	"server/pkg/jwt"
	"server/pkg/models"
	"time"
)

func (h *AuthService) Login(input *models.User) *common.APIResponse {
	user := models.User{
		Email: input.Email,
	}
	existUser := h.db.QueryOne(&user)
	if existUser.Status != common.APIStatus.Ok {
		return existUser
	}

	if err := helpers.ComparePassword(input.Password, user.Password); err != nil {
		return &common.APIResponse{
			Status:    common.APIStatus.InternalServerError,
			Message:   "Invalid email or password",
			RootCause: err,
		}
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return &common.APIResponse{
			Status:    common.APIStatus.InternalServerError,
			Message:   "Something went wrong",
			RootCause: err,
		}
	}

	now := time.Now()
	h.db.Update(&models.User{CommonModel: models.CommonModel{ID: user.ID}}, &models.User{LastLoggedIn: &now})
	user.Token = token

	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   user,
	}
}

func (h *AuthService) Register(input *models.User) *common.APIResponse {
	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return &common.APIResponse{
			Status:    common.APIStatus.InternalServerError,
			Message:   "Something went wrong",
			RootCause: err,
		}
	}

	input.Password = hashedPassword

	return h.db.Create(&input)
}
