package main

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/db"
	"nku-treehole-server/pkg/logger"
)

func main() {
	//必须先初始化配置文件
	config.Init("./conf", true)
	db.InitDB()

	r := gin.Default()
	{
		// 注册接口
		r.POST("/api/login", empty)
		r.POST("/api/logout", empty)
		r.POST("/api/register", empty)

		r.POST("/api/post/send", empty)
		r.GET("/api/post/search", empty)
		r.POST("/api/post/comment", empty)
	}
	if !config.Conf.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.Fatalf("%v", r.Run(config.Conf.GetString("addr")))
}

func empty(ctx *gin.Context) {

}
