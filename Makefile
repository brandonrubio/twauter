.PHONY: lint
lint:
	golangci-lint run \
		--config ./.golangci.yaml

.PHONY: compile
compile:
	go build .

.PHONY: build
build: lint compile
