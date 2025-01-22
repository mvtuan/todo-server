package user

import (
	"github.com/gin-gonic/gin"
	"server/pkg/common"
	"server/pkg/helpers"
	"server/pkg/models"
)

func (h *UserHandler) getUser(c *gin.Context) {
	id := helpers.ParseInt(c.Param("id"), 0)
	if id == 0 {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
		return
	}

	common.Respond(c, h.service.GetUser(id))
}

func (h *UserHandler) createUser(c *gin.Context) {
	var input *models.User
	if err := c.Bind(&input); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
		return
	}

	common.Respond(c, h.service.CreateUser(input))
}
