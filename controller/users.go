package controller

import (
	"net/http"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UpdateUser(ctx *gin.Context) {
	var input models.UserUpdateBody
	userId := ctx.Param("userId")
	ctx.BindJSON(&input)

	user, err := c.service.UpdateUser(userId, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
