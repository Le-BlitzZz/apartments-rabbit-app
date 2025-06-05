import logging
import pika

BROKER_SERVER = "rabbitmq:5672"
BROKER_USER = "apartments"
BROKER_PASSWORD = "apartments"

BROKER_RAW_PROCESSING_EXCHANGE = "raw_exchange"
BROKER_PROCESSED_EXCHANGE = "transformed_exchange"

BROKER_RAW_PROCESSING_QUEUE = "processor_raw_queue"


def broker_dsn():
    return f"amqp://{BROKER_USER}:{BROKER_PASSWORD}@{BROKER_SERVER}/"


def connect_broker():
    return pika.BlockingConnection(pika.URLParameters(broker_dsn()))


def shutdown(connection):
    connection.close()
    logging.info("closed message broker connection")


def setup_logging(level=logging.INFO):
    logging.basicConfig(
        level=level,
        format="%(asctime)s %(levelname)s %(message)s",
        datefmt="%H:%M:%S",
    )
