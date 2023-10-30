package service

import (
	"testing"

	"nku-treehole-server/config"
	"nku-treehole-server/model"
)

func setup() {
	config.Setup("../../.env.example")
}

func TestUserService_AddSession(t *testing.T) {
	setup()
	user := &model.User{ID: int64(1398604583239946240)}
	//session := &model.Session{}
	//s, err := session.GetSessionByUid(int64(1398604583239946240))
	//t.Log(s.DeletedAt, s)
	//t.Log(user, err)
	service := &UserService{}
	t.Log(service.AddSession(user))
}
