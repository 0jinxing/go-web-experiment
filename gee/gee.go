package gee

import (
	"log"
	"net/http"
)

type (
	RouterGroup struct {
		prefix      string
		middlewares []RouteHandler
		app         *Gee
	}

	Gee struct {
		*RouterGroup
		router *Router
		groups []*RouterGroup
	}
)

func New() *Gee {
	app := &Gee{router: NewRouter()}
	app.RouterGroup = &RouterGroup{app: app}
	app.groups = make([]*RouterGroup, 0)

	return app
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	app := group.app

	newGroup := &RouterGroup{
		prefix: app.prefix + prefix,
		app:    app,
	}

	app.groups = append(app.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) AddRoute(method string, component string, handler RouteHandler) {
	pattern := group.prefix + component

	log.Printf("Route %4s - %s", method, pattern)
	group.app.router.AddRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler RouteHandler) {
	group.AddRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler RouteHandler) {
	group.AddRoute("POST", pattern, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	g.router.Handle(c)
}

func (g *Gee) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}
