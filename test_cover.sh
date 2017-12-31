#!/bin/bash

# This script creates a single HTML file with the coverage of your project, per source file.
# from https://gist.github.com/toefel18/77dcdf892b35dcb955ececc784dbd91f

set -e
# clear coverage.out
cp /dev/null coverage.out

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.out
        rm profile.out
    fi
done

# remove duplicate lines (mode: count is present in each package)
awk '!a[$0]++' coverage.out > cover-profile.out
go tool cover -html=cover-profile.out -o coverage.html
rm coverage.out cover-profile.out
