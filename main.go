package main

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/db"
	"nku-treehole-server/handler"
	"nku-treehole-server/middleware"
	"nku-treehole-server/pkg/logger"
)

func main() {
	//必须先初始化配置文件
	config.Init("./conf", true)
	db.InitDB()
	r := gin.Default()
	api := r.Group("/api")
	{
		// 注册接口
		api.POST("/user/login", handler.Login)
		api.POST("/user/register", handler.Register)
		api.POST("/logout", middleware.CheckLogin(), empty)

		api.POST("/post/createPost", middleware.CheckLogin(), empty)
		api.GET("/post/getPosts", handler.Login)
		api.POST("/post/comment", middleware.CheckLogin(), empty)
	}
	if !config.Conf.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Fatalf("%v", r.Run(config.Conf.GetString("addr")))
}

func empty(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"name ": "jack"})
}
