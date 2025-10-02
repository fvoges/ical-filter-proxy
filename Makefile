.PHONY: help test lint build changelog clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Run tests using Docker
	docker run --rm -v "$$PWD":/workspace -w /workspace golang:1.23-alpine go test -v ./...

test-coverage: ## Run tests with coverage
	docker run --rm -v "$$PWD":/workspace -w /workspace golang:1.23-alpine go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

lint: ## Run golangci-lint
	docker run --rm -v "$$PWD":/app -w /app golangci/golangci-lint:latest golangci-lint run -v

build: ## Build binary using Docker
	docker run --rm -v "$$PWD":/workspace -w /workspace golang:1.23-alpine go build -v -o ical-filter-proxy .

changelog: ## Generate CHANGELOG.md from git commits
	docker run --rm -v "$$PWD":/workdir -w /workdir quay.io/git-chglog/git-chglog -o CHANGELOG.md

changelog-next: ## Preview next version changelog (unreleased commits)
	docker run --rm -v "$$PWD":/workdir -w /workdir quay.io/git-chglog/git-chglog --next-tag v0.4.0 v0.3.0..

clean: ## Clean build artifacts
	rm -f ical-filter-proxy coverage.out

docker-build: ## Build Docker image
	docker build -t ical-filter-proxy:latest .

docker-run: ## Run Docker container locally
	docker run --rm -p 8080:8080 -v "$$PWD/config.yaml":/config.yaml ical-filter-proxy:latest

run: ## Run the application locally (requires Go installed)
	go run . -config config.yaml
