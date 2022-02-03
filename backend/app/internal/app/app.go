package app

import (
	"fmt"
	"github.com/Z00mZE/url-shortner/assets"
	"github.com/Z00mZE/url-shortner/config"
	echoRouting "github.com/Z00mZE/url-shortner/internal/controller/echohttp"
	"github.com/Z00mZE/url-shortner/internal/render/html/embedded"
	"github.com/Z00mZE/url-shortner/pkg/converter"
	"github.com/Z00mZE/url-shortner/pkg/httpserver"
	"github.com/Z00mZE/url-shortner/pkg/postgres"
	"github.com/Z00mZE/url-shortner/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.AppConfig) {
	entOrm, entOrmError := postgres.NewEntPostgres(
		cfg.Database.DSN,
		postgres.MaxPoolSize(10),
		postgres.SetConnMaxLifetime(time.Minute*15),
		postgres.SetConnMaxIdleTime(time.Minute*5),
	)
	if entOrmError != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewPostgresORM: %w", entOrmError))
	}
	defer func() {
		if err := entOrm.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := echo.New()
	httpServer := httpserver.NewHTTPServer(app, httpserver.Port(cfg.HTTP.Port))
	app.Use(
		middleware.Gzip(),
	)
	app.Renderer = embedded.NewRenderer()
	//	иницализация роутинга
	shorter := usecase.NewShortener(entOrm.ShortUrl, converter.NewDefaultDecimalConverter())
	go initRoutes(app, shorter)

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
		log.Panic(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		log.Panic(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func initRoutes(app *echo.Echo, shortner *usecase.Shortener) {
	assetHandler, assetHandlerError := assets.NewAssetFileServer()
	if assetHandlerError == nil {
		app.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))
	}

	if bindRouteError := echoRouting.NewRouter(app, shortner); bindRouteError != nil {
		app.Logger.Error(bindRouteError)
	}
}
