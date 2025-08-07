# Project name 
APP_NAME := snippetbox

# Default help command
.PHONY: help
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make dev     → Start dev server using Air (live reload)"
	@echo "  make build   → Build the Go binary"
	@echo "  make run     → Run the app without Air"
	@echo "  make tidy    → Clean go.mod and go.sum"
	@echo "  make fmt     → Format all Go files"
	@echo "  make clean   → Remove the built binary"
	@echo ""

# Start Air 
.PHONY: dev
dev:
	@echo "Starting Air (live reload)..."
	air

# Build binary
.PHONY: build
build:
	go build -o $(APP_NAME) .

# Run directly without Air
.PHONY: run
run:
	go run .

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Tidy up go.mod and go.sum
.PHONY: tidy
tidy:
	go mod tidy

# Clean build
.PHONY: clean
clean:
	rm -f $(APP_NAME)
