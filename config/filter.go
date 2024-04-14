package config

import (
	"github.com/lyzzz123/illusionmvc/log"
	"net/http"
)

type FirstFilter struct {
}

func (firstFilter *FirstFilter) PreHandle(writer http.ResponseWriter, request *http.Request) error {
	log.Info("FirstFilter PreHandle:" + request.URL.Path)
	return nil
}

func (firstFilter *FirstFilter) PostHandle(writer http.ResponseWriter, request *http.Request) error {
	log.Info("FirstFilter PostHandle")

	return nil
}

func (firstFilter *FirstFilter) GetPriority() int {
	return 3
}

func (firstFilter *FirstFilter) GetPathPattern() string {
	return "/**"
}
