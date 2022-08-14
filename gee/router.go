package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type RouteHandler func(ctx *Context)

type Router struct {
	handlers map[string]RouteHandler
	roots    map[string]*TrieNode
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]RouteHandler), roots: make(map[string]*TrieNode)}
}

func (r *Router) AddRoute(method string, pattern string, handler RouteHandler) {
	splits := SplitTrieNodePath(pattern)
	key := method + "\t" + pattern

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &TrieNode{}
	}
	r.roots[method].Append(pattern, splits, 0)
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method string, path string) (*TrieNode, map[string]string) {
	search := SplitTrieNodePath(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	node := root.Match(search, 0)

	if node != nil {
		splits := SplitTrieNodePath(node.leaf)
		for index, split := range splits {
			if split[0] == ':' {
				params[split[1:]] = search[index]
			}
			if split[0] == '*' && len(split) > 1 {
				params[split[1:]] = strings.Join(search[index:], "/")
				break
			}
		}
		return node, params
	}

	return nil, nil
}

func (r *Router) Handle(ctx *Context) {
	node, params := r.GetRoute(ctx.Req.Method, ctx.Req.URL.Path)

	if node == nil {
		msg := fmt.Sprintf("404 NOT FOUND: %s\n", ctx.Req.URL.Path)
		ctx.Status(http.StatusNotFound).Text(msg)
		return
	}
	ctx.Params = params
	key := ctx.Req.Method + "\t" + node.leaf
	r.handlers[key](ctx)
}
