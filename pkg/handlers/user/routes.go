package user

import (
	"github.com/gin-gonic/gin"
	"server/pkg/middlewares"
)

func (h *UserHandler) registerGroup() *gin.RouterGroup { return h.serve.Gin.Group(h.group) }

func (h *UserHandler) routes() {
	h.router.Use(middlewares.AuthenticateRequest())

	h.router.GET("/:id", h.getUser)
	h.router.POST("/", h.createUser)
}
