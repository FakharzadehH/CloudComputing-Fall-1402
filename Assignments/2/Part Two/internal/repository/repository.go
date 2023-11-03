package repository

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/go-redis/redis/v9"
)

type Repository struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Repository {
	return &Repository{
		redis: redis,
	}
}

func (r *Repository) GetWeatherStatus(city string) ([]byte, error) {
	cfg := config.GetConfig()
	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=" + city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", cfg.NinjasAPI.APIKey)
	req.Header.Add("X-RapidAPI-Host", cfg.NinjasAPI.APIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger().Errorw("Error while getting response from Weather api")
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Logger().Errorw("Error while reading response body")
		return nil, err
	}
	return body, nil
}
func (r *Repository) StoreToRedis(key string, value string) error {
	duration := time.Duration(config.GetConfig().Redis.Expiry)
	return r.redis.Set(context.Background(), key, value, time.Second*duration).Err()
}

func (r *Repository) GetFromRedis(key string) (string, error) {
	res, err := r.redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
