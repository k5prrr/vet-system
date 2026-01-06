package port

import (
	"app/internal/app/core/domain"
	"context"
)

// Так, давай на первых этапах IUseCase будет как общий, а потом разделим
type IUseCase interface {
	/*CreateUser(ctx context.Context, tgID int64, phone string) (int64, error)
	SendAuthCode(ctx context.Context, phone string) error
	CheckAuthCode(ctx context.Context, phone, code string) (string, error)

	CurrentUser(ctx context.Context, token string) (*domain.User, error)
	*/
	Test1() (*domain.User, error)
}

type IRepoClient interface {
	Add(ctx context.Context, entity *domain.Client) (int64, error)

	Get(ctx context.Context, id int64) (*domain.Client, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.Client, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.Client, error)

	List(ctx context.Context, offset, limit int64) ([]domain.Client, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.Client, error)

	Update(ctx context.Context, id int64, entity *domain.Client) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.Client) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}
type IRepoTimesheet interface {
	Add(ctx context.Context, entity *domain.Timesheet) (int64, error)

	Get(ctx context.Context, id int64) (*domain.Timesheet, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.Timesheet, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.Timesheet, error)

	List(ctx context.Context, offset, limit int64) ([]domain.Timesheet, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.Timesheet, error)

	Update(ctx context.Context, id int64, entity *domain.Timesheet) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.Timesheet) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}

type IRepoAnimal interface {
	Add(ctx context.Context, entity *domain.Animal) (int64, error)

	Get(ctx context.Context, id int64) (*domain.Animal, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.Animal, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.Animal, error)

	List(ctx context.Context, offset, limit int64) ([]domain.Animal, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.Animal, error)

	Update(ctx context.Context, id int64, entity *domain.Animal) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.Animal) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}
type IRepoAnimalType interface {
	Add(ctx context.Context, entity *domain.AnimalType) (int64, error)

	Get(ctx context.Context, id int64) (*domain.AnimalType, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.AnimalType, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.AnimalType, error)

	List(ctx context.Context, offset, limit int64) ([]domain.AnimalType, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.AnimalType, error)

	Update(ctx context.Context, id int64, entity *domain.AnimalType) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.AnimalType) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}

type IRepoUser interface {
	Add(ctx context.Context, entity *domain.User) (int64, error)

	Get(ctx context.Context, id int64) (*domain.User, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.User, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.User, error)

	List(ctx context.Context, offset, limit int64) ([]domain.User, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.User, error)

	Update(ctx context.Context, id int64, entity *domain.User) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.User) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}
type IRepoUserRole interface {
	Add(ctx context.Context, entity *domain.UserRole) (int64, error)

	Get(ctx context.Context, id int64) (*domain.UserRole, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.UserRole, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.UserRole, error)

	List(ctx context.Context, offset, limit int64) ([]domain.UserRole, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.UserRole, error)

	Update(ctx context.Context, id int64, entity *domain.UserRole) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.UserRole) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}

type IRepoRecord interface {
	Add(ctx context.Context, entity *domain.Record) (int64, error)

	Get(ctx context.Context, id int64) (*domain.Record, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.Record, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.Record, error)

	List(ctx context.Context, offset, limit int64) ([]domain.Record, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.Record, error)

	Update(ctx context.Context, id int64, entity *domain.Record) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.Record) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}
type IRepoRecordStatus interface {
	Add(ctx context.Context, entity *domain.RecordStatus) (int64, error)

	Get(ctx context.Context, id int64) (*domain.RecordStatus, error)
	GetBy(ctx context.Context, filterKey, filterValue string) (*domain.RecordStatus, error)
	GetByInt(ctx context.Context, filterKey string, filterValue int64) (*domain.RecordStatus, error)

	List(ctx context.Context, offset, limit int64) ([]domain.RecordStatus, error)
	ListBy(ctx context.Context, filterKey, filterValue string, offset, limit int64) ([]domain.RecordStatus, error)

	Update(ctx context.Context, id int64, entity *domain.RecordStatus) error
	UpdateBy(ctx context.Context, filterKey, filterValue string, entity *domain.RecordStatus) error
	UpdateColumn(ctx context.Context, id int64, key, value string) error

	Delete(ctx context.Context, id int64, soft bool) error
	DeleteBy(ctx context.Context, filterKey, filterValue string, soft bool) error
}
