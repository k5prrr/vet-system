package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
)

func (u *UseCase) CreateClient(ctx context.Context, token string, input *domain.Client) (int64, error) {
	actor, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return 0, err
	}

	entity := *input
	entity.ParentID = actor.ID

	id, err := u.service.RepoClient.Add(ctx, &entity)
	if err != nil {
		return 0, fmt.Errorf("add client: %w", err)
	}
	return id, nil
}

func (u *UseCase) Client(ctx context.Context, token string, id int64) (*domain.Client, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	client, err := u.service.RepoClient.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get client: %w", err)
	}
	return client, nil
}

func (u *UseCase) Clients(ctx context.Context, token string) ([]domain.Client, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	clients, err := u.service.RepoClient.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("list clients: %w", err)
	}
	return clients, nil
}

func (u *UseCase) UpdateClient(ctx context.Context, token string, id int64, input *domain.Client) error {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}

	_, err = u.service.RepoClient.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("client not found: %w", err)
	}

	entity := *input
	entity.ID = id
	return u.service.RepoClient.Update(ctx, id, &entity)
}

func (u *UseCase) DeleteClient(ctx context.Context, token string, id int64) error {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}

	return u.service.RepoClient.Delete(ctx, id, true) // soft-delete
}
