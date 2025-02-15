include .env

up:
	docker compose up -d

up-build:
	docker compose up -d --build

up-log:
	docker compose up

up-log-build:
	docker compose up --build

build:
	docker compose build

down:
	docker compose down

clean:
	docker compose down --volumes

migrate-up:
	migrate -path=./migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up

lint:
	golangci-lint run
