package usecase

import (
	"app/internal/app/core/domain"
	"app/internal/app/core/port"
	"app/internal/app/core/service"
	"context"
	"errors"
	"fmt"
)

// Тут он как основной, тут только создания
type UseCase struct {
	service *service.Service
	admin   string
}

func New(service *service.Service, admin string) port.IUseCase {
	return &UseCase{
		service: service,
		admin:   admin,
	}
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden: only admins can access this resource")
)

const (
	RoleClient = 2
	RoleDoctor = 3
	RoleAdmin  = 4
)

func (u *UseCase) authorize(ctx context.Context, token string) (*domain.User, error) {
	if token == "" {
		return nil, ErrUnauthorized
	}
	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	user, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if user.RoleID != RoleDoctor && user.RoleID != RoleAdmin {
		return nil, ErrForbidden
	}

	return user, nil
}

func (u *UseCase) authorizeDoctorOrAdmin(ctx context.Context, token string) (*domain.User, error) {
	user, err := u.authorize(ctx, token)
	if err != nil {
		return nil, err
	}
	if user.RoleID != RoleDoctor && user.RoleID != RoleAdmin {
		return nil, ErrForbidden
	}

	return user, nil
}
func (u *UseCase) authorizeAdmin(ctx context.Context, token string) (*domain.User, error) {
	user, err := u.authorize(ctx, token)
	if err != nil {
		return nil, err
	}
	if user.RoleID != RoleAdmin {
		return nil, ErrForbidden
	}

	return user, nil
}
