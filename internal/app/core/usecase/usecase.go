package usecase

import (
	"app/internal/app/core/port"
	"app/internal/app/core/service"
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
