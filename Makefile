GOVERSION = $(shell go version | awk '{print $$3;}')
SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export CGO_ENABLED := 0

clean:
	rm -rf ./dist && rm -rf ./vendor
.PHONY: clean

upgrade:
	go get -t -u ./...
.PHONY: upgrade

vendor:
	go mod vendor
.PHONY: vendor

tidy:
	go mod tidy
.PHONY: tidy

lint:
	golangci-lint run --timeout=5m --output.text.colors
.PHONY: lint

test:
	gotestsum -- -tags=$(TEST_TAGS) -failfast -v -covermode count -timeout 5m $(SOURCE_FILES)
.PHONY: test

build:
	GOVERSION=$(GOVERSION) goreleaser release --snapshot --skip=sign,publish --clean --verbose
.PHONY: build

snapshot:
	GOVERSION=$(GOVERSION) goreleaser release --snapshot --clean --skip=sign --verbose
.PHONY: snapshot

release:
	GOVERSION=$(GOVERSION) goreleaser release --clean --skip=sign --verbose
.PHONY: release

docs:
	# Docs available at http://localhost:6060
	godoc -http=:6060
.PHONY: docs