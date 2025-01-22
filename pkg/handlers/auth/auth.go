package auth

import (
	"github.com/gin-gonic/gin"
	"server/pkg/common"
	"server/pkg/models"
)

func (h *AuthHandler) login(c *gin.Context) {
	var input models.User
	if err := c.Bind(&input); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
		return
	}

	common.Respond(c, h.service.Login(&input))
}

func (h *AuthHandler) register(c *gin.Context) {
	var input models.User
	if err := c.Bind(&input); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
		return
	}

	if input.Email == "" || input.Password == "" {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Missing email or password",
		})
		return
	}

	common.Respond(c, h.service.Register(&input))
}
