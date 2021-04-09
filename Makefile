export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1
.PHONY: build
build: proto
	docker-compose build
.PHONY: up
up:
	docker-compose up -d
.PHONY: proto
proto:
	# make .proto
