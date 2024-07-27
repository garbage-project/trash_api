package repository

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/garbage-project/trash_api.git/config"
	"github.com/garbage-project/trash_api.git/repository/db"
)

type Repository struct {
	cfg *config.Config
	DB  *db.SQL
	log *util.Log
}

func NewRepository(cfg *config.Config, log *util.Log) *Repository {
	r := Repository{cfg: cfg, log: log}

	r.DB = db.NewSQL(cfg.DB["mysql"], log)

	return &r
}
