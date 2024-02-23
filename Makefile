build:
	@go build -o bin/go-rest-bank-api

run:
	@air

test:
	@go test -v ./...