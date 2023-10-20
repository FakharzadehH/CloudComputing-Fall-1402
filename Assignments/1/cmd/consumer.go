package cmd

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

var consumerCMD = &cobra.Command{
	Use:   "consumer",
	Short: "Start the RabbitMQ consumer",
	Long:  `Start the RabbitMQ consumer`,
	Run:   consumerFunc,
}

func consumerFunc(cmd *cobra.Command, args []string) {
	rabbitMQ, err := amqp.Dial(config.GetConfig().RabbitMQ.GetURI())
	if err != nil {
		logger.Logger().Fatal(err)
	}
	db, err := config.NewGORMConnection(config.GetConfig())
	if err != nil {
		logger.Logger().Fatal(err)
	}
	logger.Logger().Info("consumer has started")
	repo := repository.NewRepository(db)
	svcs := service.New(repo)
	go func() { // this go routine checks rabbitMQ connection every 10 seconds
		for {
			time.Sleep(time.Second * 10)
			if rabbitMQ.IsClosed() {
				logger.Logger().Fatalw("RabbitMQ Connection is Closed!")
			}
		}
	}()
	go consume(repo, svcs, rabbitMQ) // this go routine consumes messages from rabbitMQ queue
	// stop if interrupt signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func consume(repo *repository.Repository, svcs *service.Service, conn *amqp.Connection) {
	rabbMQ := config.GetConfig().RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		logger.Logger().Fatalw("error while getting rabbitMQ connection channel")
	}
	queue, err := ch.QueueDeclare(rabbMQ.QueueName, false, false, false, false, nil)
	if err != nil {
		logger.Logger().Fatalw("error while declaring rabbitMQ connection")
	}
	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		logger.Logger().Fatalw("error while consuming messages from rabbitMQ")
	}
	go func() {
		for data := range msgs {
			userID, _ := strconv.Atoi(string(data.Body))
			if err := svcs.ProccessRequest(userID); err != nil {
				// do not ack msg from rabbitMQ if err != nil
				logger.Logger().Errorw("Error while proccessing user request", "error", err)
				continue
			}
			logger.Logger().Debugw("user authorization proccessed successfully", "user id", userID)
			data.Ack(true)
		}
	}()
}
