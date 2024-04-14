package service

import "github.com/lyzzz123/illusionmvc/log"

type TestService struct {
	Port string `property:"server.port, true"`
}

func (testService *TestService) ShowPort() {
	log.Info("port:{}", testService.Port)
}
