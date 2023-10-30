package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"nku-treehole-server/handler"
	"nku-treehole-server/middleware"
)

// V1 Router
func Setup() *gin.Engine {
	app := gin.New()

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	app.Use(gin.Recovery())
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	// Routes for api
	api := app.Group("/api")
	{

		api.POST("/user/login", handler.Login)
		api.POST("/user/register", handler.Register)
		api.POST("/logout", middleware.CheckLogin(), empty)
		api.GET("/user/getUserInfo", middleware.CheckLogin(), empty)
		api.POST("/post/createPost", middleware.CheckLogin(), handler.CreatePost)
		api.GET("/post/getPosts", middleware.CheckLogin(), handler.GetPosts)
		api.POST("/post/comment", middleware.CheckLogin(), empty)
	}

	return app
}

func empty(ctx *gin.Context) {
	ctx.String(200, "Not Implemented Yet")
}
