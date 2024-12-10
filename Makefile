.PHONY: test build run

test:
	go test -v -race ./...

build:
	go build -o bin/rest-server cmd/rest/server/main.go

run:
	bin/rest-server
