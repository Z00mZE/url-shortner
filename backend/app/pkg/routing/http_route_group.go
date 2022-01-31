package routing

import "github.com/labstack/echo/v4"

type HttpRouteGroup struct {
	prefix     string
	handlers   []IHttpHandler
	middleware []echo.MiddlewareFunc
}

func NewHTTPRouteGroup(prefix string, handlers []IHttpHandler, middleware ...echo.MiddlewareFunc) *HttpRouteGroup {
	return &HttpRouteGroup{prefix: prefix, handlers: handlers, middleware: middleware}
}

func (t *HttpRouteGroup) Prefix() string {
	return t.prefix
}

func (t *HttpRouteGroup) Middleware() []echo.MiddlewareFunc {
	return t.middleware
}

func (t *HttpRouteGroup) Handlers() []IHttpHandler {
	return t.handlers
}
