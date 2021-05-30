package middleware

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/handler"
	"nku-treehole-server/pkg/logger"
	"nku-treehole-server/service"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			logger.Errorf("未登陆 ")
			c.JSON(200, handler.Response{
				Code: handler.LOGIN_EXPIRE,
				Msg:  "未登陆",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
		userService := service.UserService{}
		if err := userService.CheckExpireAndRefresh(token); err != nil {
			c.JSON(200, handler.Response{
				Code: handler.LOGIN_EXPIRE,
				Msg:  "登陆已过期，请重新登录",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
	}
}
