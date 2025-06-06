package producer

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"

	"github.com/Le-BlitzZz/apartments-rabbit-app/producer/internal/config"
)

func Run(conf *config.Config) {
	apartments, err := readCsv(conf.FilePath())
	if err != nil {
		log.Panicf("error reading CSV file: %v", err)
	}

	log.Infof("read %d apartments from CSV file", len(apartments))

	ch, err := conf.Broker().Channel()
	if err != nil {
		log.Panicf("error getting broker channel: %v", err)
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log.Errorf("failed to close channel: %v", err)
		} else {
			log.Info("channel closed successfully")
		}
	}()

	log.Info("broker channel opened")

	jsonData, err := json.Marshal(apartments)
	if err != nil {
		log.Panicf("failed to marshal apartments to JSON: %v", err)
	}

	log.Infof("marshaled apartments to JSON")

	ch.Publish(
		conf.BrokerRawExchange(),
		"",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         jsonData,
			DeliveryMode: amqp.Persistent,
		},
	)

	log.Info("published apartments to RabbitMQ exchange")
}
