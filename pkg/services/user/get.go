package user

import (
	"server/pkg/common"
	"server/pkg/models"
)

func (s *UserService) GetUser(id int) *common.APIResponse {
	return s.db.Query(&models.User{CommonModel: models.CommonModel{ID: uint(id)}}, 0, 1)
}
