package engine

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type Error struct {
	Message string
}

func (c *Context) Error(status int, error string) {
	c.Response.WriteHeader(status)
	em := Error{Message: error}
	marsh, _ := json.Marshal(em)
	c.Response.Write(marsh)
}

func (c *Context) Print(data interface{}) {
	c.Response.Header().Set("Content-Type", "")
	marsh, _ := json.Marshal(data)
	c.Response.Write(marsh)
}

func ToStruct[T any](ctx *Context) (T, error) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var data T
	err := decoder.Decode(&data)
	return data, err
}
