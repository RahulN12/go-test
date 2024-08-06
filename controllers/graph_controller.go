package controllers

import (
	"fmt"
	"github.com/RahulN12/go-test/model"
	"github.com/RahulN12/go-test/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GraphController struct {
	service *service.GraphService
}

func NewGraphController(service *service.GraphService) *GraphController {
	return &GraphController{service: service}
}

func (c *GraphController) PostGraph(ctx *gin.Context) {
	var graph model.Graph
	if err := ctx.BindJSON(&graph); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := fmt.Sprintf("%d", c.service.GetGraphLength()+1)
	c.service.SaveGraph(id, &graph)

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (c *GraphController) GetGraph(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.service.GetGraph(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *GraphController) GetShortestPath(ctx *gin.Context) {
	id := ctx.Param("id")
	start := ctx.Query("start")
	end := ctx.Query("end")

	response, err := c.service.GetShortestPath(id, start, end)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *GraphController) DeleteGraph(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteGraph(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Graph deleted"})
}
