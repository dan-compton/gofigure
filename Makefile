APPENV ?= testenv
PROJECT := gofigure
GITCOMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
 	GITCOMMIT := $(GITCOMMIT)-dirty
	endif

all: build

fmt:
	@gofmt -w ./

proto:
	protoc --proto_path=./gofigure -I=./vendor --go_out=plugins=grpc:./gofigure ./gofigure/*.proto

build: proto $(APPENV)
	go build -o bin/server ./cmd/server/main.go
	docker build -t dan-compton/$(PROJECT):$(GITCOMMIT) .

run: $(APPENV)
	docker run \
		--env-file ./$(APPENV) \
		dan-compton/$(PROJECT):$(GITCOMMIT) server

push:
	docker push dan-compton/$(PROJECT):$(GITCOMMIT)

.PHONY: build run migrate all push
