package main

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/db"
	"nku-treehole-server/handler"
	"nku-treehole-server/middleware"
	"nku-treehole-server/pkg/logger"
	"os"
)

func main() {
	isDev := true
	if os.Getenv("DOCKER") == "true" {
		isDev = false
		logger.Infof("Server running on product env")
	}

	//必须先初始化配置文件
	config.Init("./conf", isDev)
	db.InitDB()
	r := gin.Default()
	api := r.Group("/api")
	{
		// 注册接口
		api.POST("/user/login", handler.Login)
		api.POST("/user/register", handler.Register)
		api.POST("/logout", middleware.CheckLogin(), empty)
		api.GET("/user/getUserInfo", middleware.CheckLogin(), empty)

		api.POST("/post/createPost", middleware.CheckLogin(), handler.CreatePost)
		api.GET("/post/getPosts", middleware.CheckLogin(), handler.GetPosts)
		api.POST("/post/comment", middleware.CheckLogin(), empty)
	}
	if !config.Conf.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	r.NoRoute(empty)
	logger.Fatalf("%v", r.Run(config.Conf.GetString("addr")))
}

func empty(ctx *gin.Context) {
	ctx.String(200, "nku-treehole-server")
}
