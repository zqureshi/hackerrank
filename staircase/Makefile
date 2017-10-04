BIN = $(shell basename $(shell PWD))

.PHONY: all test
all: build test

build: *.go
	go build -o $(BIN) .

test: build
	./$(BIN) < input

clean:
	rm $(BIN)
