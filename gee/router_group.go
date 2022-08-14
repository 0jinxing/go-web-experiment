package gee

type RouterGroup struct {
	app         *Gee
	router      *Router
	parent      *RouterGroup
	prefix      string
	middlewares []RouteHandler
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	app := group.app

	newGroup := &RouterGroup{
		app:         app,
		router:      group.router,
		parent:      group,
		prefix:      prefix,
		middlewares: make([]RouteHandler, 0),
	}
	
	return newGroup
}
