package main

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusiondemo/config"
	"github.com/lyzzz123/illusiondemo/controller"
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusionmvc"
)

func main() {
	config.Register()
	controller.Register()
	service.Register()
	illusion.Register(&illusionmvc.Runner{})
	illusion.Start()
}
