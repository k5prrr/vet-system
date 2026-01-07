package service

import (
	"app/internal/app/core/port"
)

type Service struct {
	Auth             port.IAuth
	RepoClient       port.IRepoClient
	RepoTimesheet    port.IRepoTimesheet
	RepoAnimal       port.IRepoAnimal
	RepoAnimalType   port.IRepoAnimalType
	RepoUser         port.IRepoUser
	RepoUserRole     port.IRepoUserRole
	RepoRecord       port.IRepoRecord
	RepoRecordStatus port.IRepoRecordStatus
}

func New(
	auth port.IAuth,
	repoClient port.IRepoClient,
	repoTimesheet port.IRepoTimesheet,
	repoAnimal port.IRepoAnimal,
	repoAnimalType port.IRepoAnimalType,
	repoUser port.IRepoUser,
	repoUserRole port.IRepoUserRole,
	repoRecord port.IRepoRecord,
	repoRecordStatus port.IRepoRecordStatus,
) *Service {
	return &Service{
		Auth:             auth,
		RepoClient:       repoClient,
		RepoTimesheet:    repoTimesheet,
		RepoAnimal:       repoAnimal,
		RepoAnimalType:   repoAnimalType,
		RepoUser:         repoUser,
		RepoUserRole:     repoUserRole,
		RepoRecord:       repoRecord,
		RepoRecordStatus: repoRecordStatus,
	}
}
