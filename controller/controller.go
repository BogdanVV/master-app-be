package controller

import (
	"net/http"

	"github.com/bogdanvv/master-app-be/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
