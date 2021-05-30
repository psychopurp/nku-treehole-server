package model

import (
	"nku-treehole-server/db"
	"time"
)

type Post struct {
	ID        int        `gorm:"column:id" json:"id" form:"id"`
	UserId    int64      `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Content   string     `gorm:"column:content" json:"content" form:"content"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (c *Post) TableName() string {
	return "posts"
}

func (c *Post) CreatePost(post *Post) error {
	conn := db.GetDBConn()
	err := conn.Table(c.TableName()).Create(post).Error
	return err
}
