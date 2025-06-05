terminal:
	docker compose exec -u root producer bash

build-go:
	rm -rf bin/producer
	go build -o bin/producer cmd/producer/producer.go

start:
	bin/producer

up:
	docker compose up

clean:
	docker compose down --volumes

build:
	docker compose build
