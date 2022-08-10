package gee

import "net/http"

type Gee struct {
	route *Route
}

func New() *Gee {
	return &Gee{route: NewRoute()}
}

func (g *Gee) GET(pattern string, handler RouteHandler) {
	g.route.AddHandler("GET", pattern, handler)
}

func (g *Gee) POST(pattern string, handler RouteHandler) {
	g.route.AddHandler("POST", pattern, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	g.route.Handle(c)
}

func (g *Gee) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}
