up-dev:
	docker compose -f .docker/docker-compose.dev.yml up -d

down-dev:
	docker compose -f .docker/docker-compose.dev.yml down

clean-dev:
	docker compose -f .docker/docker-compose.dev.yml down --volumes
