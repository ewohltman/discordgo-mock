.PHONY: fmt lint test

fmt:
	gofmt -s -w . && goimports -local github.com/ewohltman/ephemeral-roles -w . && go mod tidy

lint: fmt
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.out ./...
	@ echo "all tests passed"
