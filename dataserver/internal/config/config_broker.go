package config

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

const (
	brokerServer   = "rabbitmq:5672"
	brokerUser     = "apartments"
	brokerPassword = "apartments"
)

const (
	brokerRawExchange         = "raw_exchange"
	brokerTransformedExchange = "transformed_exchange"
)

const (
	brokerRawQueue         = "uploader_raw_queue"
	brokerTransformedQueue = "uploader_transformed_queue"
)

func (c *Config) Broker() *amqp.Connection {
	if c.broker == nil {
		log.Fatal("config: message broker not connected")
	}

	return c.broker
}

func (c *Config) BrokerUser() string {
	if c.options.BrokerUser == "" {
		return brokerUser
	}
	return c.options.BrokerUser
}

func (c *Config) BrokerPassword() string {
	if c.options.BrokerPassword == "" {
		return brokerPassword
	}
	return c.options.BrokerPassword
}

func (c *Config) BrokerServer() string {
	if c.options.BrokerServer == "" {
		return brokerServer
	}
	return c.options.BrokerServer
}

func (c *Config) BrokerRawExchange() string {
	if c.options.BrokerRawExchange == "" {
		return brokerRawExchange
	}
	return c.options.BrokerRawExchange
}

func (c *Config) BrokerTransformedExchange() string {
	if c.options.BrokerTransformedExchange == "" {
		return brokerTransformedExchange
	}
	return c.options.BrokerTransformedExchange
}

func (c *Config) BrokerRawQueue() string {
	if c.options.BrokerRawQueue == "" {
		return brokerRawQueue
	}
	return c.options.BrokerRawQueue
}

func (c *Config) BrokerTransformedQueue() string {
	if c.options.BrokerTransformedQueue == "" {
		return brokerTransformedQueue
	}
	return c.options.BrokerTransformedQueue
}

func (c *Config) BrokerDsn() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s/",
		c.BrokerUser(),
		c.BrokerPassword(),
		c.BrokerServer(),
	)
}

func (c *Config) connectBroker() error {
	brokerDsn := c.BrokerDsn()

	broker, err := amqp.Dial(brokerDsn)
	if err != nil || broker == nil {
		log.Infof("config: waiting for the message broker to become available")

		for range 5 {
			broker, err = amqp.Dial(brokerDsn)
			if broker != nil && err == nil {
				break
			}

			time.Sleep(2 * time.Second)
		}

		if err != nil || broker == nil {
			return err
		}
	}

	log.Info("RabbitMQ: connection established")

	c.broker = broker

	return nil
}

func (c *Config) closeBroker() error {
	if c.broker != nil {
		if err := c.broker.Close(); err != nil {
			return err
		}

		c.broker = nil
	}

	return nil
}
