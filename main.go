package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panhdjf/scrum/controllers"
	"github.com/panhdjf/scrum/initializers"
	"github.com/panhdjf/scrum/routes"
)

var (
	server              *gin.Engine
	TaskController      controllers.TaskController
	TaskRouteController routes.TaskRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	TaskController = controllers.NewTaskController(initializers.DB)
	TaskRouteController = routes.NewRouteTaskController(TaskController)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	TaskRouteController.TaskRouter(router)
	log.Fatal(server.Run(":" + config.SeverPort))
}
