package service

import (
	"fmt"
	"time"

	"nku-treehole-server/config"
	"nku-treehole-server/model"
	"nku-treehole-server/pkg/jwt"
	"nku-treehole-server/pkg/logger"
)

type UserService struct{}

// Add a new session for each login and return a token by writing it to the session table.
func (s *UserService) AddSession(user *model.User) (token string, err error) {
	ss := &model.Session{}
	// First, clear the user's previous sessions.
	err = ss.DeleteOldSession(user.ID)
	if err != nil {
		logger.Errorf("login error: %v", err)
		return "", err
	}

	token, err = jwt.GetJWTCrypto().GenerateToken(fmt.Sprint(user.ID))
	if err != nil {
		logger.Errorf("login error: %v", err)
		return "", err
	}

	expireAt := time.Now().Add(config.EXPIRE_DURATION) // Set to expire in three days
	err = ss.CreateSession(user.ID, token, expireAt)
	if err != nil {
		logger.Errorf("login error: %v", err)
		return "", err
	}
	return token, nil
}

// Check if the token has expired and refresh it if it hasn't.
func (s *UserService) CheckExpireAndRefresh(token string) error {
	ss := &model.Session{}
	session, err := ss.GetSessionByToken(token)
	if err != nil {
		logger.Errorf("failed to retrieve session")
		return fmt.Errorf("failed to retrieve session: err=%v", err)
	}
	if !session.ExpiredAt.After(time.Now()) {
		logger.Errorf("token has expired")
		return fmt.Errorf("token has expired")
	}

	err = ss.Refresh(token, time.Now().Add(config.EXPIRE_DURATION))
	if err != nil {
		logger.Errorf("token update failed: err=%v", err)
	}

	return nil
}
