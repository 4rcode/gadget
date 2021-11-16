cover = go tool cover
test = go test ./... -cover -count=1
tmp = coverage.test

.PHONY: all clean generate details html update

all: generate
	$(test)

clean:
	go get ./...
	git clean -fdX

generate: clean
	go install github.com/4rcode/gnomic/cmd/optgen@1.0
	go generate ./...

details: clean $(tmp)
	$(cover) -func=$(tmp)

html: clean $(tmp)
	$(cover) -html=$(tmp)

update:
	go get -u ./...

$(tmp): generate
	$(test) -coverprofile=$(tmp)
