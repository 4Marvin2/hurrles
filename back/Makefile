HURRLES_BINARY=hurrles

PROJECT_DIR := ${CURDIR}

DOCKER_DIR := ${CURDIR}/docker

.PHONY: build-go
build-go:
	go build -o ${HURRLES_BINARY} cmd/hurrles/main.go

.PHONY: local-build
local-build:
	docker build -t hurrles_go_base -f ${DOCKER_DIR}/hurrles_go_base.Dockerfile .
	docker build -t hurrles_tarantool -f ${DOCKER_DIR}/hurrles_tarantool.Dockerfile .
	docker build -t hurrles -f ${DOCKER_DIR}/hurrles.Dockerfile .

.PHONY: local-run
local-run:
	docker-compose -f docker-compose.yml up --build --no-deps

.PHONY: local
local: local-build local-run

.PHONY: relocal
relocal: clean local-build local-run