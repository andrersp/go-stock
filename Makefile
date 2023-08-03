.PHONY: test-coverage
test-coverage:
	go test -v ./... -coverprofile=coverage.out
