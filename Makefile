all: build

.PHONY: build
build:
	go build -o bin/crawler main.go
