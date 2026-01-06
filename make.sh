#!/bin/bash

case "$1" in
    init)
        go mod init app
        go get github.com/jackc/pgx/v5
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
    base)
        #podman compose down
        podman compose -f db.yml down
        podman compose -f db.yml up -d
        podman ps -a
        ;;
    rm_base)
        ./make.sh stop_podman
        sudo rm pg_data -rf
        ;;
    stop_podman)
        podman stop -a
        podman rm -a
        podman ps -a
        ;;
    gpt)
        chmod +x scripts/createFileForGPT.sh
        scripts/createFileForGPT.sh
        ;;
    *)
        echo "Usage: $0 {build|run}"
        exit 1
        ;;
esac
