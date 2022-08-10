package gee

import (
	"fmt"
	"log"
	"net/http"
)

type RouteHandler func(ctx *Context)

type Route struct {
	handlers map[string]RouteHandler
}

func NewRoute() *Route {
	return &Route{handlers: make(map[string]RouteHandler)}
}

func (r *Route) AddHandler(method string, pattern string, handler RouteHandler) {
	log.Printf("Route %s\t%s", method, pattern)

	key := method + "\t" + pattern
	r.handlers[key] = handler
}

func (r *Route) Handle(ctx *Context) {
	key := ctx.Req.Method + "\t" + ctx.Req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
		return
	}
	msg := fmt.Sprintf("404 NOT FOUND: %s\n", ctx.Req.URL.Path)
	ctx.Status(http.StatusNotFound).Text(msg)
}
