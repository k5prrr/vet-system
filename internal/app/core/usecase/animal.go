package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
)

func (u *UseCase) CreateAnimal(ctx context.Context, token string, input *domain.Animal) (int64, error) {
	user, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return 0, err
	}
	entity := *input
	entity.ParentID = user.ID
	id, err := u.service.RepoAnimal.Add(ctx, &entity)
	if err != nil {
		return 0, fmt.Errorf("add animal: %w", err)
	}
	return id, nil
}

func (u *UseCase) Animal(ctx context.Context, token string, id int64) (*domain.Animal, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}
	animal, err := u.service.RepoAnimal.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get animal: %w", err)
	}
	return animal, nil
}

func (u *UseCase) Animals(ctx context.Context, token string) ([]domain.Animal, error) {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}
	animals, err := u.service.RepoAnimal.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("list animals: %w", err)
	}
	return animals, nil
}

func (u *UseCase) UpdateAnimal(ctx context.Context, token string, id int64, input *domain.Animal) error {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}
	_, err = u.service.RepoAnimal.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("animal not found: %w", err)
	}

	// Обновляем только поля из input
	entity := *input
	entity.ID = id

	return u.service.RepoAnimal.Update(ctx, id, &entity)
}

func (u *UseCase) DeleteAnimal(ctx context.Context, token string, id int64) error {
	_, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return err
	}
	return u.service.RepoAnimal.Delete(ctx, id, true) // soft-delete
}
