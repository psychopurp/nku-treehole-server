package middleware

import "github.com/gin-gonic/gin"

// TODO: check login and refresh token
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("token")

	}
}
