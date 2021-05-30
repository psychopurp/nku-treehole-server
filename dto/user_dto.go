package dto

import (
	"fmt"
	"nku-treehole-server/model"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserProfile struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Avatar string `json:"avatar"`
	Level  int    `json:"level"`
	Token  string `json:"token"`
}

func NewUserProfile(user *model.User, token string) *UserProfile {
	sex := "男"
	if user.Sex == 1 {
		sex = "女"
	}
	return &UserProfile{
		Id:     fmt.Sprint(user.ID),
		Name:   user.Name,
		Sex:    sex,
		Avatar: user.Avatar,
		Level:  int(user.Level),
		Token:  token,
	}
}
