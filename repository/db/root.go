package db

import (
	"database/sql"
	"fmt"

	util "github.com/04Akaps/go-util/log"
	"github.com/garbage-project/trash_api.git/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

type SQL struct {
	cfg  config.DB
	sess db.Session
	log  *util.Log

	stock db.Collection

	joinAs map[db.Collection]string
}

func NewSQL(cfg config.DB, log *util.Log) *SQL {
	s := SQL{
		cfg:    cfg,
		log:    log,
		joinAs: make(map[db.Collection]string),
	}
	var err error

	if s.sess, err = mysql.Open(mysql.ConnectionURL{
		Database: cfg.Database,
		Host:     cfg.Host,
		User:     cfg.User,
		Password: cfg.Password,
		// Options:  map[string]string{}, // TODO 필요하다면 Group By를 위한 세션에서의 조건 수정
	}); err != nil {
		panic(err)
	} else if err = s.sess.Ping(); err != nil {
		panic(err)
	} else {

		register := func(name, as string) db.Collection {
			table := s.sess.Collection(name)
			s.joinAs[table] = fmt.Sprintf("%s AS %s", name, as)
			return table
		}

		s.stock = register("stock", "s")
		return &s
	}
}

func (s *SQL) sql() db.SQL {
	return s.sess.SQL()
}

func (s *SQL) GetPreparedRows(query string) (*sql.Stmt, error) {
	return s.sql().Prepare(query)
}
