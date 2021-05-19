package model

import (
	"nku-treehole-server/db"
	"time"
)

type Session struct {
	ID        int       `gorm:"column:id" json:"id" form:"id"`
	UserId    int64     `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Token     string    `gorm:"column:token" json:"token" form:"token"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at" form:"expired_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) CreateSession(userId int64, token string, expiredTime time.Time) error {
	conn := db.GetDBConn()
	obj := &Session{UserId: userId, Token: token, ExpiredAt: expiredTime}
	err := conn.Table(s.TableName()).Create(obj).Error
	return err
}
