package gee

import (
	"encoding/json"
	"net/http"
)

type H map[string]any

type Context struct {
	W   http.ResponseWriter
	Req *http.Request

	Params map[string]string
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{W: w, Req: req}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// handle Response
func (ctx *Context) Status(code int) *Context {
	ctx.W.WriteHeader(code)
	return ctx
}

func (ctx *Context) Header(key string, value string) *Context {
	ctx.W.Header().Set(key, value)
	return ctx
}

func (ctx *Context) Byte(val []byte) {
	ctx.W.Write(val)
}

func (ctx *Context) Text(val string) {
	ctx.Header("Content-Type", "text/plain")
	ctx.Byte([]byte(val))
}

func (ctx *Context) HTML(val string) {
	ctx.Header("Content-Type", "text/html")
	ctx.Byte([]byte(val))
}

func (ctx *Context) JSON(val any) {
	ctx.Header("Content-Type", "application/json")

	if err := json.NewEncoder(ctx.W).Encode(val); err != nil {
		http.Error(ctx.W, err.Error(), 500)
	}
}
