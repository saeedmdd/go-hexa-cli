#! /usr/bin/make
PROJRCT_NAME=go-hexa-cli

all : tidy build

tidy:
	go mod tidy

run:
	go run .
build:
	rm -rf build
	@go build -tags netgo -ldflags='-extldflags "-static" -s -w' -o build/$(PROJRCT_NAME) -v

upgrade:
	go get -u all
fmt:
	@echo "Formatting code"
	gofmt -l -s -w ./
	goimports -l -w ./
