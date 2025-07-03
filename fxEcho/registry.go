package FxEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// RouteRegistryIf defines the interface for route registration
type RouteRegistryIf interface {
	Method() string
	Path() string
	Handle(ctx echo.Context) error
}

// GroupRegistryIf defines the interface for route group registration
type GroupRegistryIf interface {
	Prefix() string
	Register(g *echo.Group)
}

// MiddlewareRegistryIf defines the interface for middleware registration
type MiddlewareRegistryIf interface {
	Priority() int
	Middleware() echo.MiddlewareFunc
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(RouteRegistryIf)),
		fx.ResultTags(`group:"routes"`),
	)
}

// AsGroup annotates a group constructor for Fx.
func AsGroup(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(GroupRegistryIf)),
		fx.ResultTags(`group:"groups"`),
	)
}

// AsMiddleware annotates a middleware constructor for Fx.
func AsMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(MiddlewareRegistryIf)),
		fx.ResultTags(`group:"middlewares"`),
	)
}

// RouteBuilder provides a fluent interface for building routes
type RouteBuilder struct {
	method string
	path   string
	handle echo.HandlerFunc
}

// NewRoute creates a new route builder
func NewRoute(method, path string, handle echo.HandlerFunc) *RouteBuilder {
	return &RouteBuilder{
		method: method,
		path:   path,
		handle: handle,
	}
}

// GET creates a GET route
func GET(path string, handle echo.HandlerFunc) *RouteBuilder {
	return NewRoute("GET", path, handle)
}

// POST creates a POST route
func POST(path string, handle echo.HandlerFunc) *RouteBuilder {
	return NewRoute("POST", path, handle)
}

// PUT creates a PUT route
func PUT(path string, handle echo.HandlerFunc) *RouteBuilder {
	return NewRoute("PUT", path, handle)
}

// DELETE creates a DELETE route
func DELETE(path string, handle echo.HandlerFunc) *RouteBuilder {
	return NewRoute("DELETE", path, handle)
}

// PATCH creates a PATCH route
func PATCH(path string, handle echo.HandlerFunc) *RouteBuilder {
	return NewRoute("PATCH", path, handle)
}

// Build returns the route registry interface
func (rb *RouteBuilder) Build() RouteRegistryIf {
	return &routeRegistry{
		method: rb.method,
		path:   rb.path,
		handle: rb.handle,
	}
}

// routeRegistry implements RouteRegistryIf
type routeRegistry struct {
	method string
	path   string
	handle echo.HandlerFunc
}

func (r *routeRegistry) Method() string {
	return r.method
}

func (r *routeRegistry) Path() string {
	return r.path
}

func (r *routeRegistry) Handle(ctx echo.Context) error {
	return r.handle(ctx)
}

// GroupBuilder provides a fluent interface for building route groups
type GroupBuilder struct {
	prefix     string
	routes     []RouteRegistryIf
	children   []GroupRegistryIf
	middleware []echo.MiddlewareFunc
}

// NewGroup creates a new group builder
func NewGroup(prefix string) *GroupBuilder {
	return &GroupBuilder{
		prefix:     prefix,
		routes:     make([]RouteRegistryIf, 0),
		children:   make([]GroupRegistryIf, 0),
		middleware: make([]echo.MiddlewareFunc, 0),
	}
}

// AddRoute adds a route to the group
func (gb *GroupBuilder) AddRoute(route RouteRegistryIf) *GroupBuilder {
	gb.routes = append(gb.routes, route)
	return gb
}

// AddGroup adds a child group
func (gb *GroupBuilder) AddGroup(group GroupRegistryIf) *GroupBuilder {
	gb.children = append(gb.children, group)
	return gb
}

// Use adds middleware to the group
func (gb *GroupBuilder) Use(middleware ...echo.MiddlewareFunc) *GroupBuilder {
	gb.middleware = append(gb.middleware, middleware...)
	return gb
}

// Build returns the group registry interface
func (gb *GroupBuilder) Build() GroupRegistryIf {
	return &groupRegistry{
		prefix:     gb.prefix,
		routes:     gb.routes,
		children:   gb.children,
		middleware: gb.middleware,
	}
}

// groupRegistry implements GroupRegistryIf
type groupRegistry struct {
	prefix     string
	routes     []RouteRegistryIf
	children   []GroupRegistryIf
	middleware []echo.MiddlewareFunc
}

func (g *groupRegistry) Prefix() string {
	return g.prefix
}

func (g *groupRegistry) Register(group *echo.Group) {
	// Apply middleware to the group
	for _, m := range g.middleware {
		group.Use(m)
	}

	// Register routes in this group
	for _, route := range g.routes {
		group.Add(route.Method(), route.Path(), route.Handle)
	}

	// Register child groups
	for _, child := range g.children {
		childGroup := group.Group(child.Prefix())
		child.Register(childGroup)
	}
}
