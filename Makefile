SHELL := /bin/bash

WORKDIR := $(shell pwd)

proto:
	@protoc -I proto/ proto/note.proto --go_out=plugins=grpc:proto
	@gofmt -w proto/*.go;

build: clean setup
	@echo "::build";
	cd $(WORKDIR)/note-api && \
	CGO_ENABLED=0 GOOS=linux go build -o /go/bin/note-api . && \
	cd $(WORKDIR)/note-svc && \
	CGO_ENABLED=0 GOOS=linux go build -o /go/bin/note-svc .;

clean:
	rm -rf /go/bin/note-api /go/bin/note-svc;

setup:
	@echo "::setup";
	go get -u github.com/micro/protoc-gen-micro;
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go};
	go get -u golang.org/x/lint/golint;
	go get -u golang.org/x/tools/cmd/cover;
