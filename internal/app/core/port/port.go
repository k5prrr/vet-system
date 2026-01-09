package port

import (
	"app/internal/app/core/domain"
	"context"
)

// Так, давай на первых этапах IUseCase будет как общий, а потом разделим
type IUseCase interface {
	CreateAdmin(ctx context.Context) error
	Login(ctx context.Context, phone, password, ip string) (string, error)
	CurrentUser(ctx context.Context, token string) (*domain.User, error)
	InitData(ctx context.Context, token string) (*domain.InitData, error)

	CreateUser(ctx context.Context, token string, input *domain.User) (int64, error)
	User(ctx context.Context, token string, id int64) (*domain.User, error)
	Users(ctx context.Context, token string) ([]domain.User, error)
	UpdateUser(ctx context.Context, token string, id int64, input *domain.User) error
	DeleteUser(ctx context.Context, token string, id int64) error

	CreateAnimal(ctx context.Context, token string, input *domain.Animal) (int64, error)
	Animal(ctx context.Context, token string, id int64) (*domain.Animal, error)
	Animals(ctx context.Context, token string) ([]domain.Animal, error)
	UpdateAnimal(ctx context.Context, token string, id int64, input *domain.Animal) error
	DeleteAnimal(ctx context.Context, token string, id int64) error

	CreateClient(ctx context.Context, token string, input *domain.Client) (int64, error)
	Client(ctx context.Context, token string, id int64) (*domain.Client, error)
	Clients(ctx context.Context, token string) ([]domain.Client, error)
	UpdateClient(ctx context.Context, token string, id int64, input *domain.Client) error
	DeleteClient(ctx context.Context, token string, id int64) error

	CreateRecord(ctx context.Context, token string, input *domain.Record) (int64, error)
	Record(ctx context.Context, token string, id int64) (*domain.Record, error)
	Records(ctx context.Context, token string) ([]domain.Record, error)
	UpdateRecord(ctx context.Context, token string, id int64, input *domain.Record) error
	DeleteRecord(ctx context.Context, token string, id int64) error

	CreateTimesheet(ctx context.Context, token string, input *domain.Timesheet) (int64, error)
	Timesheet(ctx context.Context, token string, id int64) (*domain.Timesheet, error)
	Timesheets(ctx context.Context, token string) ([]domain.Timesheet, error)
	UpdateTimesheet(ctx context.Context, token string, id int64, input *domain.Timesheet) error
	DeleteTimesheet(ctx context.Context, token string, id int64) error
}
type IAuth interface {
	Token(id int64, secret, ip string) string
	UserID(token, secret string) (int64, error)
	Secret(nBytes int) (string, error)
	Code() (int64, error)
	PasswordHash(password string) (string, error)
	PasswordVerify(password, encodedHash string) bool
}
