DB_BACKUP_PATH=/tmp/$(shell date +'%Y%m%d_%H%M%S').dump
DB_BACKUP_ARCHIVE=${DB_BACKUP_PATH}.zip

all: build

build:
	go build -o bin/crawler main.go

test:
	go test ./... -v -count=1 -coverprofile=coverage.txt -covermode=atomic

cov-html:
	go tool cover -html=coverage.txt

.PHONY: backup
backup:
	@pg_dump -Fc -d "$$DATABASE_URL" > "${DB_BACKUP_PATH}"
	@zip "${DB_BACKUP_ARCHIVE}" "${DB_BACKUP_PATH}" &>/dev/null
	@rm "${DB_BACKUP_PATH}"
	@echo "${DB_BACKUP_ARCHIVE}"

load: load-shares

load-shares: build
	./bin/crawler crawl shares
