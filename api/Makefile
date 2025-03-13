BINARY_PATH=./tmp/server
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
	migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/expense-tracker?sslmode=disable" up
migrate-down:
	migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/expense-tracker?sslmode=disable" down 1
migrate-down-all:
	migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/expense-tracker?sslmode=disable" down