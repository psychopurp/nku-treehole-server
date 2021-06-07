package model

import (
	"nku-treehole-server/db"
	"time"
)

type User struct {
	ID        int64      `gorm:"column:id" json:"id" form:"id"`
	Level     int32      `gorm:"column:level" json:"level" form:"level"`
	Name      string     `gorm:"column:name" json:"name" form:"name"`
	Sex       int32      `gorm:"column:sex" json:"sex" form:"sex"`
	Birthday  *time.Time `gorm:"column:birthday" json:"birth" form:"birthday"`
	Email     string     `gorm:"column:email" json:"email" form:"email"`
	Avatar    string     `gorm:"column:avatar" json:"avatar" form:"avatar"`
	Phone     string     `gorm:"column:phone" json:"phone" form:"phone"`
	Password  string     `gorm:"column:password" json:"password" form:"password"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) CreateUser(user *User) error {
	conn := db.GetDBConn()
	err := conn.Table(u.TableName()).Create(user).Error
	return err
}

func (u *User) SearchUserByID(uid int64) (*User, error) {
	conn := db.GetDBConn()
	res := &User{}
	err := conn.Table(u.TableName()).Where("id=?", uid).First(res).Error
	return res, err
}

func (u *User) FindInBatches(userIds []int64) ([]*User, error) {
	conn := db.GetDBConn()
	var res []*User
	err := conn.Table(u.TableName()).Where("id in ?", userIds).Find(&res).Error
	return res, err
}

func (u *User) SearchUserByEmail(email string) (*User, error) {
	conn := db.GetDBConn()
	res := &User{}
	err := conn.Table(u.TableName()).Where("email=?", email).First(res).Error
	return res, err
}
