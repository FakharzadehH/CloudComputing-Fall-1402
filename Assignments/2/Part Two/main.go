package main

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/server"
)

func main() {
	err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}
	logger.Init()
	logger.Logger().Fatal(server.Start())
}
