package uploader

import (
	"encoding/json"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func consumeDeliveries(ch *amqp.Channel, queue string) (<-chan amqp.Delivery, error) {
	return ch.Consume(queue, "", false, false, false, false, nil)
}

func processDeliveries[T entity.Apartments](deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		var apartments T

		if err := json.Unmarshal(delivery.Body, &apartments); err != nil {
			log.Errorf("failed to unmarshal messages: %v", err)
			delivery.Nack(false, false)
			continue
		}

		if err := apartments.Create(); err != nil {
			log.Errorf("failed to insert %s payload: %v", entity.TypeName(apartments), err)
			delivery.Nack(false, false)
			continue
		}

		log.Infof("successfully inserted %s payload", entity.TypeName(apartments))

		delivery.Ack(false)
	}
}
