package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"server/pkg/handlers"
	"server/pkg/services/user"
)

type UserHandler struct {
	serve   *handlers.Server
	group   string
	router  gin.IRoutes
	service *user.UserService
}

func NewUserHandler(s *handlers.Server, groupName string, service *user.UserService, validate *validator.Validate) {

	userHandler := &UserHandler{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
	}
	userHandler.router = userHandler.registerGroup()
	userHandler.routes()
	//usrHandler.registerValidator()
}
