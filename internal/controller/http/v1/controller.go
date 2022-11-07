package httpctrl1

import (
	service "plutoscrm/internal/service/v1"
	"time"
)

type Controller struct {
	HelloWorld *HelloWorldController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		HelloWorld: NewHelloWorld(service),
	}
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
	Time  time.Time   `json:"time"`
}

func NewResponse(v interface{}) (r *Response) {
	r = new(Response)
	r.Time = time.Now()

	if err, ok := v.(error); ok {
		r.Error = err
		return
	}

	r.Data = v
	return
}
