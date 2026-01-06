// Dependency injection / Внедрение зависимостей
package app

import (
	"app/internal/app/adapter/api"
	"app/internal/app/core/port"
	"app/internal/app/core/service"
	"app/internal/app/core/usecase"
	"app/internal/repository/postgres"
	"app/migration"
	"app/pkg/database"
	"app/pkg/env"
	"app/pkg/server"
	"fmt"
)

type dependencyInjection struct {
	conf      *env.Env
	db        database.IDB
	migration *migration.Migration

	repoClient       port.IRepoClient
	repoTimesheet    port.IRepoTimesheet
	repoAnimal       port.IRepoAnimal
	repoAnimalType   port.IRepoAnimalType
	repoUser         port.IRepoUser
	repoUserRole     port.IRepoUserRole
	repoRecord       port.IRepoRecord
	repoRecordStatus port.IRepoRecordStatus

	service *service.Service
	useCase port.IUseCase

	router *api.Router
	server *server.Server
}

func NewDependencyInjection() *dependencyInjection {
	return &dependencyInjection{}
}
func (d *dependencyInjection) Conf() *env.Env {
	if d.conf == nil {
		d.conf = env.New(".env")
	}

	return d.conf
}
func (d *dependencyInjection) DB() database.IDB {
	if d.db == nil {
		conf := d.Conf()

		db, err := database.New(&database.DBConfig{
			User:     conf.Get("POSTGRES_USER", ""),
			Password: conf.Get("POSTGRES_PASSWORD", ""),
			Name:     conf.Get("POSTGRES_DB", ""),
			Host:     conf.Get("APP_POSTGRES_HOST", "localhost"),
			Port:     conf.Get("APP_POSTGRES_PORT", "15432"),
		})
		if err != nil {
			panic(fmt.Sprintf("failed to initialize database: %v", err))
		}
		d.db = db
	}

	return d.db
}
func (d *dependencyInjection) Migration() *migration.Migration {
	if d.migration == nil {
		db := d.DB()
		d.migration = migration.New(db.Pool())
	}

	return d.migration
}

func (d *dependencyInjection) RepoClient() port.IRepoClient {
	if d.repoClient == nil {
		d.repoClient = postgres.NewRepoClient(d.DB())
	}

	return d.repoClient
}
func (d *dependencyInjection) RepoTimesheet() port.IRepoTimesheet {
	if d.repoTimesheet == nil {
		d.repoTimesheet = postgres.NewRepoTimesheet(d.DB())
	}

	return d.repoTimesheet
}
func (d *dependencyInjection) RepoAnimal() port.IRepoAnimal {
	if d.repoAnimal == nil {
		d.repoAnimal = postgres.NewRepoAnimal(d.DB())
	}

	return d.repoAnimal
}
func (d *dependencyInjection) RepoAnimalType() port.IRepoAnimalType {
	if d.repoAnimalType == nil {
		d.repoAnimalType = postgres.NewRepoAnimalType(d.DB())
	}

	return d.repoAnimalType
}
func (d *dependencyInjection) RepoUser() port.IRepoUser {
	if d.repoUser == nil {
		d.repoUser = postgres.NewRepoUser(d.DB())
	}

	return d.repoUser
}
func (d *dependencyInjection) RepoUserRole() port.IRepoUserRole {
	if d.repoUserRole == nil {
		d.repoUserRole = postgres.NewRepoUserRole(d.DB())
	}

	return d.repoUserRole
}
func (d *dependencyInjection) RepoRecord() port.IRepoRecord {
	if d.repoRecord == nil {
		d.repoRecord = postgres.NewRepoRecord(d.DB())
	}

	return d.repoRecord
}
func (d *dependencyInjection) RepoRecordStatus() port.IRepoRecordStatus {
	if d.repoRecordStatus == nil {
		d.repoRecordStatus = postgres.NewRepoRecordStatus(d.DB())
	}

	return d.repoRecordStatus
}

func (d *dependencyInjection) Services() *service.Service {
	if d.service == nil {
		conf := d.Conf()
		d.service = service.New(
			conf.Get("APP_SALT", ""),
			d.RepoClient(),
			d.RepoTimesheet(),
			d.RepoAnimal(),
			d.RepoAnimalType(),
			d.RepoUser(),
			d.RepoUserRole(),
			d.RepoRecord(),
			d.RepoRecordStatus(),
		)
	}

	return d.service
}
func (d *dependencyInjection) UseCase() port.IUseCase {
	if d.useCase == nil {
		d.useCase = usecase.New(d.Services())
	}

	return d.useCase
}
func (d *dependencyInjection) Router() *api.Router {
	if d.router == nil {
		d.router = api.New(
			d.UseCase(),
		)
	}
	return d.router
}
func (d *dependencyInjection) Server() *server.Server {
	if d.server == nil {
		conf := d.Conf()
		d.server = server.New(
			conf.Get("APP_SERVER_ADDRESS", ":8080"),
			d.Router(),
			0,
		)
	}
	return d.server
}
