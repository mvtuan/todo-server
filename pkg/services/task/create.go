package task

import (
	"server/pkg/common"
	"server/pkg/models"
)

func (ts *TaskService) CreateTask(task *models.Task) *common.APIResponse {
	return ts.db.Create(&task)
}
