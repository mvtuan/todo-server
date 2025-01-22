package user

import "server/pkg/models"

type UserService struct {
	db  models.Database
	env string
}

func NewUserService(db models.Database, env string) *UserService { return &UserService{db, env} }
