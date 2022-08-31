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

 	// init client
	// logging that client is inited
	// defer close client

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
