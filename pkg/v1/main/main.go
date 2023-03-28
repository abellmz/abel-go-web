package main

import (
	webv1 "abel-go-web/pkg/v1"
	"fmt"
	"net/http"
)

func main() {
	server := webv1.NewHandlerBasedOnTree("test-server")

	//server.Route("/", home)
	//server.Route("/body/once", user)
	//server.Route("/body/multi", createUser)
	//server.Route("/url/query", order)
	server.Route(http.MethodGet, "/user/sigup", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
func SignUp(ctx *webv1.Context) {
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
