PROJECT_NAME := UltimaVSaveGameEditor
MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
MAIN_PATH := ${MAKEFILE_DIR}/cmd/ultima_v_save_game_editor
BIN_PATH := ${MAKEFILE_DIR}/bin


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
	cd ${MAIN_PATH}; go build -o ${BIN_PATH}/$(PROJECT_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(PROJECT_NAME)
