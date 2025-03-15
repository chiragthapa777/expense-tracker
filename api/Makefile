# Load environment variables from .env
include .env

DATABASE_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
BINARY_PATH=./tmp/server.exe
MAIN_PATH=./cmd/server/main.go

build:
	go build -o $(BINARY_PATH) $(MAIN_PATH)
clean:
	rm $(BINARY_PATH)
start:
	$(BINARY_PATH)
run:
	go build -o $(BINARY_PATH) $(MAIN_PATH)
	$(BINARY_PATH)
dev:
	air -c ".air.toml"
migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down 1

migrate-down-all:
	migrate -path ./migrations -database "$(DATABASE_URL)" down