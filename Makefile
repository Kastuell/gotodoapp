build:
	@go build cmd/main.go -o bin/gotodoapp

run: build
	@./bin/gotodoapp

test:
	@go test -v ./...