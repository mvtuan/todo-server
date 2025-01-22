package task

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/pkg/common"
	"server/pkg/helpers"
	"server/pkg/models"
)

func (t *TaskHandler) createTask(c *gin.Context) {
	var input *models.Task
	if err := c.Bind(&input); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Bad request",
			RootCause: err,
		})
		return
	}

	common.Respond(c, t.service.CreateTask(input))
}

func (t *TaskHandler) getTasks(c *gin.Context) {
	offset := helpers.ParseInt(c.Query("offset"), 0)
	limit := helpers.ParseInt(c.Query("limit"), 20)
	q := c.Query("q")
	var query []*models.Task

	if q != "" {
		var input *models.Task
		err := json.Unmarshal([]byte(q), &input)
		if err != nil {
			common.Respond(c, &common.APIResponse{
				Status:  common.APIStatus.Invalid,
				Message: "Invalid query parameter",
			})
			return
		}

		query = append(query, input)
	}

	common.Respond(c, t.service.GetTasks(query, offset, limit))
}

func (t *TaskHandler) updateTask(c *gin.Context) {
	id := helpers.ParseInt(c.Param("id"), 0)
	if id == 0 {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Missing task id",
		})
		return
	}
	filter := &models.Task{Model: gorm.Model{ID: uint(id)}}

	var updater *models.Task
	if err := c.Bind(&updater); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
	}

	common.Respond(c, t.service.UpdateTask(filter, updater))
}

func (t *TaskHandler) updateTaskStatus(c *gin.Context) {
	id := helpers.ParseInt(c.Param("id"), 0)
	if id == 0 {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Missing task id",
		})
		return
	}
	filter := &models.Task{Model: gorm.Model{ID: uint(id)}}

	var updater *models.Task
	if err := c.Bind(&updater); err != nil {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Bad request",
		})
		return

	}

	if updater.Status == "" {
		common.Respond(c, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Missing updater status",
		})
		return
	}

	common.Respond(c, t.service.UpdateTask(filter, updater))
}
