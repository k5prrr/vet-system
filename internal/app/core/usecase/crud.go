package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
)

func (u *UseCase) Users(ctx context.Context, token string) ([]domain.User, error) {
	if token == "" {
		return nil, ErrUnauthorized
	}

	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w")
	}

	currentUser, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w")
	}

	if currentUser.RoleID != 4 {
		return nil, ErrForbidden
	}

	// Дополнительная защита на запрос от админа (только на изменение пользователей)
	userID, err = u.service.Auth.UserID(token, currentUser.AuthSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid token2: %w")
	}

	users, err := u.service.RepoUser.List(ctx, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	for i := range users {
		users[i].PasswordHash = ""
		users[i].AuthSecret = ""
	}

	return users, nil
}
func (u *UseCase) CreateUser(ctx context.Context, token string, input *domain.User) (int64, error) {
	if token == "" {
		return 0, ErrUnauthorized
	}

	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	currentUser, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("current user not found: %w", err)
	}
	if currentUser.RoleID != 4 { // 4 = admin
		return 0, ErrForbidden
	}

	// Дополнительная защита на запрос от админа (только на изменение пользователей)
	userID, err = u.service.Auth.UserID(token, currentUser.AuthSecret)
	if err != nil {
		return 0, fmt.Errorf("invalid token2: %w")
	}

	if input.PasswordHash == "" {
		return 0, fmt.Errorf("password is required")
	}

	hashed, err := u.service.Auth.PasswordHash(input.PasswordHash)
	if err != nil {
		return 0, fmt.Errorf("hash password: %w", err)
	}

	secret, err := u.service.Auth.Secret(64)
	if err != nil {
		return 0, fmt.Errorf("generate auth secret: %w", err)
	}

	entity := input
	entity.ParentID = currentUser.ID
	entity.PasswordHash = hashed
	entity.AuthSecret = secret

	newUserID, err := u.service.RepoUser.Add(ctx, entity)
	if err != nil {
		return 0, fmt.Errorf("add new user: %w", err)
	}

	return newUserID, nil
}

func (u *UseCase) UpdateUser(ctx context.Context, token string, id int64, input *domain.User) error {
	if token == "" {
		return ErrUnauthorized
	}

	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	currentUser, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}
	if currentUser.RoleID != 4 {
		return ErrForbidden
	}
	// Дополнительная защита на запрос от админа (только на изменение пользователей)
	userID, err = u.service.Auth.UserID(token, currentUser.AuthSecret)
	if err != nil {
		return fmt.Errorf("invalid token2: %w")
	}

	// Получаем обновляемого пользователя
	targetUser, err := u.service.RepoUser.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("target user not found: %w", err)
	}

	output := input
	output.ID = id
	output.AuthSecret = targetUser.AuthSecret

	// Если пришёл новый пароль - обновляем
	if input.PasswordHash != "" {
		hashed, err := u.service.Auth.PasswordHash(input.PasswordHash)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		output.PasswordHash = hashed
	} else {
		output.PasswordHash = targetUser.PasswordHash
	}

	return u.service.RepoUser.Update(ctx, id, targetUser)
}

func (u *UseCase) DeleteUser(ctx context.Context, token string, id int64) error {
	if token == "" {
		return ErrUnauthorized
	}

	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	currentUser, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}
	if currentUser.RoleID != 4 {
		return ErrForbidden
	}
	// Дополнительная защита на запрос от админа (только на изменение пользователей)
	userID, err = u.service.Auth.UserID(token, currentUser.AuthSecret)
	if err != nil {
		return fmt.Errorf("invalid token2: %w")
	}

	// запрет удалить самого себя
	if id == currentUser.ID {
		return fmt.Errorf("cannot delete yourself")
	}

	return u.service.RepoUser.Delete(ctx, id, true)
}
