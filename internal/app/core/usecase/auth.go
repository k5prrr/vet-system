package usecase

import (
	"app/internal/app/core/domain"
	"context"
	"fmt"
	"strings"
)

func (u *UseCase) CreateAdmin(ctx context.Context) error {
	parts := strings.Split(u.admin, "|")
	phone, password := parts[0], parts[1]

	user, _ := u.service.RepoUser.GetBy(ctx, "phone", phone)
	if user != nil {
		return fmt.Errorf("admin already exists")
	}

	secret, err := u.service.Auth.Secret(64)
	if err != nil {
		return fmt.Errorf("create admin Secret: %w", err)
	}

	passwordHash, err := u.service.Auth.PasswordHash(password)
	if err != nil {
		return fmt.Errorf("create admin passwordHash: %w", err)
	}

	_, err = u.service.RepoUser.Add(ctx, &domain.User{
		FIO:      "Админ админович",
		RoleID:   4,
		Phone:    phone,
		ParentID: 1,
		//Description:  nil,
		PasswordHash: passwordHash,
		AuthSecret:   secret,

		//CreatedAt: time.Now(),
		//UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("create admin: %w", err)
	}
	return nil
}
func (u *UseCase) Login(ctx context.Context, phone, password, ip string) (string, error) {
	user, err := u.service.RepoUser.GetBy(ctx, "phone", phone)
	if err != nil {
		return "", fmt.Errorf("Failed get user: %w", err)
	}

	if !u.service.Auth.PasswordVerify(password, user.PasswordHash) {
		return "", fmt.Errorf("invalid credentials")
	}

	token := u.service.Auth.Token(user.ID, user.AuthSecret, ip)

	return token, nil
}
func (u *UseCase) CurrentUser(ctx context.Context, token string) (*domain.User, error) {
	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	userID, err := u.service.Auth.UserID(token, "")
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	user, err := u.service.RepoUser.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	user.PasswordHash = ""
	user.AuthSecret = ""

	return user, nil
}
