package gee

import "net/http"

type Gee struct {
	route *Router
}

func New() *Gee {
	return &Gee{route: NewRouter()}
}

func (g *Gee) GET(pattern string, handler RouteHandler) {
	g.route.AddRoute("GET", pattern, handler)
}

func (g *Gee) POST(pattern string, handler RouteHandler) {
	g.route.AddRoute("POST", pattern, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	g.route.Handle(c)
}

func (g *Gee) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}
