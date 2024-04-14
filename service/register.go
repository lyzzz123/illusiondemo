package service

import "github.com/lyzzz123/illusion"

func Register() {
	illusion.Register(&TestService{})
}
