package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nku-treehole-server/handler"
)

// No Method Handler global middleware
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, handler.Response{
			Msg: "method not permitted",
		})
	}
}

// No Route Handler global middleware
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusNotFound, handler.Response{
			Msg: "the processing function of the request route was not found",
		})
	}
}
