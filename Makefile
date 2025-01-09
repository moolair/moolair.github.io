# Task automation
# Basic test targets
.PHONY: run test build

run:
	go run cmd/blog-server/main.go

build:
	go build -o bin/blog-server cmd/blog-server/main.go

test:
	go test ./...