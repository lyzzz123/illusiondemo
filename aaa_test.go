package main

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusiondemo/controller"
	"github.com/lyzzz123/illusiondemo/service"
	"testing"
)

type GetObjects struct {
	TS *service.TestService       `require:"true"`
	TC *controller.TestController `require:"true"`
}

func TestCookieParam(t *testing.T) {
	RegisterAll()
	getO := &GetObjects{}
	illusion.Register(getO)
	illusion.TestStart()
	fmt.Println()
}