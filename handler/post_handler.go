package handler

import (
	"github.com/gin-gonic/gin"

	"nku-treehole-server/config"
	"nku-treehole-server/dto"
	"nku-treehole-server/model"
	"nku-treehole-server/service"
)

func CreatePost(c *gin.Context) {
	var reqParam dto.CreatePostRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "Parameter error")
		return
	}
	var exist bool
	var uid int64
	userId, exist := c.Get(config.UID)
	if exist {
		uid, exist = userId.(int64)
	}
	if !exist {
		ErrorResponse(c, "Failed to send, please log in again")
		return
	}

	post := &model.Post{}
	post.UserId = uid
	post.Content = reqParam.Content
	err := post.CreatePost(post)
	if err != nil {
		ErrorResponse(c, "Failed to send, please log in again")
		return
	}
	SuccessResponse(c, map[string]interface{}{})
}

func GetPosts(c *gin.Context) {
	var reqParam dto.PageQuery
	if err := c.ShouldBindQuery(&reqParam); err != nil || (reqParam.Limit == 0 && reqParam.Page == 0) {
		ErrorResponse(c, "Parameter error")
		return
	}
	postService := &service.PostService{}
	res, err := postService.GetPosts(reqParam.Page, reqParam.Limit)
	if err != nil {
		ErrorResponse(c, "Query failed")
		return
	}
	SuccessResponse(c, res)
}
