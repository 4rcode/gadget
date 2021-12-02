.PHONY: all clean generate details html test update

all: clean test

clean:
	git clean -fX

generate:
	go install github.com/4rcode/gnomic/cmd/optgen@latest
	go generate ./...

details: coverage.test
	go tool cover -func=$<

html: coverage.test
	go tool cover -html=$<

test: generate
	go test -cover -count=1 ./...

update:
	go get -u ./...

%.txt %.test: generate
	go test -cover -count=1 -coverprofile=$@ ./...
