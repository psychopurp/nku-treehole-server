package main

import (
	"github.com/gin-gonic/gin"
	"nku-treehole-server/config"
	"nku-treehole-server/db"
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

		r.POST("/api/post/create", empty)
		r.GET("/api/post/search", empty)
		r.POST("/api/post/comment", empty)

	}
}

func empty(ctx *gin.Context) {

}
