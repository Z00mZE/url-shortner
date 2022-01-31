package routing

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type EndpointAdd interface {
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

func BindHTTPRouteGroup(app *echo.Echo, groups ...IHttpRouteGroup) []error {
	var errors []error
	for _, group := range groups {
		endpointGroup := app.Group(group.Prefix(), group.Middleware()...)
		{
			if bindError := BindHTTPRouteEndpoint(endpointGroup, group.Handlers()...); bindError != nil {
				errors = append(errors, bindError)
			}
		}
	}
	return errors
}
func BindHTTPRouteEndpoint(app EndpointAdd, handlers ...IHttpHandler) error {
	if len(handlers) == 0 {
		return fmt.Errorf("handler not implemented")
	}
	for _, handler := range handlers {
		app.Add(
			handler.Method(),
			handler.Endpoint(),
			handler.Handler(),
			handler.Middleware()...,
		)
	}

	return nil
}
