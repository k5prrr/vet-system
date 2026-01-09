package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
)

func (u *UseCase) CreateTimesheet(ctx context.Context, token string, input *domain.Timesheet) (int64, error) {
	admin, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return 0, err
	}

	entity := *input
	entity.ParentID = admin.ID

	id, err := u.service.RepoTimesheet.Add(ctx, &entity)
	if err != nil {
		return 0, fmt.Errorf("add timesheet: %w", err)
	}

	return id, nil
}

func (u *UseCase) Timesheet(ctx context.Context, token string, id int64) (*domain.Timesheet, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	timesheet, err := u.service.RepoTimesheet.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get timesheet: %w", err)
	}

	return timesheet, nil
}

func (u *UseCase) Timesheets(ctx context.Context, token string) ([]domain.Timesheet, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	timesheets, err := u.service.RepoTimesheet.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("list timesheets: %w", err)
	}

	return timesheets, nil
}

func (u *UseCase) UpdateTimesheet(ctx context.Context, token string, id int64, input *domain.Timesheet) error {
	_, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return err
	}

	_, err = u.service.RepoTimesheet.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("timesheet not found: %w", err)
	}

	entity := *input
	entity.ID = id
	return u.service.RepoTimesheet.Update(ctx, id, &entity)
}

func (u *UseCase) DeleteTimesheet(ctx context.Context, token string, id int64) error {
	_, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return err
	}

	return u.service.RepoTimesheet.Delete(ctx, id, true) // soft-delete
}
