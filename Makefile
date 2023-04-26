.PHONY: build test

# Go
GOPKG := go.indent.com/access

build: cmd/access
	go build -v $(GOPKG)/$<

test:
	go test -v ./...
