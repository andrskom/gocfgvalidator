SRV = $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))
PROJECT = github.com/andrskom/${SRV}

all: lint test
.PHONY: all

test:
	@echo "+ $@"
	@go test -cover ./...
.PHONY: test

lint:
	@echo "+ $@"
	@docker run --rm -i  \
		-v ${GOPATH}/src/${PROJECT}:/go/src/${PROJECT} \
		-w /go/src/${PROJECT} golangci/golangci-lint:v1.12 golangci-lint run --enable-all --skip-dirs vendor,version,pkg/gen ./...
.PHONY: lint
