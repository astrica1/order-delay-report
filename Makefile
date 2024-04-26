.PHONY: migrate seed start test

migrate:
	go run ./cmd/migration/main.go

seed:
	go run ./cmd/seeder/main.go

start:
	go run ./cmd/server/main.go

test:
	go test ./... -cover -v -coverprofile=coverage.out