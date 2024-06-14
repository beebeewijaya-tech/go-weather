package http

import (
	"go-weather/internal/usecases/weather"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	echo           *echo.Echo
	weatherUsecase weather.IWeatherUsecase
}

func New(
	weatherUsecase weather.IWeatherUsecase,
) *Router {
	r := &Router{
		weatherUsecase: weatherUsecase,
	}
	e := echo.New()

	e.Use(middleware.Logger())

	r.echo = e
	r.initRoutes()

	return r
}

func (r *Router) initRoutes() {
	router := r.echo.Group("/api/v1")
	router.POST("/weather/search", r.weatherUsecase.GetCurrentWeather)
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
