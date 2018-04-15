.PHONY: default

all:
	go build

test:
	go test -cover ./...

test_cover:
	./test_cover.sh

# run a single test:
# go test ./result/result_internal_test.go ./result/result.go -run Test_map_labels2

default: all
