package app

import (
	"context"
	"fmt"
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

	migration := a.di.Migration()
	if err := migration.Run(ctx); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	server := a.di.Server()
	err := server.Run(ctx)

	return err
}
