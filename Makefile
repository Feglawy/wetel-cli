# Binary name and main source file
BINARY_NAME := wetel_cli
MAIN := cmd/wetel-cli/main.go
OUTPUT_DIR := bin
ARGS ?=



# Platform-specific build
.PHONY: build build-linux build-mac build-windows

.PHONY: all
all: build-linux build-mac build-windows

build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(MAIN)

run: build
	./$(OUTPUT_DIR)/$(BINARY_NAME) $(ARGS)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/linux/$(BINARY_NAME) $(MAIN)

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/mac/$(BINARY_NAME) $(MAIN)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/windows/$(BINARY_NAME).exe $(MAIN)

clean:
	rm -rf $(OUTPUT_DIR)
