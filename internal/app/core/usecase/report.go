package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
)

func (u *UseCase) Report(ctx context.Context, token string) (*domain.Report, error) {
	// Только админ может запрашивать отчёт
	_, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	// Клиентов всего
	clientsTotal, err := u.service.RepoClient.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to count clients: %w", err)
	}

	// Записей всего
	recordsTotal, err := u.service.RepoRecord.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to count all records: %w", err)
	}

	// Отменённых записей (статус "cancel" = id 5)
	cancelled, err := u.service.RepoRecord.ListBy(ctx, "status_id", "5", 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to count cancelled records: %w", err)
	}

	return &domain.Report{
		ClientsTotal:     int64(len(clientsTotal)),
		RecordsTotal:     int64(len(recordsTotal)),
		RecordsCancelled: int64(len(cancelled)),
	}, nil
}
