build:
	@go build -o bin/go-rest-bank-api

run: build
	@./bin/go-rest-bank-api

test:
	@go test -v ./...