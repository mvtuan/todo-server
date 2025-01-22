package auth

import "github.com/gin-gonic/gin"

func (h *AuthHandler) registerGroup() *gin.RouterGroup { return h.serve.Gin.Group(h.group) }

func (h *AuthHandler) routes() {
	h.router.POST("/register", h.register)
	h.router.POST("/login", h.login)
}
