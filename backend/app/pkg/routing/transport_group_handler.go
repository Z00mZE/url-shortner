package routing

import "github.com/labstack/echo/v4"

type HTTPRouteHandler struct {
	endpoint   string
	method     string
	handler    echo.HandlerFunc
	middleware []echo.MiddlewareFunc
}

func NewHTTPRouteHandler(endpoint string, method string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *HTTPRouteHandler {
	return &HTTPRouteHandler{endpoint: endpoint, method: method, handler: handler, middleware: middleware}
}

func (t *HTTPRouteHandler) Endpoint() string {
	return t.endpoint
}

func (t *HTTPRouteHandler) Method() string {
	return t.method
}

func (t *HTTPRouteHandler) Handler() echo.HandlerFunc {
	return t.handler
}

func (t *HTTPRouteHandler) Middleware() []echo.MiddlewareFunc {
	return t.middleware
}
