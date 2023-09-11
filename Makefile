.PHONY: build all test clean

all: build 

build:
	go build -o bin/easycode-authenticator-backend ./src/cmd/authenticator-backend/main.go
	go build -o bin/easycode-authenticator-backend-internal ./src/cmd/authenticator-backend-internal/main.go

test:
	PROJECT_PWD=$(shell pwd) go test -race ./...

test-cover:
	go test -cover --count=1 ./...

cover-total:
	go test -cover --count=1 ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total 

cov:
	PROJECT_PWD=$(shell pwd) go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path './*_test.go')

fmt-check:
	@gofmt -l $(shell find . -type f -name '*.go' -not -path './*_test.go')

clean:
	@ro -fR bin