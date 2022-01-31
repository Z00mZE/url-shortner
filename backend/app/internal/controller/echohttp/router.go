package v1

import (
	"context"
	"fmt"
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
	SaveLink(ctx context.Context, url url.URL) error
	ViewLink(ctx context.Context, hash string)
	Redirect(ctx context.Context, url string) error
}

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
						minTTL := time.Now()
						maxTTL := minTTL.AddDate(1, 0, 0)
						return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{
							"ttlMin": minTTL.Format("2006-01-02"),
							"ttlMax": maxTTL.Format("2006-01-02"),
						})
					},
				),
				routing.NewHTTPRouteHandler(
					"/link",
					http.MethodPost,
					func(ctx echo.Context) error {
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
