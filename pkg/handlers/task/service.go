package task

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"server/pkg/handlers"
	"server/pkg/services/task"
)

type TaskHandler struct {
	serve    *handlers.Server
	group    string
	router   gin.IRoutes
	service  *task.TaskService
	validate *validator.Validate
}

func NewTaskHandler(s *handlers.Server, groupName string, service *task.TaskService, validate *validator.Validate) {

	taskHandler := &TaskHandler{
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

//func (s *UserHandler) registerValidator() {
//	_ = s.validate.RegisterValidation("name", userValidate.NameValidator)
//	_ = s.validate.RegisterValidation("email", userValidate.EmailValidator)
//}
