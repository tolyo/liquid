# Shared dev configs
include ./pkg/conf/dev.env

# Backend
GO = go
GOFMT = $(GO) fmt
GOTEST = $(GO) test
GOLINT = golangci-lint
LDFLAGS="-s -w"  # To strip debug information from the binary (optional)
BINARY_NAME=liquid
BUILD_DIR=build
VENDORS_DIR=vendor
# Frontend
FRONTEND_CONTEXT = make -C web -f Makefile

# Default target to build the project
.PHONY: all
all: build

# Setup dev env
.PHONY: setup
setup:
	@echo "Setting up deps ..."
	@curl -Lo tigerbeetle.zip https://linux.tigerbeetle.com && unzip tigerbeetle.zip && ./tigerbeetle version
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/air-verse/air@latest
    $(FRONTEND_CONTEXT) setup
	@go get ./...
	

# Build the binary
.PHONY: build
build:
	@echo "Building the binary..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/

# Run the application
.PHONY: run
run:
	@echo "Running the application..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Run the application in dev env
.PHONY: dev
dev:
	$(MAKE) -j 3 frontend backend beetle

frontend:	
	$(FRONTEND_CONTEXT) serve

backend:
	@air demostore --env DEV

beetle: 
	./tigerbeetle start --addresses=$(TIGER_PORT) --development 0_0.tigerbeetle

# Install dependencies
.PHONY: install
install:
	@go mod tidy
	@go mod vendor

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -cover -v ./...

# Lint the code (use golangci-lint if installed)
.PHONY: lint
lint:
	@echo "Running linting..."
	$(GOLINT) run

# Format the code with gofmt
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)

DB_DSN:="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"
MIGRATE_OPTIONS=-allow-missing -dir="./sql"

db-up: ## Migrate down on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DB_DSN) up
	./tigerbeetle format --cluster=0 --replica=0 --replica-count=1 --development 0_0.tigerbeetle

db-down: ## Migrate up on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DB_DSN) reset
	rm 0_0.tigerbeetle

db-rebuild: ## Reset the database
	make db-down
	make db-up