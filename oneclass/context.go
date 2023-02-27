package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) REadJson(req interface{}) error {
	r := c.R
	// 读取body，处理json
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

// 以下为扩展的便捷方法
func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequest(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
