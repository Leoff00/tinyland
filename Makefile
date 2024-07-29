run:
	@go run attributes.go main.go -$(dirName)

t:
	@go test -v 

bin:
	@go build -o bin/tinyland

.PHONY: run bin t