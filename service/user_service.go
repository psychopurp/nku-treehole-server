package service

import (
	"fmt"
	"nku-treehole-server/config"
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

	expireAt := time.Now().Add(config.EXPIRE_DURATION) //设为三天后过期
	err = ss.CreateSession(user.ID, token, expireAt)
	if err != nil {
		logger.Errorf("Login err = %v", err)
		return "", err
	}
	return token, nil
}

// 检查token 是否过期，如果没过期则刷新token
func (s *UserService) CheckExpireAndRefresh(token string) error {
	ss := &model.Session{}
	session, err := ss.GetSessionByToken(token)
	if err != nil {
		logger.Errorf("Session 获取失败")
		return fmt.Errorf("Session 获取失败 err=%v ", err)
	}
	if !session.ExpiredAt.After(time.Now()) {
		logger.Errorf("token 已过期")
		return fmt.Errorf("token 已过期")
	}

	err = ss.Refresh(token, time.Now().Add(config.EXPIRE_DURATION))
	if err != nil {
		logger.Errorf("token 更新失败 err=%v ", err)
	}

	return nil
}
