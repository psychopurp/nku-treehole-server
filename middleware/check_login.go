package middleware

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/handler"
	"nku-treehole-server/pkg/jwt"
	"nku-treehole-server/pkg/logger"
	"nku-treehole-server/service"
	"strconv"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			logger.Errorf("未登陆 ")
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "未登陆",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
		var err error
		var userId int64
		claims, err := jwt.ParseToken(token)
		if err == nil {
			userId, err = strconv.ParseInt(claims.UserID, 10, 64)
		}
		if err != nil {
			logger.Errorf("token 解析失败")
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "登陆已过期，请重新登录",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
		// set userId to Context
		c.Set(config.UID, userId)

		userService := service.UserService{}
		if err := userService.CheckExpireAndRefresh(token); err != nil {
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "登陆已过期，请重新登录",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
	}
}
