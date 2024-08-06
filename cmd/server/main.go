package main

import (
	"github.com/RahulN12/go-test/config"
	"github.com/RahulN12/go-test/controllers"
	"github.com/RahulN12/go-test/repository"
	"github.com/RahulN12/go-test/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()

	r := gin.Default()

	repo := repository.NewGraphRepository()
	service := service.NewGraphService(repo)
	controller := controllers.NewGraphController(service)

	r.POST("/graph", controller.PostGraph)
	r.GET("/graph/:id/shortest-path", controller.GetGraph)
	r.GET("/graph/:id/shortest-path", controller.GetShortestPath)
	r.DELETE("/graph/:id", controller.DeleteGraph)

	r.Run(":8080")

}
