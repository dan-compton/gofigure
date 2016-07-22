PROJECT := gofigure
ENTRYPOINT ?= gfserver
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

# Build like: make ENTRYPOINT="server" or make ENTRYPOINT="test-client"
build: proto
	CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' -o bin/$(ENTRYPOINT) ./cmd/$(ENTRYPOINT)
	docker build --build-arg ENTRYPOINT="/$(ENTRYPOINT)" --build-arg PROJECT=$(PROJECT) --build-arg VERSION=$(GITCOMMIT) -t dan-compton/$(PROJECT):$(GITCOMMIT) .

run: build
	docker run \
		dan-compton/$(PROJECT):$(GITCOMMIT)

push:
	docker push dan-compton/$(PROJECT):$(GITCOMMIT)

.PHONY: build run migrate all push
