package routing

import "github.com/labstack/echo/v4"

type IHttpRouteGroup interface {
	Prefix() string
	Middleware() []echo.MiddlewareFunc
	Handlers() []IHttpHandler
}
type IHttpHandler interface {
	Endpoint() string
	Method() string
	Handler() echo.HandlerFunc
	Middleware() []echo.MiddlewareFunc
}
