package cmd

import (
	"os"
	"os/signal"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/spf13/cobra"
)

var monoCMD = &cobra.Command{
	Use:   "mono",
	Short: "mono",
	Long:  "run everything",
	Run:   monoFunc,
}

func monoFunc(cmd *cobra.Command, args []string) {
	go apiFunc(cmd, args)
	go consumerFunc(cmd, args)

	logger.Logger().Info("Starting assignment-1 mono")
	// stop if interrupt signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

}
