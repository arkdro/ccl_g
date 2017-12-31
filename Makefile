.PHONY: default

all:
	go build

test:
	go test -cover ./...

test_cover:
	./test_cover.sh

default: all
