package model

import (
	"nku-treehole-server/db"
	"time"
)

type Session struct {
	ID        int        `gorm:"column:id" json:"id" form:"id"`
	UserId    int64      `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Token     string     `gorm:"column:token" json:"token" form:"token"`
	ExpiredAt time.Time  `gorm:"column:expired_at" json:"expired_at" form:"expired_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) CreateSession(userId int64, token string, expireTime time.Time) error {
	conn := db.GetDB()
	obj := &Session{UserId: userId, Token: token, ExpiredAt: expireTime}
	err := conn.Table(s.TableName()).Create(obj).Error
	return err
}

func (s *Session) GetSessionByUid(userId int64) (*Session, error) {
	conn := db.GetDB()
	res := &Session{}
	err := conn.Table(s.TableName()).Where("user_id=? and deleted_at is null", userId).First(res).Error
	return res, err
}

func (s *Session) DeleteOldSession(userId int64) error {
	conn := db.GetDB()
	err := conn.Table(s.TableName()).Where("user_id=? and deleted_at is null", userId).Update("deleted_at", time.Now()).Error
	return err
}

func (s *Session) GetSessionByToken(token string) (*Session, error) {
	conn := db.GetDB()
	res := &Session{}
	err := conn.Table(s.TableName()).Where("token=? and deleted_at is null", token).First(res).Error
	return res, err
}

func (s *Session) Refresh(token string, expireTime time.Time) error {
	conn := db.GetDB()
	err := conn.Table(s.TableName()).Where("token=? and deleted_at is null", token).Update("expired_at", expireTime).Error
	return err
}
