package payloads

import "github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"

type CheckWeatherRequest struct {
	CityName string `json:"city_name,omitempty"`
}
type WeatherAPIResponse struct {
	domain.NinjasResponse
}
type CheckWeatherResponse struct {
	TempMin string `json:"temp_min,omitempty"`
	TempMax string `json:"temp_max,omitempty"`
}
