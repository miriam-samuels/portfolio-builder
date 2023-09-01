build:
	@go build -o cmd/portfolio

run:
	@go run cmd/portfolio

test:
	@go test -v ./...