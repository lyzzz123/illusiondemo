package test

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusiondemo/controller"
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusionmvc/log"
	"testing"
)

type GetObjects struct {
	TS *service.TestService       `require:"true"`
	TC *controller.TestController `require:"true"`
}

func TestCookieParam(t *testing.T) {
	Register()

	getO := &GetObjects{}
	illusion.Register(getO)
	illusion.TestStart()

	fr, _ := getO.TC.Download()

	log.Info(fr.Name)
}
