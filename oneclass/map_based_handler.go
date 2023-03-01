package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}
type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// key 应该是method + url
	handlers map[string]func(ctx *Context)
}

// Route 路由注册
func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	//  检测路由是否注册过
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("not any Route match"))
	}
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%S", method, path)
}

// 确保&HandlerBasedOnMap{} 一定实现了Handler接口  断定它是否某类型的常见写法
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 128),
	}
}
