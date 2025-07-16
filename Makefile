BINARY=we_cli
build:
	@go build -o bin/$(BINARY) cmd/wetel-cli/main.go

run: build
	@./bin/$(BINARY)