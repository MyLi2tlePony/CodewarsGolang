test:
	go test -race -count 100 ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.0

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: test install-lint-deps lint