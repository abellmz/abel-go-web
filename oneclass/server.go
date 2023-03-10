package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	// method Post,Get,Put,Delete
	//Route(method string, pattern string, handleFunc func(ctx *Context))
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

// Route 路由注册
func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})

	//key := s.handler.key(method,pattern)
	//s.handler.handlers[key] = handleFunc
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

// 是否支持路由级别 有些filter只允许在user路径下使用，有些允许它在全路径下触发
func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	// 定义root的
	var root Filter = handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	res := &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}

	return res
}

// func SignUp(w http.ResponseWriter, r *http.Request) {
func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		//fmt.Fprintf(w, "err: %v", err)
		//return
		ctx.BadRequest(err)
	}
	//body,err := io.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Fprintf(w,"read body failed: %v", err)
	//	return
	//}
	//// 将body中对应的数据放入req的字段中
	//err = json.Unmarshal(body, req)
	//if err != nil {
	//	fmt.Fprintf(w,"deserialized failed: %v", err)
	//	return
	//}

	resp := &commonResponse{
		Data: 123,
	}
	//respJson, err :=json.Marshal(resp)
	//if err != nil {
	//}
	//w.WriteHeader(http.StatusOK)
	//// 返回一个虚拟的user id，表示注册成功
	////fmt.Fprintf(w, "invalid request: %v", 123)
	//// byte切片和string可以互相转
	//fmt.Fprintf(w, string(respJson))

	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		// 日志
		fmt.Printf("写入响应失败： %v", err)
	}
}

type signUpReq struct {
	// `json:"email"` Tag 特性：可运行时通过反射拿到，反序列化的名称可以是`json:"email_xxx"  声明式的api
	Email             string `json:"email"`
	Password          string `json:password`
	ConfirmedPassword string `json:confirmed_password`
}

type commonResponse struct {
	BizCode int         `json:biz_code`
	Msg     string      `json:msg`
	Data    interface{} `json:data`
}
