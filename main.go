package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"nku-treehole-server/config"
	"nku-treehole-server/db"
	v1 "nku-treehole-server/router/v1"
)

func main() {
	Run(".env.example")
}

// Set configuration
// Change this func to "exported"  to make Test package can access it
func SetConfiguration(configPath string) {
	// Setup config from path
	// Default is .env in root folder
	config.Setup(configPath)
	// Calling setup db
	db.SetupDB()
	// Calling cloudinary storage
	// config.InitializeCloudinary()
	gin.SetMode(config.GetConfig().Server.Mode)
}

// Run the new API with designated configuration
func Run(configPath string) {
	if configPath == "" {
		configPath = ".env"
	}
	SetConfiguration(configPath)
	conf := config.GetConfig()

	// Routing
	web := v1.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
