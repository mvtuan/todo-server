package models

import "time"

type CommonFields struct {
	AndCondition []Condition
	OrCondition  []Condition
}

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type CommonModel struct {
	ID        uint       `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
