import processor.processor as processor
import processor.config as config


def _app():
    broker = config.connect_broker()

    try:
        processor.start(broker)
    finally:
        config.shutdown(broker)


def main():
    config.setup_logging()

    _app()


if __name__ == "__main__":
    main()
