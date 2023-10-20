package cmd

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/server"
	"github.com/spf13/cobra"
)

var apiCMD = &cobra.Command{
	Use:   "api",
	Short: "Start the rest API server",
	Long:  `Start the rest API server`,
	Run:   apiFunc,
}

func apiFunc(cmd *cobra.Command, args []string) {
	logger.Logger().Info("Starting http server")
	logger.Logger().Fatal(server.Start())
}
