#!/usr/bin/env bash

PROJECT_ROOT :=$(HOME)/LittleCai
OUTPATH :=$(HOME)/LittleCai/bin

all:build run

auto:
	sh $(PROJECT_ROOT)/backend/shell/auto.sh

.PHONY: static
static:
	sh $(PROJECT_ROOT)/backend/shell/mvStatic.sh

.PHONY: build
build:
	go build -v -o $(OUTPATH)/web main.go

.PHONY: run
run:
	sh $(OUTPATH)/run.sh restart

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o $(OUTPATH)/web_linux -v main.go

docker-build:
        #docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

clean:
	rm -rf $OUTPATH/*