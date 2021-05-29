package service

import (
	"fmt"
	"nku-treehole-server/model"
	"nku-treehole-server/pkg/jwt"
	"nku-treehole-server/pkg/logger"
	"time"
)

type UserService struct {
}

/*
每次主动登陆新增一条session
写入session表并返回token
*/
func (s *UserService) AddSession(user *model.User) (token string, err error) {
	ss := &model.Session{}
	// 先将该用户之前的session 清除
	err = ss.DeleteOldSession(user.ID)
	if err != nil {
		logger.Errorf("Login err = %v", err)
		return "", err
	}

	token, err = jwt.GenerateToken(fmt.Sprint(user.ID))
	if err != nil {
		logger.Errorf("Login err = %v", err)
		return "", err
	}

	expireAt := time.Now().Add(3 * 24 * 60 * time.Minute) //设为三天后过期
	err = ss.CreateSession(user.ID, token, expireAt)
	if err != nil {
		logger.Errorf("Login err = %v", err)
		return "", err
	}
	return token, nil
}

// 写入session表并返回token
func (s *UserService) Register(user *model.User) (string, error) {
	return "", nil
}
