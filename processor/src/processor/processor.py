import json
import logging
import sys
import pika
import processor.config as config
from processor.transformation import transform


def start(connection) -> None:
    channel = connection.channel()
    channel.basic_qos(prefetch_count=1)

    try:
        channel.exchange_declare(
            exchange=config.BROKER_RAW_PROCESSING_EXCHANGE,
            passive=True,
        )
        channel.exchange_declare(
            exchange=config.BROKER_PROCESSED_EXCHANGE,
            passive=True,
        )
        channel.queue_declare(
            queue=config.BROKER_RAW_PROCESSING_QUEUE,
            passive=True,
        )
    except pika.exceptions.ChannelClosedByBroker as e:
        logging.error("Topology missing: %s", e)
        sys.exit(1)

    def handle(ch, method, _props, body):
        try:
            raw_rows = json.loads(body)
            logging.info("Processor: consumed %d raw rows", len(raw_rows))

            processed_rows = transform(raw_rows)

            ch.basic_publish(
                exchange=config.BROKER_PROCESSED_EXCHANGE,
                routing_key="",
                body=json.dumps(processed_rows).encode(),
                properties=pika.BasicProperties(
                    content_type="application/json",
                    delivery_mode=2,
                ),
            )
            logging.info("Processor: published %d processed rows", len(processed_rows))

            ch.basic_ack(method.delivery_tag)
            ch.stop_consuming()
        except Exception as exc:
            logging.exception("Processor failed - dropping batch: %s", exc)
            ch.basic_nack(method.delivery_tag, requeue=False)

    channel.basic_consume(
        queue=config.BROKER_RAW_PROCESSING_QUEUE,
        on_message_callback=handle,
        auto_ack=False,
    )

    logging.info("Processor waiting for the single batch …")
    channel.start_consuming()

    channel.close()
    logging.info("Processor finished - channel closed")
