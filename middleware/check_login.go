package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"nku-treehole-server/config"
	"nku-treehole-server/handler"
	"nku-treehole-server/pkg/jwt"
	"nku-treehole-server/pkg/logger"
	"nku-treehole-server/service"
)

func CheckLogin() gin.HandlerFunc {
	jwtHelper := jwt.GetJWTCrypto()

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			logger.Errorf("Not logged in")
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "Not logged in",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
		var err error
		var userId int64

		userID, err := jwtHelper.ValidateToken(token)
		if err == nil {
			userId, err = strconv.ParseInt(userID, 10, 64)
		}
		if err != nil {
			logger.Errorf("Token parsing failed")
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "Login has expired, please log in again",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
		// Set userId to Context
		c.Set(config.UID, userId)

		userService := service.UserService{}
		if err := userService.CheckExpireAndRefresh(token); err != nil {
			c.JSON(200, handler.Response{
				Code: config.LOGIN_EXPIRE,
				Msg:  "Login has expired, please log in again",
				Data: map[string]interface{}{},
			})
			c.Abort()
			return
		}
	}
}
