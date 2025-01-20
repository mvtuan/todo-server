package task

import (
	"server/pkg/common"
	"server/pkg/models"
)

func (ts *TaskService) GetTasks(filter []*models.Task, offset, limit int) *common.APIResponse {
	return ts.db.Query(&filter, offset, limit)
}
