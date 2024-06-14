package weather

import (
	"go-weather/internal/entities"
	"go-weather/internal/services/weather"
	"go-weather/internal/utilities/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IWeatherUsecase interface {
	GetCurrentWeather(ctx echo.Context) error
}

type WeatherUsecase struct {
	weatherSvc weather.IWeatherService
	appConfig  *app.AppConfig
}

func New(
	appConfig *app.AppConfig,
	weatherSvc weather.IWeatherService,
) *WeatherUsecase {
	return &WeatherUsecase{
		appConfig:  appConfig,
		weatherSvc: weatherSvc,
	}
}

func (u *WeatherUsecase) GetCurrentWeather(ctx echo.Context) error {
	var weather entities.Weather
	if err := ctx.Bind(&weather); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	weatherRequest := entities.WeatherRequest{
		Weather: entities.Weather{
			Lat: weather.Lat,
			Lon: weather.Lon,
		},
		AppID: u.appConfig.Config.GetString("weather.apiKey"),
	}
	res, err := u.weatherSvc.GetCurrentWeather(ctx.Request().Context(), weatherRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
