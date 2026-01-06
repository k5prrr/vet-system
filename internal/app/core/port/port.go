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
