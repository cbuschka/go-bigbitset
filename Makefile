PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test:
	@cd $(PROJECT_DIR) && \
	go test ./...

tidy:
	@cd $(PROJECT_DIR) && \
	go mod tidy
