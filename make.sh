#!/bin/bash

show_help() {
    echo "Usage: $0 <command>"
    echo
    echo "Commands:"
    # Ищем строки вида: `# help: <команда> - <описание>` непосредственно перед `case`-ветками
    grep -E '^[[:space:]]*# help:' "$0" | sed 's/^[[:space:]]*# help:[[:space:]]*//; s/ - /:  /' | column -t -s ':'
    exit 0
}

case "$1" in
    # help: init - инициализация модуля, зависимостей и .env
    init)
        go mod init app
        go get github.com/jackc/pgx/v5
        go mod tidy

        [ -f .env.example ] && [ ! -f .env ] && cp .env.example .env && echo ".env был создан из .env.example"
        ;;
    # help: build - сборка бинарника в ./bin/main
    build)
        go build -o bin/main ./cmd/app
        ;;
    # help: run - форматирование кода и запуск приложения
    run)
        go fmt ./cmd/... ./internal/... ./pkg/...
        clear
        go run cmd/app/main.go || echo "exit code: $?"
        ;;
    # help: base - запуск БД через Podman Compose (db.yml)
    base)
        podman compose -f db.yml down
        podman compose -f db.yml up -d
        podman ps -a
        ;;
    # help: rm_base - остановка Podman + удаление pg_data
    rm_base)
        ./make.sh stop_podman
        sudo rm -rf pg_data
        ;;
    # help: stop_podman - остановить и удалить все контейнеры Podman
    stop_podman)
        podman stop -a
        podman rm -a
        podman ps -a
        ;;
    # help: gpt - запуск скрипта подготовки файлов для GPT
    gpt)
        chmod +x scripts/createFileForGPT.sh
        scripts/createFileForGPT.sh
        ;;
    # help: help - показать это сообщение
    help|"")
        show_help
        ;;
    *)
        echo "Неизвестная команда: '$1'"
        echo
        show_help
        exit 1
        ;;
esac