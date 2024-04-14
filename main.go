package main

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusiondemo/config"
	"github.com/lyzzz123/illusiondemo/controller"
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusionmvc"
	"reflect"
)

func RegisterAll() {
	illusion.Register(&config.TimeConvert{})
	illusion.Register(&config.StringResponseWriter{ResponseType: reflect.TypeOf(*new(config.StringResponse))})
	illusion.Register(&config.XmlConverter{})
	illusion.Register(&config.FirstFilter{})
	illusion.Register(&config.ListenerTest{})

	illusion.Register(&service.TestService{})
	illusion.Register(&controller.TestController{})
	illusion.Register(&illusionmvc.Runner{})
}
func main() {
	RegisterAll()
	illusion.Start()
}
