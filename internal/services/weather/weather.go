package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"go-weather/internal/entities"
	"go-weather/internal/repositories/cache"
	"go-weather/internal/utilities/app"
	"net/http"
	"time"
)

type IWeatherService interface {
	GetCurrentWeather(ctx context.Context, weather entities.WeatherRequest) (entities.WeatherResponse, error)
	GetCurrentWeatherCache(ctx context.Context, key string) (entities.WeatherResponse, error)
}

type WeatherService struct {
	appConfig    *app.AppConfig
	cacheStorage cache.ICacheStorage
}

func New(appConfig *app.AppConfig, cacheStorage cache.ICacheStorage) *WeatherService {
	return &WeatherService{
		appConfig:    appConfig,
		cacheStorage: cacheStorage,
	}
}

func (w *WeatherService) GetCurrentWeatherCache(ctx context.Context, key string) (entities.WeatherResponse, error) {
	var weatherResponse entities.WeatherResponse
	if key == "" {
		return weatherResponse, fmt.Errorf("key is required")
	}

	res, err := w.cacheStorage.Get(ctx, key)
	if err != nil {
		return weatherResponse, err
	}

	weatherString, ok := res.Value.(string)
	if !ok {
		weatherString = ""
	}

	err = json.Unmarshal([]byte(weatherString), &weatherResponse)
	if err != nil {
		w.appConfig.Logger.Errorf("error when unmarshal cache %v", err.Error())
		return weatherResponse, err
	}

	return weatherResponse, nil
}

func (w *WeatherService) GetCurrentWeather(ctx context.Context, weather entities.WeatherRequest) (entities.WeatherResponse, error) {
	var weatherResponse entities.WeatherResponse
	cacheKey := fmt.Sprintf("%f,%f", weather.Lat, weather.Lon)

	val, err := w.GetCurrentWeatherCache(ctx, cacheKey)
	if err == nil {
		return val, nil
	}

	host := w.appConfig.Config.GetString("weather.host")
	urlPost := fmt.Sprintf("%s/weather?lat=%f&lon=%f&appid=%s", host, weather.Lat, weather.Lon, weather.AppID)

	resp, err := http.Get(urlPost)
	if err != nil {
		return weatherResponse, fmt.Errorf("error when calling the weather app %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return weatherResponse, fmt.Errorf("error when parsing json from weather response %v", err)
	}

	err = w.cacheStorage.Set(ctx, cacheKey, weatherResponse, time.Minute*5)
	if err != nil {
		w.appConfig.Logger.Errorf("error when setting cache %v", err.Error())
	}

	return weatherResponse, nil
}
