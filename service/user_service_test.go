package service

import (
	"nku-treehole-server/config"
	"nku-treehole-server/db"
	"nku-treehole-server/model"
	"testing"
)

func setup() {
	config.Init("../conf", false)
	db.InitDB()
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
