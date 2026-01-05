package service

import (
	"app/internal/app/core/port"
)

type Service struct {
	salt     []byte
	Tg       port.Itg
	MapAuth  port.IMapAuth
	Repo     port.IRepo
	RepoUser port.IRepoUser
	RepoAuth port.IRepoAuth
}

func New(
	salt string,
	repoUser port.IRepoUser,
	repoAuth port.IRepoAuth,
) *Service {
	return &Service{
		salt:     []byte(salt),
		RepoUser: repoUser,
		RepoAuth: repoAuth,
	}
}
