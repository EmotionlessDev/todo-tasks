up-dev:
	docker compose up -d

build-dev:
	docker compose build

down-dev:
	docker compose down

clean-dev:
	docker compose down --volumes
