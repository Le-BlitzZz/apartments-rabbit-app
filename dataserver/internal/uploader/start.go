// uploader/run.go
package uploader

import (
	"context"
	"sync"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/config"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	log "github.com/sirupsen/logrus"
)

func Start(ctx context.Context, conf *config.Config) {
	broker := conf.Broker()

	ch, err := broker.Channel()

	if err != nil {
		log.Panic(err)
	}

	rawDeliveries, err := consumeDeliveries(ch, conf.BrokerRawQueue())
	if err != nil {
		log.Panicf("failed to start raw consuming: %v", err)
	}

	transformedDeliveries, err := consumeDeliveries(ch, conf.BrokerTransformedQueue())
	if err != nil {
		log.Panicf("failed to start transformed consuming: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		log.Info("starting raw deliveries processing")

		processDeliveries[entity.RawApartments](rawDeliveries)

		log.Info("raw deliveries processing stopped")
	}()

	go func() {
		defer wg.Done()

		log.Info("starting transformed deliveries processing")

		processDeliveries[entity.TransformedApartments](transformedDeliveries)

		log.Info("transformed deliveries processing stopped")
	}()

	<-ctx.Done()

	if err := ch.Close(); err != nil {
		log.Infof("failed to close channel: %v", err)
	} else {
		log.Info("channel closed successfully")
	}

	wg.Wait()

	log.Info("uploader stopped")
}
