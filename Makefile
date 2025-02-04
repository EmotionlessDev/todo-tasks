up-dev:
	docker compose -f .docker/docker-compose.dev.yml up -d

build-dev:
	docker compose -f .docker/docker-compose.dev.yml build

down-dev:
	docker compose -f .docker/docker-compose.dev.yml down

clean-dev:
	docker compose -f .docker/docker-compose.dev.yml down --volumes
