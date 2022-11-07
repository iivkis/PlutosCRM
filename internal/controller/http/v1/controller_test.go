package httpctrl1

import (
	service "plutoscrm/internal/service/v1"

	"github.com/gin-gonic/gin"
)

func engine() *gin.Engine {
	e := gin.Default()

	ctrl := NewController(service.NewService())
	ctrl.Register(e)

	return e
}
