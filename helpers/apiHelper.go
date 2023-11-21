package helpers

import "github.com/beego/beego/v2/server/web"

type ResponseSuccess struct {
	ResStatus int         `json:"status"`
	Message   string      `json:"message"`
	Result    interface{} `json:"data"`
}

type ResponseFailed struct {
	ResStatus int    `json:"status"`
	Message   string `json:"message"`
}

func ApiSuccessResponse(c *web.Controller, result interface{}, message string) {
	if result == "" {
		result = map[int]int{}
	}
	response := ResponseSuccess{
		Message:   message,
		ResStatus: 1,
		Result:    result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func ApiFailedResponse(c *web.Controller, message string) {
	response := ResponseFailed{
		Message:   message,
		ResStatus: 0,
	}
	c.Data["json"] = response
	c.ServeJSON()
}
