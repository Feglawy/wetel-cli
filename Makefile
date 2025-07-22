BINARY=we_cli.exe
ARGS=

build:
	@go build -o bin/$(BINARY) cmd/wetel-cli/main.go

run: build
	@./bin/$(BINARY) $(ARGS)
