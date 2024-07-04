package ginx

import (
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"net/http"
)

type Result struct {
	Code statusx.Status `json:"code"`
	Msg  string         `json:"message"`
	Data interface{}    `json:"data"`
}

type ListResponses struct {
	List interface{} `json:"list"`
	Page interface{} `json:"page,omitempty"`
}

type InfoResponses struct {
	Info interface{} `json:"info"`
}

// Render gin Success
func (c *Context) RenderSuccess(data interface{}) {
	c.Header("Cache-Control", "no-cache")

	_code := http.StatusOK
	result := Result{
		Code: statusx.StatusOk,
		Msg:  statusx.GetMsg(statusx.StatusOk),
		Data: data,
	}

	c.JSON(_code, result)
	return
}

// Render gin render
func (c *Context) Render(st statusx.Status, data interface{}, httpCode ...int) {
	c.Header("Cache-Control", "no-cache")

	_code := http.StatusOK
	if len(httpCode) > 0 {
		_code = httpCode[0]
	}

	result := Result{
		Code: st,
		Msg:  statusx.GetMsg(st),
		Data: data,
	}

	c.JSON(_code, result)
	return
}

// RenderWithMsg ...
func (c *Context) RenderWithMsg(st statusx.Status, data interface{}, msg string, httpCode ...int) {
	c.Header("Cache-Control", "no-cache")

	_code := http.StatusOK
	if len(httpCode) > 0 {
		_code = httpCode[0]
	}

	result := Result{
		Code: st,
		Msg:  msg,
		Data: data,
	}

	c.JSON(_code, result)
	return
}
