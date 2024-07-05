package ginx

import (
	"bytes"
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// NewContext ...
func NewContext(ctx *gin.Context) *Context {
	return &Context{ctx}
}

// Handle ...
func Handle(f func(ctx *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(NewContext(c))
	}
}

// Context ...
type Context struct {
	*gin.Context
}

// ParamsString ...
func (c *Context) ParamsString() string {
	var paramsBuff bytes.Buffer
	for _, param := range c.Params {
		paramsBuff.WriteString(param.Key)
		paramsBuff.WriteString("=")
		paramsBuff.WriteString(param.Value)
		paramsBuff.WriteString(",")
	}
	paramsBuff.WriteString("consumes_time")
	paramsBuff.WriteString("=")
	paramsBuff.WriteString(strconv.FormatInt(c.processTime(), 10))
	paramsBuff.WriteString(",")

	if paramsBuff.String() != "" {
		return strings.TrimSuffix(paramsBuff.String(), ",")
	}
	return paramsBuff.String()
}

// ParamsMap ...
func (c *Context) ParamsMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["path"] = c.Request.URL
	resp["params"] = c.Params
	resp["consumes_time"] = c.processTime()
	return resp
}

// processTime ...
func (c *Context) processTime() (processTime int64) {
	startAt := c.GetInt64("start_at")
	now := time.Now().UnixMilli()
	if startAt > 0 {
		processTime = now - startAt
	}
	return
}

type ginKey struct{}

// NewGinContext returns a new Context that carries gin.Context value.
func NewGinContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, ginKey{}, c)
}

// FromGinContext returns the gin.Context value stored in ctx, if any.
func FromGinContext(ctx context.Context) (c *gin.Context, ok bool) {
	c, ok = ctx.Value(ginKey{}).(*gin.Context)
	return
}
