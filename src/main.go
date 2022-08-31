package main

import (
	"context"
	"fmt"
	"log"
	"textShare/router"
	"textShare/utils"

	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	ctx := context.Background()

	client := utils.InitClient()

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	app := gin.New()
	app.Use(gin.Recovery())

	baseGroup := app.Group("/api")
	{
		router.InitRoutes(baseGroup)
	}

	return app
}

func RunServer() {
	config := utils.Config()

	app := InitApp()
	//defer utils.CloseClient(10)

	log.Println("Starting server...")

	app.Run(config.ServerDsn)
}

func main() {
	RunServer()
}
