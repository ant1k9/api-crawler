all: build

build:
	go build -o bin/crawler main.go

test:
	go test ./... -v -count=1 -coverprofile=coverage.txt -covermode=atomic

cov-html:
	go tool cover -html=coverage.txt

dump:

load: load-shares

load-shares:
	go run main.go crawl share
