APP_NAME=workers
BUILD_DIR=bin
MAIN_PKG= ./cmd/workers

.PHONY: run build clean unit-test

run:
	go run $(MAIN_PKG)

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

clean:
	rm -rf ./bin/*

unit-test:
	go test ./... -v -cover -short=true
