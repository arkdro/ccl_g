.PHONY: default

all:
	go build

test:
	go test -cover ./...

default: all
