package task

import "github.com/gin-gonic/gin"

func (t *TaskHandler) registerGroup() *gin.RouterGroup { return t.serve.Gin.Group(t.group) }

func (t *TaskHandler) routes() {
	t.router.POST("/", t.createTask)
	t.router.GET("/", t.getTasks)
	t.router.PUT("/:id", t.updateTask)
	t.router.PUT("/:id/status", t.updateTaskStatus)
}
