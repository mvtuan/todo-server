package user

import (
	"errors"
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

func (h *UserHandler) getMe(c *gin.Context) {
	curUserVal, exists := c.Get("currentUser")
	if !exists {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Unauthorized request",
		})
		return
	}

	curUser, ok := curUserVal.(*models.User)
	if !ok {
		common.Respond(c, &common.APIResponse{
			Status:    common.APIStatus.InternalServerError,
			Message:   "Something went wrong",
			RootCause: errors.New("unable to assert current user"),
		})
		return
	}

	if curUser.ID == 0 {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Unauthorized request",
		})
		return
	}

	meRes := h.service.GetMe(curUser)

	common.Respond(c, meRes)
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
