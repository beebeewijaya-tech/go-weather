package main

import (
	"fmt"
	"go-weather/internal/delivery/http"
	"go-weather/internal/repositories/cache"
	weatherSvc "go-weather/internal/services/weather"
	weatherUsecase "go-weather/internal/usecases/weather"
	"go-weather/internal/utilities/app"
	"go-weather/internal/utilities/config"
	"go-weather/internal/utilities/logger"
	"log"
)

func Run() error {
	c := config.New()
	log := logger.New(c)

	app := app.New(
		c,
		log,
	)

	// register dependencies of storage
	cacheStorage := cache.New(app)

	// register services
	weatherSvc := weatherSvc.New(app, cacheStorage)

	// register usecase
	weatherUsecase := weatherUsecase.New(app, weatherSvc)

	// register routers
	r := http.New(weatherUsecase)

	address := c.GetInt64("web.port")
	if err := r.Start(fmt.Sprintf(":%d", address)); err != nil {
		return fmt.Errorf("error when starting address %v", err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
