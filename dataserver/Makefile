terminal:
	docker compose exec -u root dataserver bash

build-go:
	rm -rf bin/dataserver
	go build -o bin/dataserver cmd/dataserver/dataserver.go

start:
	bin/dataserver

up:
	docker compose up

clean:
	docker compose down --volumes

build:
	docker compose build
