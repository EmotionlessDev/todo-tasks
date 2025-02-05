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
	migrate -path=./migrations -database "${TASKS_POSTGRES_DSN}" up
