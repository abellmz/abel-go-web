package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}
type Handler interface {
	http.Handler
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

func (h *HandlerBasedOnMap) ServeHttp(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	//  检测路由是否注册过
	if handler, ok := h.handlers[key]; ok {
		c := NewContext(writer, request)
		handler(c)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("not any Route match"))
	}
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%S", method, path)
}

// 确保&HandlerBasedOnMap{} 一定实现了Handler接口
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 128),
	}
}
