package v1

import (
	"context"
	"fmt"
	"github.com/Z00mZE/url-shortner/ent"
	"github.com/Z00mZE/url-shortner/pkg/routing"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"time"
)

type UseCaseError interface {
	StatusCode() int
	Message() string
}

type UseCase interface {
	SaveLink(context.Context, *url.URL) (ent.Urls, error)
	ViewLink(context.Context, string) (ent.Urls, error)
	Redirect(context.Context, string) (ent.Urls, error)
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
								onSaveError: urlEntitySaveError.Error(),
							})
						}
						fmt.Println("urlEntity:", urlEntity)
						//useCases.SaveLink()
						//useCaseContext := context.Background()
						//useCases.Redirect(useCaseContext, "asd")
						return ctx.Redirect(http.StatusSeeOther, "/link/asDAW123")
					}),
				routing.NewHTTPRouteHandler(
					"/link/:hash",
					http.MethodGet,
					func(ctx echo.Context) error {
						var hash string
						if err := ctx.Bind(&hash); err != nil {
							fmt.Println("error", err)
							return ctx.Redirect(http.StatusTemporaryRedirect, "/")
						}
						//useCaseContext := context.Background()
						//useCases.ViewLink(useCaseContext, hash)
						return ctx.Render(
							http.StatusOK,
							"link.html",
							map[string]interface{}{
								"url":      "https://ya.ru",
								"shortUrl": "https://asd.ru/" + hash,
								"ttl":      time.Now().Format("2006-01-02"),
							},
						)
					},
				),
				routing.NewHTTPRouteHandler(
					"/short/:link",
					http.MethodGet,
					func(ctx echo.Context) error {
						var hash string
						if err := ctx.Bind(&hash); err != nil {
							fmt.Println("error", err)
							return ctx.Redirect(http.StatusTemporaryRedirect, "/")
						}

						//useCaseContext := context.Background()
						//useCases.Redirect(useCaseContext, hash)

						return ctx.Redirect(http.StatusPermanentRedirect, "http://yandex.ru")
					},
				),
			},
		),
	)
	return nil
}
