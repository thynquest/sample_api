MAIN_GO := main.go
BUILD_NAME := "invoice-app"

.PHONY: build
build:
	go build -mod=vendor -o bin/$(BUILD_NAME) $(MAIN_GO)

.PHONY: start
start: build
	./bin/$(BUILD_NAME)