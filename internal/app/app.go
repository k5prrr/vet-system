package app

import (
	"context"
	"fmt"
	"log"
	"time"
)

type App struct {
	di *dependencyInjection
}

func New() *App {
	return &App{
		di: NewDependencyInjection(),
	}
}

func (a *App) Run(ctx context.Context) error {
	// Ждём, пока БД станет доступной
	for {
		if err := a.di.DB().Pool().Ping(ctx); err == nil {
			break
		}
		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	migration := a.di.Migration()
	if err := migration.Run(ctx); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	server := a.di.Server()
	err := server.Run(ctx)

	return err
}
