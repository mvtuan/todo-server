package user

import (
	"server/pkg/common"
	"server/pkg/models"
)

func (s *UserService) CreateUser(input *models.User) *common.APIResponse {
	return s.db.Create(&input)
}
