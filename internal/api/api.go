package api

import (
	httpctrl1 "plutoscrm/internal/controller/http/v1"
	service1 "plutoscrm/internal/service/v1"

	"github.com/gin-gonic/gin"
)

func Launch() {
	engine := gin.Default()

	s1 := service1.NewService()
	hc1 := httpctrl1.NewController(s1)
	hc1.Register(engine)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
