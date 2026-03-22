include .env
export

SHELL := /bin/bash

export PROJECT_ROOT=$(shell cygpath -m $(shell pwd))

env-up:
	@docker compose up -d go-list-postgres
env-down:
	@docker compose down go-list-postgres
env-cleanup:
	@read -p "Clean data? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down go-list-postgres go-list-port-forwarder && \
		rm -rf out/pgdata && \
		echo "Done"; \
	fi;
env-port-forward:
	@docker compose up -d go-list-port-forwarder
env-port-close:
	@docker compose down go-list-port-forwarder
migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "Need to set seq as argument. Example: make migrate-create seq=init"; \
  		exit 1; \
  	fi;
	@docker compose run --rm go-list-postgres-migrate \
		create \
		-ext sql \
		-dir ./migrations \
		-seq "$(seq)"
migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down
migrate-action:
	@if [ -z "$(action)" ]; then \
  		echo "Need to set action as argument. Example: make migrate-action action=up 1"; \
  		exit 1; \
  	fi;
	@docker compose run --rm go-list-postgres-migrate \
		-path ./migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@go-list-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)
golist-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=127.0.0.1 \
	go mod tidy && \
	go run cmd/golist/main.go


