package controller

import (
	"net/http"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	var input models.UserUpdateBody
	ctx.BindJSON(&input)

	user, err := c.service.UpdateUser(userId, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *Controller) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := c.service.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
