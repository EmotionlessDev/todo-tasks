up:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down

clean:
	docker compose down --volumes
