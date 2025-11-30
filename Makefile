.PHONY: help run test new clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run a specific day (usage: make run DAY=01)
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY is required. Usage: make run DAY=01"; \
		exit 1; \
	fi
	@go run ./day$(DAY)

test: ## Run all tests
	@go test ./...

test-day: ## Run tests for a specific day (usage: make test-day DAY=01)
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY is required. Usage: make test-day DAY=01"; \
		exit 1; \
	fi
	@go test ./day$(DAY)

new: ## Create a new day (usage: make new DAY=01)
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY is required. Usage: make new DAY=01"; \
		exit 1; \
	fi
	@./scripts/new_day.sh $(DAY)

clean: ## Remove build artifacts
	@go clean
	@rm -rf bin/
