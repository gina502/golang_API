package main

import (
	"golangAPI/database"
	"golangAPI/src"

	_ "golangAPI/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// schemes http
// @BasePath /v1/users

func main() {
	router := gin.New()

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)
	go func() {
		database.DD()
		database.RR()
	}()
	router.Run(":8000")
}
