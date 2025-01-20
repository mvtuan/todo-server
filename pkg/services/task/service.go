package task

import "server/pkg/models"

type TaskService struct {
	db  models.Database
	env string
}

func NewTaskService(db models.Database, env string) *TaskService {
	return &TaskService{db: db, env: env}
}
