.PHONY: build

build:
	@echo "Building..."
	@go build -o bin/slack-status main.go
