APP_NAME=user_service
MAIN=cmd/server/main.go

include .env
export

.PHONY: run build wire test clean

wire:
	wire ./internal/wire/wire.go

run: wire
	go run -v $(MAIN)

build: wire
	go build -o bin/$(APP_NAME) $(MAIN)

test:
	go test ./... -cover

clean:
	rm -rf bin
