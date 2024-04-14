package controller

import (
	"fmt"
	"github.com/lyzzz123/illusiondemo/proto"
	"github.com/lyzzz123/illusiondemo/service"
	"github.com/lyzzz123/illusiondemo/utils"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/controller"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/response"
	"mime/multipart"
	"os"
)

type TestController struct {
	TestServiceValue *service.TestService `require:"true"`
}

type JSONParam struct {
	Hello     string `json:"hello"`
	Name      string `json:"name"`
	UserAgent string `paramValue:"User-Agent"`
}

// JsonRequestTest Content-Type must be application/json
func (testController *TestController) JsonRequestTest(param *JSONParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: &utils.Response{
			Code:    0,
			Message: "success",
			Data:    message,
		},
	}, nil
}

type PathValueParam struct {
	Hello     string `paramValue:"hello_a"`
	Name      string `paramValue:"name_b"`
	UserAgent string `paramValue:"User-Agent"`
}

func (testController *TestController) PathValueTest(param *PathValueParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	log.Info("testController message:{}", message)
	testController.TestServiceValue.ShowPort()
	return &response.JSONResponse{
		Data: &utils.Response{
			Code:    0,
			Message: "success",
			Data:    message,
		},
	}, nil
}

// ProtobufTest protobuf test
func (testController *TestController) ProtobufTest(student *proto.Student) (*response.ProtobufResponse, error) {
	fmt.Println(student)
	student.Age = student.Age + 10
	return &response.ProtobufResponse{Data: student}, nil
}

type FormDataParam struct {
	Hello    string                `paramValue:"hello"`
	Name     string                `paramValue:"name"`
	TestFile *multipart.FileHeader `paramValue:"testFile"`
}

// FormDataTest Content-Type must be multipart/form-data
func (testController *TestController) FormDataTest(param *FormDataParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name + " " + param.TestFile.Filename
	return &response.JSONResponse{
		Data: &utils.Response{
			Code:    0,
			Message: "success",
			Data:    message,
		},
	}, nil
}

type FormUrlEncodedParam struct {
	Hello string `paramValue:"hello"`
	Name  string `paramValue:"name"`
}

// FormUrlEncodedTest Content-Type must be application/x-www-form-urlencoded
func (testController *TestController) FormUrlEncodedTest(param *FormUrlEncodedParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	return &response.JSONResponse{
		Data: &utils.Response{
			Code:    0,
			Message: "success",
			Data:    message,
		},
	}, nil
}

func (testController *TestController) Download() (*response.FileResponse, error) {
	fr := &response.FileResponse{}
	fileInfo, _ := os.Stat("E:\\temp\\restart.sh")
	file, _ := os.Open("E:\\temp\\restart.sh")
	fr.Name = fileInfo.Name()
	fr.Size = fileInfo.Size()
	fr.File = file
	return fr, nil
}

type CookieParam struct {
	Cookie1 string `paramValue:"Cookie_1"`
}

func (testController *TestController) CookieParamTest(param *CookieParam) (*response.JSONResponse, error) {
	message := param.Cookie1
	return &response.JSONResponse{
		Data: &utils.Response{
			Code:    0,
			Message: "success",
			Data:    message,
		},
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
		{
			Path:          "/protobuf_test",
			HttpMethod:    []string{httpmethod.POST},
			HandlerMethod: testController.ProtobufTest,
		},
		{
			Path:          "/form_url_encoded_test",
			HttpMethod:    []string{httpmethod.GET},
			HandlerMethod: testController.FormUrlEncodedTest,
		},
		{
			Path:          "/form_data_test",
			HttpMethod:    []string{httpmethod.POST},
			HandlerMethod: testController.FormDataTest,
		},
		{
			Path:          "/download",
			HttpMethod:    []string{httpmethod.GET},
			HandlerMethod: testController.Download,
		},
		{
			Path:          "/cookie_param_test",
			HttpMethod:    []string{httpmethod.GET},
			HandlerMethod: testController.CookieParamTest,
		},
	}
}
