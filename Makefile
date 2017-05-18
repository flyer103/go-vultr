all:
	cd cmd/ && go build go-vultr.go && mv go-vultr ..

install: all

test:
	go test -cover ./...

clean:
	go clean -i ./...

gofmt:
	find . -name "*.go" | xargs -l1 go fmt

.PHONY : all install test clean gofmt
