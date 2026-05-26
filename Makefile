.PHONY: all build test clean

all: build test

build:
	go build -o push-swap cmd/push-swap/main.go
	go build -o checker cmd/checker/main.go

test:
	go test ./... -v

clean:
	rm -f push-swap checker