package main

import (
	"context"
	"fmt"
	"net/http"

	"os"
	"os/signal"

	"github.com/datvvan/sample1/api"
	appMiddleware "github.com/datvvan/sample1/middlewares"

	"github.com/datvvan/sample1/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

const (
	defaultConfigPath = "./.env"
)

var (
	engine *echo.Echo
	ctx    context.Context
	cancel context.CancelFunc
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())

	//init config
	config.Init(defaultConfigPath)
	appMiddleware.InitLogger()
	validator := validator.New()

	//connect db
	// _, err := db.New()
	// if err != nil {
	// 	log.Fatalln("Connect db error: ", err.Error())
	// }

	engine = echo.New()
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	engine.Use(appMiddleware.Logging)
	engine.Use(middleware.Recover())
	engine.Validator = &appMiddleware.CustomValidator{Validator: validator}
}

func setupGracefulShutdown(ctx context.Context, port string, engine *echo.Echo) {
	signalForExit := make(chan os.Signal, 1)
	signal.Notify(signalForExit, os.Interrupt)

	go func() {
		if err := engine.Start(fmt.Sprintf(":%v", port)); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Application failed", err)
		}
	}()
	log.WithFields(log.Fields{"bind": port}).Info("Running application")

	stop := <-signalForExit
	log.Info("Stop signal Received", stop)
	if err := engine.Shutdown(ctx); err != nil {
		log.Info("engine.Shutdown err", err)
	}
}

func main() {
	api.RegisterAPI(engine)
	setupGracefulShutdown(ctx, config.Default.PORT, engine)
	cancel()
}
