PROJECT_NAME=mix
DOCKER_COMPOSE_FILE=deployment/docker/docker-compose.yml
DOCKER_COMPOSE_CMD=docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${PROJECT_NAME}

.PHONY: all
all:

##
# Go modules section
##

.PHONY: update
update:
	go get -u all

.PHONY: prune
prune:
	go mod tidy

##
# Development
##

.PHONY: gen
gen:
	CGO_ENABLED=0 go generate -v ./... && scripts/protoc.sh

.PHONY: run_account
run_account:
	go run app/account/cmd/main.go --config=configs/account/config.toml

.PHONY: run_api
run_api:
	go run app/api/cmd/main.go --config=configs/api/config.toml

##
# Docker section for development
##

.PHONY: container_start
container_start:
	${DOCKER_COMPOSE_CMD} up -d ${CONTAINER}

.PHONY: container_stop
container_stop:
	${DOCKER_COMPOSE_CMD} stop ${CONTAINER} && ${DOCKER_COMPOSE_CMD} rm --force ${CONTAINER}

.PHONY: db_start
db_start:
	CONTAINER=mongo $(MAKE) container_start

.PHONY:db_stop
db_stop:
	CONTAINER=mongo $(MAKE) container_stop

.PHONY: docker_start
docker_start: db_start

.PHONY: docker_stop
docker_stop: db_stop
