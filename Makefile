.PHONY: build test
build:
	 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo

test:
	@go test -v $(shell glide novendor)
