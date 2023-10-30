package model

import (
	"testing"

	"nku-treehole-server/config"
)

// Initialize the environment before running the test
func setup() {
	config.Setup("../../.env.example")
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

func TestUser_FindInBatches(t *testing.T) {
	setup()
	uid := int64(1398604583239946240)
	u := &User{}
	users, err := u.FindInBatches([]int64{uid})
	t.Log(users, err)
}
