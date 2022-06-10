MAIN_GO := main.go
BUILD_NAME := "invoice-app"

.PHONY: build
build:
	go build -mod=vendor -o bin/$(BUILD_NAME) $(MAIN_GO)

.PHONY: start
start: build startmongo
	./bin/$(BUILD_NAME)

.PHONY: startmongo
startmongo:
	docker run -d  --name mongodb -p 27017:27017 mongo:5.0.9 

.PHONY: stopmongo
stopmongo:
	docker stop mongodb 
	docker rm mongodb 