package domain

type NinjasResponse struct {
	CloudPct    int     `json:"cloud_pct,omitempty"`
	Temp        int     `json:"temp,omitempty"`
	FeelsLike   int     `json:"feels_like,omitempty"`
	Humidity    int     `json:"humidity,omitempty"`
	MinTemp     int64   `json:"min_temp,omitempty"`
	MaxTemp     int64   `json:"max_temp,omitempty"`
	WindSpeed   float64 `json:"wind_speed,omitempty"`
	WindDegrees int     `json:"wind_degrees,omitempty"`
	Sunrise     int64   `json:"sunrise,omitempty"`
	Sunset      int64   `json:"sunset,omitempty"`
}
