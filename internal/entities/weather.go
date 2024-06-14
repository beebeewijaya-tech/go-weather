package entities

import "encoding/json"

type Weather struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type WeatherRequest struct {
	Weather
	AppID string
}

type WeatherResponse struct {
	Coord       Weather              `json:"coord"`
	WeatherInfo []WeatherInformation `json:"weather"`
	Main        WeatherMain          `json:"main"`
}

type WeatherInformation struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherMain struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Humidity float64 `json:"humidity"`
}

func (w WeatherResponse) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(w)
	return bytes, err
}
