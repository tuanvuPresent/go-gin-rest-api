package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-gin/api"
	"go-gin/config"
	"go-gin/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.Init()
	_, err = database.ConnectDB()
	app := gin.Default()
	api.ApplyRoutes(app)
	err = app.Run(":" + config.Config.APP_PORT)
	if err != nil {
		return
	}
}
