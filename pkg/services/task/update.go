package task

import (
	"server/pkg/common"
	"server/pkg/models"
)

func (ts *TaskService) UpdateTask(filter, updater *models.Task) *common.APIResponse {
	return ts.db.Update(&filter, &updater)
}
