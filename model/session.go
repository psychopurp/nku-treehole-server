package model

import "time"

type Session struct {
	ID        int       `gorm:"column:id" json:"id" form:"id"`
	UserId    int64     `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Token     string    `gorm:"column:token" json:"token" form:"token"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at" form:"expired_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}

func (c *Session) Session() string {
	return "sessions"
}
