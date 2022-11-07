package httpctrl1

import "github.com/gin-gonic/gin"

func (c *Controller) Register(engine *gin.Engine) {
	r := engine.Group("api/v1")

	welcome := r.Group("/welcome")
	{
		welcome.GET("", c.HelloWorld.Welcome)
	}
}
