package main

import (
	"fmt"
	"net/http"
)

type handler interface {
	http.Handler
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

type HandlerBasedOnMap struct {
	// key 应该是method + url
	handlers map[string]func(ctx *Context)
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
