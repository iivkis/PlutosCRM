package service1

import "context"

type HelloWorldServie struct{}

func NewHelloWorld() *HelloWorldServie {
	return &HelloWorldServie{}
}

func (s *HelloWorldServie) Welcome(ctx context.Context) string {
	return "Hello world"
}
