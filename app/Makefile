.PHONY: all
all: test build

.PHONY: build
build:
	go build -tags prod

.PHONY: test
test:
	go test ./...

.PHONY: serve
serve:
	go run main.go