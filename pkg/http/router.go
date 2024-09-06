package http

import (
	"github.com/gin-gonic/gin"
)

type defaultRouter interface {
	GET(relativePath string, handlers ...interface{})
	POST(relativePath string, handlers ...interface{})
	PUT(relativePath string, handlers ...interface{})
	PATCH(relativePath string, handlers ...interface{})
	DELETE(relativePath string, handlers ...interface{})
	Use(middleware ...gin.HandlerFunc)
}

type Router interface {
	defaultRouter
	Group(relativePath string, handlers ...interface{}) RouterGroup
}

type RouterGroup interface {
	defaultRouter
	Group(relativePath string, handlers ...interface{}) RouterGroup
}

type ginRouter struct {
	engine *gin.Engine
}

type ginRouterGroup struct {
	engine *gin.RouterGroup
}

// implements httpserver.Router.
func (r *ginRouter) GET(relativePath string, handlers ...interface{}) {
	r.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

// implements httpserver.Router.
func (r *ginRouter) POST(relativePath string, handlers ...interface{}) {
	r.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

// implements httpserver.Router.
func (r *ginRouter) PUT(relativePath string, handlers ...interface{}) {
	r.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}

// PATCH implements httpserver.Router.
func (rg *ginRouter) PATCH(relativePath string, handlers ...interface{}) {
	rg.engine.PATCH(relativePath, convertHandlersToGin(handlers)...)
}

// implements httpserver.Router.
func (r *ginRouter) DELETE(relativePath string, handlers ...interface{}) {
	r.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

// implements httpserver.Router.
func (r *ginRouter) Group(relativePath string, handlers ...interface{}) RouterGroup {
	group := r.engine.Group(relativePath, convertHandlersToGin(handlers)...)
	return &ginRouterGroup{engine: group}
}

func (r *ginRouter) Use(middleware ...gin.HandlerFunc) {
	r.engine.Use(middleware...)
}

// GET implements httpserver.RouterGroup.
func (rg *ginRouterGroup) GET(relativePath string, handlers ...interface{}) {
	rg.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

// DELETE implements httpserver.RouterGroup.
func (rg *ginRouterGroup) DELETE(relativePath string, handlers ...interface{}) {
	rg.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

// POST implements httpserver.RouterGroup.
func (rg *ginRouterGroup) POST(relativePath string, handlers ...interface{}) {
	rg.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

// PUT implements httpserver.RouterGroup.
func (rg *ginRouterGroup) PUT(relativePath string, handlers ...interface{}) {
	rg.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}

// PATCH implements httpserver.RouterGroup.
func (rg *ginRouterGroup) PATCH(relativePath string, handlers ...interface{}) {
	rg.engine.PATCH(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) Use(middleware ...gin.HandlerFunc) {
	rg.engine.Use(middleware...)
}

func (rg *ginRouterGroup) Group(relativePath string, handlers ...interface{}) RouterGroup {
	group := rg.engine.Group(relativePath, convertHandlersToGin(handlers)...)
	return &ginRouterGroup{engine: group}
}

func NewRouter(isDebug bool, middlewares ...gin.HandlerFunc) *ginRouter {
	var r *gin.Engine

	if isDebug {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	// Add the custom middlewares
	r.Use(middlewares...)

	return &ginRouter{engine: r}
}

func convertHandlersToGin(handlers []interface{}) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc

	for _, h := range handlers {

		if handlerFunc, ok := h.(func(c *gin.Context)); ok {
			ginHandlers = append(ginHandlers, gin.HandlerFunc(handlerFunc))
		} else if ginHandler, ok := h.(gin.HandlerFunc); ok {
			ginHandlers = append(ginHandlers, ginHandler)
		} else if handlerFunc, ok := h.(func(c Context)); ok {
			ginHandlers = append(ginHandlers, convertToGinHandler(handlerFunc))
		} else {
			panic("unimplemented")
		}
	}

	return ginHandlers
}

func convertToGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&ginContext{Context: c})
	}
}
