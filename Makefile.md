.PHONY: all build test clean

all: build test

build:
	@mkdir -p bin
	go build -o bin/push-swap cmd/push-swap/main.go
	go build -o bin/checker cmd/checker/main.go

test:
	go test ./... -v

clean:
	rm -rf bin