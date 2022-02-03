package v1

import (
	"context"
	"github.com/Z00mZE/url-shortner/ent/service"
	"github.com/Z00mZE/url-shortner/pkg/routing"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

type UseCaseError interface {
	StatusCode() int
	Message() string
}

type UseCase interface {
	SaveLink(context.Context, *url.URL) (*service.ShortUrl, error)
	FindLink(context.Context, string) (*service.ShortUrl, error)
}

const (
	formParamUserURL = "url"
	onSaveError      = "errorOnSave"
)

func NewRouter(app *echo.Echo, useCases UseCase) error {
	routing.BindHTTPRouteGroup(
		app,
		routing.NewHTTPRouteGroup(
			"",
			[]routing.IHttpHandler{
				routing.NewHTTPRouteHandler(
					"",
					http.MethodGet,
					func(ctx echo.Context) error {
						return ctx.Render(http.StatusOK, "index.html", nil)
					},
				),
				routing.NewHTTPRouteHandler(
					"/",
					http.MethodPost,
					func(ctx echo.Context) error {
						userRawURL := ctx.FormValue(formParamUserURL)
						if userRawURL == "" {
							return ctx.Render(http.StatusOK, "index.html", map[string]string{
								onSaveError: "url is empty",
							})
						}
						userURL, userParseError := url.ParseRequestURI(userRawURL)
						if userParseError != nil {
							return ctx.Render(http.StatusOK, "index.html", map[string]string{
								"userRawURL": userRawURL,
								onSaveError:  "url invalid",
							})
						}

						urlEntity, urlEntitySaveError := useCases.SaveLink(context.Background(), userURL)
						if urlEntitySaveError != nil {
							return ctx.Render(http.StatusOK, "index.html", map[string]string{
								"userRawURL": userRawURL,
								onSaveError:  urlEntitySaveError.Error(),
							})
						}
						return ctx.Redirect(http.StatusSeeOther, "/v"+urlEntity.HashID)
					}),
				routing.NewHTTPRouteHandler(
					"/v:hash",
					http.MethodGet,
					func(ctx echo.Context) error {
						hash := ctx.Param("hash")
						urlEntity, urlEntityFindError := useCases.FindLink(context.Background(), hash)
						if urlEntityFindError != nil {
							ctx.Logger().Error(urlEntityFindError)
							return ctx.Redirect(http.StatusSeeOther, "/")
						}
						//	@todo: нужно переделать формирование `shortUrl`
						return ctx.Render(
							http.StatusOK,
							"link.html",
							map[string]interface{}{
								"url":      urlEntity.URL,
								"shortUrl": "http://localhost:8080/r" + urlEntity.HashID,
								"ttl":      urlEntity.ExpiredAt.Format("2006-01-02"),
							},
						)
					},
				),
				routing.NewHTTPRouteHandler(
					"/r:hash",
					http.MethodGet,
					func(ctx echo.Context) error {
						hash := ctx.Param("hash")
						urlEntity, urlEntityFindError := useCases.FindLink(context.Background(), hash)
						if urlEntityFindError != nil {
							ctx.Logger().Error(urlEntityFindError)
							return ctx.Redirect(http.StatusSeeOther, "/")
						}

						return ctx.Redirect(http.StatusPermanentRedirect, urlEntity.URL)
					},
				),
			},
		),
	)
	return nil
}
