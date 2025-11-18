# Makefile for CryptoFortress Go Backend Services

# Default target
.PHONY: help
help: ## Show this help
	@echo "CryptoFortress Go Backend Services - Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build all services
.PHONY: build
build: ## Build all services
	@echo "Building all services..."
	cd auth && go build -o bin/auth-service cmd/main.go
	cd encryption && go build -o bin/encryption-service cmd/main.go
	cd keymgmt && go build -o bin/keymgmt-service cmd/main.go
	cd audit && go build -o bin/audit-service cmd/main.go
	@echo "All services built successfully."

# Run all services with docker-compose
.PHONY: run
run: ## Run all services with docker-compose
	@echo "Starting all services with docker-compose..."
	docker-compose up --build

# Run all services in detached mode
.PHONY: run-detached
run-detached: ## Run all services in detached mode
	@echo "Starting all services in detached mode..."
	docker-compose up --build -d

# Stop all services
.PHONY: stop
stop: ## Stop all services
	@echo "Stopping all services..."
	docker-compose down

# Test all services
.PHONY: test
test: ## Run tests for all services
	@echo "Running tests for all services..."
	cd auth && go test ./...
	cd encryption && go test ./...
	cd keymgmt && go test ./...
	cd audit && go test ./...

# Test all services with coverage
.PHONY: test-coverage
test-coverage: ## Run tests with coverage for all services
	@echo "Running tests with coverage for all services..."
	cd auth && go test ./... -cover
	cd encryption && go test ./... -cover
	cd keymgmt && go test ./... -cover
	cd audit && go test ./... -cover

# Clean build artifacts
.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf auth/bin
	rm -rf encryption/bin
	rm -rf keymgmt/bin
	rm -rf audit/bin
	@echo "Build artifacts cleaned."

# Update dependencies
.PHONY: deps
deps: ## Update dependencies for all services
	@echo "Updating dependencies for all services..."
	cd auth && go mod tidy
	cd encryption && go mod tidy
	cd keymgmt && go mod tidy
	cd audit && go mod tidy
	@echo "Dependencies updated."

# Run health checks
.PHONY: health
health: ## Run health checks for all services
	@echo "Running health checks for all services..."
	./test-services.sh

# Setup development environment
.PHONY: setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development environment setup complete."