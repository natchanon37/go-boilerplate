package httpserver

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Context interface {
	Bind(interface{}) error
	JSON(code int, obj interface{})
	Redirect(code int, location string)
	GetQuery(string) string
	GetQueryInt(string) (int, error)
	GetParam(string) string
	GetParamInt(string) (int, error)
	AttachError(err error)
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	BindForm(interface{}) error
	SetHeader(string, string)
	GetHeader(string) string
	Data(httpCode int, contentType string, data []byte)
	GetRequestCtx() context.Context
}

type ginContext struct {
	*gin.Context
}

func (c *ginContext) GetRequestCtx() context.Context {
	return c.Request.Context()
}

func (c *ginContext) JSON(code int, obj interface{}) {
	c.Context.PureJSON(code, obj)
}

func (c *ginContext) Bind(obj interface{}) error {
	if err := c.Context.ShouldBindJSON(obj); err != nil {
		return &BindError{Message: err.Error()}
	}
	return nil
}

func (c *ginContext) GetQuery(key string) string {
	return c.Context.Query(key)
}

func (c *ginContext) Set(key string, value any) {
	c.Context.Set(key, value)
}

func (c *ginContext) Get(key string) (value any, exists bool) {
	return c.Context.Get(key)
}

func (c *ginContext) GetParamInt(key string) (int, error) {
	param := c.Context.Param(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

func (c *ginContext) GetQueryInt(key string) (int, error) {
	param := c.Context.Query(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

func (c *ginContext) GetParam(key string) string {
	return c.Context.Param(key)
}

func (c *ginContext) AttachError(err error) {
	c.Context.Error(err)
}

func (c *ginContext) SetHeader(key string, value string) {
	c.Header(key, value)
}

func (c *ginContext) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *ginContext) BindForm(v interface{}) error {
	return c.Context.ShouldBind(v)
}

func (c *ginContext) Data(httpCode int, contentType string, data []byte) {
	c.Context.Data(httpCode, contentType, data)
}

func (c *ginContext) Redirect(code int, location string) {
	c.Context.Redirect(code, location)
}

func NewContext(c *gin.Context) Context {
	return &ginContext{Context: c}
}
