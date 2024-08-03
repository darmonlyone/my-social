-include .env

run:
	go run cmd/server/main.go

STEP = 1

migrate-up:
	migrate -path postgres/migration -database $(POSTGRES_URL) up

migrate-down:
	migrate -path postgres/migration -database $(POSTGRES_URL) down $(STEP)

gensqlboiler:
	sqlboiler psql --config postgres/sqlboiler.toml

# gen:
# 	make gensqlboiler mockery swag-init

CURRENT_DIR := $(shell pwd)

infra-up:
	cd $(CURRENT_DIR)/deployment/postgres && make up

infra-down:
	cd $(CURRENT_DIR)/deployment/postgres && make down

.PHONY: test
test:
	sh script/test.sh