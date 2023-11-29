coberturaEsperada = 99
testcoverage = $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')

.PHONY: test-coverage
test-cov:
	@echo "  >  Running tests and generating coverage output ..."
	@echo "Cobertura esperada: $(coberturaEsperada)"
	@go test ./... -coverprofile coverage.out -covermode count
	@sleep 2
	@echo "Cobertura atual : $(testcoverage) %"
	@if [ "$(shell echo "$(testcoverage) < $(coberturaEsperada)" | bc -l)" -eq 1 ]; then \
        echo "Please add more unit tests or adjust the threshold to a lower value."; \
        echo "Failed"; \
        exit 1; \
    else \
        echo "OK"; \
    fi
	
	

.PHONY: run
run:
	swag init -g cmd/api/main.go --parseInternal
	go run cmd/api/main.go