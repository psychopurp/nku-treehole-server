package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nku-treehole-server/config"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: config.SUCCESS_CODE,
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
