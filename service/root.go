package service

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/garbage-project/trash_api.git/config"
	"github.com/garbage-project/trash_api.git/repository"
)

type Service struct {
	cfg        *config.Config
	log        *util.Log
	repository *repository.Repository
}

func NewService(cfg *config.Config, log *util.Log, repository *repository.Repository) *Service {
	s := Service{cfg: cfg, log: log, repository: repository}

	return &s
}
