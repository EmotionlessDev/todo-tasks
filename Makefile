include .env

up:
	docker compose up -d

up-log:
	docker compose up

build:
	docker compose build

down:
	docker compose down

clean:
	docker compose down --volumes

migrate-up:
	migrate -path=./migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up
