package model

import (
	"nku-treehole-server/config"
	"nku-treehole-server/db"
	"testing"
)

// 运行测试前得初始化环境
func setup() {
	config.Init("../conf", true)
	db.InitDB()
}

func TestUser_CreateUser(t *testing.T) {
	setup()
	u := &User{}
	t.Log(u.CreateUser(u))
}

func TestUser_SearchUserByID(t *testing.T) {
	setup()
	uid := int64(1395197430445641728)
	u := &User{}
	u, err := u.SearchUserByID(uid)
	t.Log(u, err)
}
