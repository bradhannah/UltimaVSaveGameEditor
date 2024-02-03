PROJECT_NAME := UltimaVSaveGameEditor

.PHONY: all format lint test build clean

all: format lint test build

format:
	@echo "Formatting code..."
	@gofmt -s -w .

lint:
	@echo "Linting code..."
	@golint ./...

test:
	@echo "Running tests..."
	@go test ./...

build:
	@echo "Building $(PROJECT_NAME)..."
	@go build -o $(PROJECT_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(PROJECT_NAME)
