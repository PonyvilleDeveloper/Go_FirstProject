package service

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Data     []byte
	Id       uint32
}

type httperror struct {
	Message string
}

func (c *Context) SendError(status int, mes string) {
	c.Response.WriteHeader(status)
	em := httperror{Message: mes}
	marsh, _ := json.Marshal(em)
	c.Response.Write(marsh)
}

func (c *Context) SendAnswer(data interface{}) {
	c.Response.Header().Set("Content-Type", "text/json")
	marsh, _ := json.Marshal(data)
	c.Response.Write(marsh)
}

func ExtractData[T any](ctx *Context) (ent T, err error) {
	err = json.NewDecoder(ctx.Request.Body).Decode(&ent)
	return
}
