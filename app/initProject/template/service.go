package template

var ServiceUserTmp = `package service

import (
	"context"
	"{{projectName}}/dao"
	"{{projectName}}/model"
	"{{projectName}}/model/entity"
)

var insUser = &SvrUser{}

func User() *SvrUser {
	return insUser
}

type SvrUser struct{}

func (s *SvrUser) GetUserByUserName(ctx context.Context, username string) (out *model.UserOutput, err error) {
	var user entity.User
	if err = dao.User.Ctx(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return
	}

	out = &model.UserOutput{
		Username: user.Username,
		Password: user.Password,
	}
	return
}`
