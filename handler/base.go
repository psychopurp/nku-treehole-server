package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// 请求成功状态码
	SUCCESS_CODE int = 200

	// 登录过期，或者未登录
	LOGIN_EXPIRE int = 400
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS_CODE,
		Msg:  "",
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: map[string]interface{}{},
	})
}
