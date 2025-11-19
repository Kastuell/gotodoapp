build:
	@go build cmd/app/main.go -o bin/gotodoapp

run: build
	@./bin/gotodoapp

test:
	@go test -v ./...