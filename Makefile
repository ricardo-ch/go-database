.PHONY: build test
build:
	 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo

test:
	@go test -race -v $(shell go list ./... | grep -v /vendor/)
