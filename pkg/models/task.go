package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model

	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	Priority    int        `json:"priority,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

func (t *Task) TableName() string {
	return "tasks"
}
