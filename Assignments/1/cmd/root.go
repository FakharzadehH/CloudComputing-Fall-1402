package cmd

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Short: "ilenkrad",
	Long:  "iLenkrad GPS tracking platform",
}

func init() {
	err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}
	logger.Init()
	rootCMD.AddCommand(apiCMD)
	rootCMD.AddCommand(consumerCMD)
	rootCMD.AddCommand(monoCMD)
}

func Execute() {
	rootCMD.Execute()
}
