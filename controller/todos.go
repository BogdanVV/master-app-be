package controller

import (
	"net/http"
	"strconv"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateTodo(ctx *gin.Context) {
	var input models.TodoCreateBody
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := c.service.CreateTodo(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": todo})
}

func (c *Controller) GetAllTodos(ctx *gin.Context) {
	todos, err := c.service.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}

func (c *Controller) GetTodoById(ctx *gin.Context) {
	id, isIdParam := ctx.Params.Get("id")
	idInt, err := strconv.Atoi(id)
	if !isIdParam || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	todo, err := c.service.GetTodoById(idInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todo})
}

func (c *Controller) DeleteTodoById(ctx *gin.Context) {
	id, isIdParam := ctx.Params.Get("id")
	idInt, err := strconv.Atoi(id)
	if !isIdParam || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	err = c.service.DeleteTodoById(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) UpdateTodoById(ctx *gin.Context) {
	id, isIdParam := ctx.Params.Get("id")
	idInt, err := strconv.Atoi(id)
	if !isIdParam || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	var input models.TodoUpdateBody
	err = ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedTodo, err := c.service.UpdateTodoById(idInt, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedTodo})
}
