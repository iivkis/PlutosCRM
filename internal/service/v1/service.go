package service1

type Service struct {
	HelloWorld *HelloWorldServie
}

func NewService() *Service {
	return &Service{
		HelloWorld: NewHelloWorld(),
	}
}
