.PHONY: all
all: test build

.PHONY: build
build:
	go generate ./...
	go build -tags prod

.PHONY: test
test:
	go test ./...

.PHONY: serve
serve:
	go run main.go