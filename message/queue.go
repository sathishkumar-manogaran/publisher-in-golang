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
	connection, err := connectToQueue()

	getChannel(err, connection)

	createExchangeQueue(err)
}

func createExchangeQueue(err error) {
	err = Channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err,
			"errorCode": 500,
		}).Panic("could not declare exchange")
	}
}

func getChannel(err error, connection *amqp.Connection) {
	channel, err := connection.Channel()
	Channel = channel

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("could not open RabbitMQ channel")
	}
}

func connectToQueue() (*amqp.Connection, error) {
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
	return connection, err
}
