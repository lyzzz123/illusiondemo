package controller

import (
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/controller"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/response"
)

type TestController struct {
	TestServiceValue *service.TestService `require:"true"`
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

type PathValueParam struct {
	Hello string `paramValue:"hello_a"`
	Name  string `paramValue:"name_b"`
}

func (testController *TestController) PathValueTest(param *PathValueParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	log.Info("testController message:{}", message)
	testController.TestServiceValue.ShowPort()
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
		{
			Path:          "/path_value_test/{hello_a}/{name_b}",
			HttpMethod:    []string{httpmethod.GET},
			HandlerMethod: testController.PathValueTest,
		},
	}
}
