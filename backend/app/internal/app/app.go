package app

import (
	"fmt"
	"github.com/Z00mZE/url-shortner/config"
	echoRouting "github.com/Z00mZE/url-shortner/internal/controller/echohttp"
	"github.com/Z00mZE/url-shortner/pkg/httpserver"
	"github.com/Z00mZE/url-shortner/pkg/render"
	"github.com/Z00mZE/url-shortner/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.AppConfig) {
	app := echo.New()
	httpServer := httpserver.NewHTTPServer(app, httpserver.Port(cfg.HTTP.Port))
	app.Use(
		middleware.Gzip(),
	)
	app.Renderer = render.NewTemplateRenderer("template/*.html")
	//	иницализация роутинга
	shortner := &usecase.Shortner{}
	go initRoutes(app, shortner)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
		if shutdownError := httpServer.Shutdown(); shutdownError != nil {
			log.Panicln(shutdownError)
		}
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func initRoutes(app *echo.Echo, shortner *usecase.Shortner) {
	app.Static("/assets", "assets")
	if bindRouteError := echoRouting.NewRouter(app, shortner); bindRouteError != nil {
		app.Logger.Error(bindRouteError)
	}
}