package main

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusiondemo/controller"
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusionmvc"
)

func main() {
	illusion.Register(&service.TestService{})
	illusion.Register(&controller.TestController{})
	illusion.Register(&illusionmvc.Runner{})
	illusion.Start()
}
