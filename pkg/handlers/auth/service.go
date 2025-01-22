package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"server/pkg/handlers"
	"server/pkg/services/auth"
)

type AuthHandler struct {
	serve    *handlers.Server
	group    string
	router   gin.IRoutes
	service  *auth.AuthService
	validate *validator.Validate
}

func NewAuthHandler(s *handlers.Server, groupName string, service *auth.AuthService, validate *validator.Validate) {

	taskHandler := &AuthHandler{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
		validate,
	}
	taskHandler.router = taskHandler.registerGroup()
	taskHandler.routes()
	//usrHandler.registerValidator()
}
