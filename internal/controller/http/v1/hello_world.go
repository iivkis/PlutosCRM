package httpctrl1

import (
	"net/http"
	servicev1 "plutoscrm/internal/service/v1"

	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	service *servicev1.Service
}

func NewHelloWorld(service *servicev1.Service) *HelloWorldController {
	return &HelloWorldController{
		service: service,
	}
}

type HelloWorldWelcomeResponse struct {
	Text string `json:"text"`
}

func (c *HelloWorldController) Welcome(ctx *gin.Context) {
	resp := HelloWorldWelcomeResponse{
		Text: c.service.HelloWorld.Welcome(ctx),
	}

	ctx.JSON(http.StatusOK, NewResponse(&resp))
}
