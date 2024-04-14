package config

import (
	"github.com/lyzzz123/illusion"
	"reflect"
)

func Register() {
	illusion.Register(&TimeConvert{})
	illusion.Register(&StringResponseWriter{ResponseType: reflect.TypeOf(*new(StringResponse))})
	illusion.Register(&XmlConverter{})
	illusion.Register(&FirstFilter{})
	illusion.Register(&ListenerTest{})

}
