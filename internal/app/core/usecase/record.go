package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
	"strconv"
)

func (u *UseCase) CreateRecord(ctx context.Context, token string, input *domain.Record) (int64, error) {
	user, err := u.authorizeClientOrDoctorOrAdmin(ctx, token)
	if err != nil {
		return 0, err
	}

	entity := *input
	entity.ParentID = user.ID
	entity.ParentRoleID = user.RoleID

	if user.RoleID == RoleDoctor {
		entity.UserID = user.ID
	}
	if user.RoleID == RoleClient {
		entity.UserID = 0
		client, err := u.service.RepoClient.GetBy(ctx, "phone", user.Phone)
		if err != nil {
			return 0, fmt.Errorf("list animals: %w", err)
		}
		entity.ClientID = client.ID
	}

	id, err := u.service.RepoRecord.Add(ctx, &entity)
	if err != nil {
		return 0, fmt.Errorf("add record: %w", err)
	}
	return id, nil
}

func (u *UseCase) Record(ctx context.Context, token string, id int64) (*domain.Record, error) {
	user, err := u.authorizeClientOrDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	record, err := u.service.RepoRecord.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("record not found: %w", err)
	}

	// Если это не запись клиента, то запрет ему на просмотр
	if user.RoleID == RoleClient && record.ClientID != user.ID {
		return nil, ErrForbidden
	}

	return record, nil
}

func (u *UseCase) Records(ctx context.Context, token string) ([]domain.Record, error) {
	var err error

	user, err := u.authorizeClientOrDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	var records []domain.Record
	if user.RoleID == RoleAdmin {
		records, err = u.service.RepoRecord.List(ctx, 0, 0)
	} else if user.RoleID == RoleDoctor {
		records, err = u.service.RepoRecord.ListBy(ctx, "user_id", strconv.FormatInt(user.ID, 10), 0, 0)
	} else {
		client, err := u.service.RepoClient.GetBy(ctx, "phone", user.Phone)
		if err != nil {
			return nil, fmt.Errorf("list animals: %w", err)
		}
		records, err = u.service.RepoRecord.ListBy(ctx, "client_id", strconv.FormatInt(client.ID, 10), 0, 0)
	}
	if err != nil {
		return nil, fmt.Errorf("list records: %w", err)
	}

	return records, nil
}

func (u *UseCase) UpdateRecord(ctx context.Context, token string, id int64, input *domain.Record) error {
	actor, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}

	existing, err := u.service.RepoRecord.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	entity := *input
	entity.ID = id
	entity.ParentID = actor.ID
	entity.ParentRoleID = actor.RoleID

	// Сохраняем оригинальные значения, если новые не заданы
	if entity.ClientID == 0 {
		entity.ClientID = existing.ClientID
	}
	if entity.UserID == 0 {
		entity.UserID = existing.UserID
	}
	if entity.TimesheetID == 0 {
		entity.TimesheetID = existing.TimesheetID
	}
	if entity.StatusID == 0 {
		entity.StatusID = existing.StatusID
	}
	if entity.AnimalID == 0 {
		entity.AnimalID = existing.AnimalID
	}
	if entity.DateTime.IsZero() {
		entity.DateTime = existing.DateTime
	}
	if entity.Complaints == nil {
		entity.Complaints = existing.Complaints
	}
	if entity.Examination == "" {
		entity.Examination = existing.Examination
	}
	if entity.DS == nil {
		entity.DS = existing.DS
	}
	if entity.Recommendations == "" {
		entity.Recommendations = existing.Recommendations
	}
	if entity.Description == nil {
		entity.Description = existing.Description
	}

	return u.service.RepoRecord.Update(ctx, id, &entity)
}

func (u *UseCase) DeleteRecord(ctx context.Context, token string, id int64) error {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}

	return u.service.RepoRecord.Delete(ctx, id, true) // soft-delete
}
