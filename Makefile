DOCKER_COMPOSE_VERSION=1.24.0
NAMESPACE=sr2020
SERVICE := platform
IMAGE := $(or ${image},${image},gateway)
IMAGE_TEST := $(or ${image},${image},gateway-convey)
GIT_TAG := $(shell git tag -l --points-at HEAD | cut -d "v" -f 2)
TAG := :$(or ${tag},${tag},$(or ${GIT_TAG},${GIT_TAG},latest))
ENV := $(or ${env},${env},local)
cest := $(or ${cest},${cest},)

include .env
export $(shell sed 's/=.*//' .env)

current_dir = $(shell pwd)

build:
	docker build -t ${NAMESPACE}/${IMAGE}${TAG} -t ${NAMESPACE}/${IMAGE}:latest .

build-test:
	docker build -t ${NAMESPACE}/${IMAGE_TEST}${TAG} -t ${NAMESPACE}/${IMAGE_TEST}:latest ./tests

push:
	docker push ${NAMESPACE}/${IMAGE}

up:
	docker-compose up -d

down:
	docker-compose down

reload:
	make down
	make up

restart:
	docker-compose down -v
	docker-compose up -d

install:
	cp .env.example .env

install-docker-compose:
	curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-Linux-x86_64 > /tmp/docker-compose
	chmod +x /tmp/docker-compose
	sudo mv /tmp/docker-compose /usr/local/bin/docker-compose
	docker-compose -v

test:
	cd tests && go clean -testcache && go test -v .

convey:
	cd tests && goconvey -port 8448

load:
	docker run -v $(current_dir)/tests/loadtest:/var/loadtest --net host --entrypoint /usr/local/bin/yandex-tank -it direvius/yandex-tank -c production.yaml

test-dev:
	make restart
	make test
