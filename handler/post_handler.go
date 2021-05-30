package handler

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/model"
)

type CreatePostRequest struct {
	Content string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var reqParam CreatePostRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "参数错误")
		return
	}
	var exist bool
	var uid int64
	userId, exist := c.Get(config.UID)
	if exist {
		uid, exist = userId.(int64)
	}
	if !exist {
		ErrorResponse(c, "发送失败，请重新登陆")
		return
	}

	post := &model.Post{}
	post.UserId = uid
	post.Content = reqParam.Content
	err := post.CreatePost(post)
	if err != nil {
		ErrorResponse(c, "发送失败，请重新登陆")
		return
	}
	SuccessResponse(c, map[string]interface{}{})
}
