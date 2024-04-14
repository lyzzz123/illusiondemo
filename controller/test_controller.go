package controller

import (
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/controller"
	"github.com/lyzzz123/illusionmvc/response"
)

type TestController struct {
}

type JSONParam struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

// JsonRequestTest Content-Type must be application/json
func (testController *TestController) JsonRequestTest(param *JSONParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func (testController *TestController) Export() []*controller.Exporter {
	return []*controller.Exporter{
		{
			Path:          "/json_request_test",
			HttpMethod:    []string{httpmethod.POST},
			HandlerMethod: testController.JsonRequestTest,
		},
	}
}
