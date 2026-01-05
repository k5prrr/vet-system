#!/bin/bash

case "$1" in
    init)
        go mod init app
        go mod tidy
        [ -f .env.example ] && [ ! -f .env ] && cp .env.example .env && echo ".env был создан из .env.example"
        ;;
    build)
        go build -o bin/main ./cmd/app
        ;;
    run)
        go fmt ./cmd/... ./internal/... ./pkg/...
        clear
        go run cmd/app/main.go || echo "exit code: $?"
        ;;
    *)
        echo "Usage: $0 {build|run}"
        exit 1
        ;;
esac
