include .envrc

MIGRATIONS_PATH = ./cmd/migrate/migrations
DB_MIGRATOR_ADDR = postgres://postgres:Krish.jiyani%401@localhost:5432/socialnetworkx_ai?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down build run

# Target to create a new migration
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

# Target to apply all up migrations
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_MIGRATOR_ADDR) up

# Target to apply down migrations
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
 @go run cmd/migrate/seed/main.go
# Build the application
build:
	go build -o social.exe cmd/main.go

# Build and run the application
run: build
	./social.exe
