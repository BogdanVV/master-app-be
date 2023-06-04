package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Signup(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	newUserId, err := c.service.Signup(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": gin.H{"id": newUserId}})
}

func (c *Controller) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invlaid body"})
		return
	}

	user, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("Authorization", user.AccessToken, 3600*24*30, "/", ctx.Request.URL.Hostname(), false, true)
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	token, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	newToken, err := c.service.RefreshAccessTokenToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("Authorization", newToken, 3600*24*30, "/", ctx.Request.URL.Hostname(), false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "the token was refreshed successfully"})
}
