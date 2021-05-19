package model

import "time"

type Comment struct {
	ID        int        `gorm:"column:id" json:"id" form:"id"`
	ReplyTo   int64      `gorm:"column:reply_to" json:"reply_to" form:"reply_to"`
	UserId    int64      `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Content   string     `gorm:"column:content" json:"content" form:"content"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (c *Comment) Comment() string {
	return "comments"
}
