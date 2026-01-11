package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"errors"
	"fmt"
)

func (u *UseCase) CreateUser(ctx context.Context, token string, input *domain.User) (int64, error) {
	// Только админ может создавать пользователей
	admin, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return 0, err
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

	copyUser, err := u.service.RepoUser.GetBy(ctx, "phone", input.Phone)
	if err != nil && errors.Is(err, errors.New("not found")) {
		return 0, fmt.Errorf("get user by phone: %w", err)
	}
	if copyUser != nil {
		return 0, fmt.Errorf("phone already exists")
	}

	entity := *input
	entity.ParentID = admin.ID
	entity.PasswordHash = hashed
	entity.AuthSecret = secret

	newUserID, err := u.service.RepoUser.Add(ctx, &entity)
	if err != nil {
		return 0, fmt.Errorf("add new user: %w", err)
	}
	return newUserID, nil
}

func (u *UseCase) User(ctx context.Context, token string, id int64) (*domain.User, error) {
	// Только админ может просматривать конкретного пользователя
	if _, err := u.authorizeAdmin(ctx, token); err != nil {
		return nil, err
	}

	user, err := u.service.RepoUser.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Убираем чувствительные поля
	user.PasswordHash = ""
	user.AuthSecret = ""
	return user, nil
}

func (u *UseCase) Users(ctx context.Context, token string) ([]domain.User, error) {
	// Только админ может получать список всех пользователей
	var err error
	user, err := u.authorizeDoctorOrAdmin(ctx, token)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	if user.RoleID == RoleAdmin {
		users, err = u.service.RepoUser.List(ctx, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to list users: %w", err)
		}
	} else {
		users = []domain.User{*user}
	}

	for i := range users {
		users[i].PasswordHash = ""
		users[i].AuthSecret = ""
	}
	return users, nil
}

func (u *UseCase) UpdateUser(ctx context.Context, token string, id int64, input *domain.User) error {
	// Только админ может обновлять пользователей
	admin, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return err
	}

	// Запрет обновлять самого себя через этот путь (опционально)
	if id == admin.ID {
		return fmt.Errorf("cannot update yourself via admin endpoint")
	}

	targetUser, err := u.service.RepoUser.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("target user not found: %w", err)
	}

	// Подготавливаем обновлённую сущность
	updatedUser := *input
	updatedUser.ID = id
	updatedUser.AuthSecret = targetUser.AuthSecret

	if input.PasswordHash != "" {
		//fmt.Printf("Изменение пароля на %s\n", input.PasswordHash)
		hashed, err := u.service.Auth.PasswordHash(input.PasswordHash)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		updatedUser.PasswordHash = hashed
		//fmt.Printf("Хеш %s\n", hashed)
	} else {
		updatedUser.PasswordHash = targetUser.PasswordHash
	}

	return u.service.RepoUser.Update(ctx, id, &updatedUser)
}

func (u *UseCase) DeleteUser(ctx context.Context, token string, id int64) error {
	// Только админ может удалять пользователей
	admin, err := u.authorizeAdmin(ctx, token)
	if err != nil {
		return err
	}

	if id == admin.ID {
		return fmt.Errorf("cannot delete yourself")
	}

	return u.service.RepoUser.Delete(ctx, id, true) // soft-delete
}
