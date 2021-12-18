GOFILES := $(shell find . -name "*.go")
CONFIG_ENV := default

.PHONY: check
check: fmt lint test

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	goimports -w -local github.com/hara/roomheatmap $(GOFILES)

.PHONY: build
build:
	./scripts/build.sh

.PHONY: deploy
deploy: build
	./scripts/deploy.sh $(CONFIG_ENV)
