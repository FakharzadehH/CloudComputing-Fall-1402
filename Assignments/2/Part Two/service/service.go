package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain/payloads"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"github.com/go-redis/redis/v9"
)

type Service struct {
	repos *repository.Repository
}

func New(repos *repository.Repository) *Service {
	return &Service{
		repos: repos,
	}
}

func (s *Service) CheckWeather(ctx context.Context, payload payloads.CheckWeatherRequest) (*payloads.CheckWeatherResponse, error) {
	//first check if response is already in redis
	cfg := config.GetConfig().APP
	city := cfg.CityName
	if payload.CityName != "" {
		city = payload.CityName
	}
	redisKey := city
	val, err := s.repos.GetFromRedis(redisKey)
	if err != redis.Nil && err != nil {
		logger.Logger().Errorw("error while getting values from redis", "error", err)
		return nil, err
	}
	if err == nil {
		temps := strings.Split(val, ",")
		if len(temps) == 2 {
			logger.Logger().Debugw("Got the values from redis!", "values", temps)
			tempMin := temps[0]
			tempMax := temps[1]
			return &payloads.CheckWeatherResponse{
				TempMin: tempMin,
				TempMax: tempMax,
			}, nil
		}
	}

	//if there is no response in redis, send request to weatherAPI and store result in redis
	respBody, err := s.repos.GetWeatherStatus(city)
	if err != nil {
		return nil, err
	}
	apiResponse := payloads.WeatherAPIResponse{}
	err = json.Unmarshal(respBody, &apiResponse)
	if err != nil {
		logger.Logger().Errorw("eror while unmarshaling json", "error", err)
		return nil, err
	}
	if apiResponse.MinTemp == 0 && apiResponse.MaxTemp == 0 {
		return nil, errors.New("Can't send request to WeatherAPIResponse, is your vpn connected?")
	}

	tempMin := strconv.FormatInt(apiResponse.MinTemp, 10)
	tempMax := strconv.FormatInt(apiResponse.MaxTemp, 10)
	redisValue := tempMin + "," + tempMax
	if err := s.repos.StoreToRedis(redisKey, redisValue); err != nil {
		logger.Logger().Errorw("eror while saving values to redis", "error", err)
		return nil, err
	}
	logger.Logger().Debugw("Saved the values to redis!", "key", redisKey, "values", redisValue)

	return &payloads.CheckWeatherResponse{
		TempMin: tempMin,
		TempMax: tempMax,
	}, nil

}
