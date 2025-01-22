package auth

import "server/pkg/models"

type AuthService struct {
	db  models.Database
	env string
}

func NewAuthService(db models.Database, env string) *AuthService { return &AuthService{db, env} }
