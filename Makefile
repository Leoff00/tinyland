PROJECT=tinyland_example
GPURL=github.com/leoff00/foo
ROOT=false

run:
	@go run attributes.go main.go -project $(PROJECT) -gpUrl $(GPURL) -root $(ROOT)
t:
	@go test -v 

bin:
	@go build -o bin/tinyland

.PHONY: run bin t