.PHONY: default

all:
	go build

test:
	go test ./...

default: all
