package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var (
	Channel *amqp.Channel
)

func InitChannel() {
	// Get the connection string from the environment variable
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(url)

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("could not establish connection with RabbitMQ")
	}

	channel, err := connection.Channel()
	Channel = channel

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("could not open RabbitMQ channel")
	}

	err = Channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err,
			"errorCode": 500,
		}).Panic("could not declare exchange")
	}
}
