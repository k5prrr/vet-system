package service

import (
	"app/internal/app/core/port"
)

type Service struct {
	salt       []byte
	RepoUser   port.IRepoUser
	RepoRecord port.IRepoRecord
}

func New(
	salt string,
	repoUser port.IRepoUser,
	repoRecord port.IRepoRecord,
) *Service {
	return &Service{
		salt:       []byte(salt),
		RepoUser:   repoUser,
		RepoRecord: repoRecord,
	}
}
