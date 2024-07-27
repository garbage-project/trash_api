package service

import (
	"errors"

	. "github.com/garbage-project/trash_api.git/types"
	"github.com/upper/db/v4"
)

func (s *Service) StockList(req StockListReq) ([]*StockList, error) {
	if res, err := s.repository.DB.StockList(req); err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return []*StockList{}, nil
		} else {
			s.log.ErrLog("Failed To Get StockList", "err", err.Error())
			return nil, err
		}
	} else {
		return res, nil
	}
}

func (s *Service) StockDetail(stockID int64) (*Stock, error) {
	return s.repository.DB.StockDetail(stockID)
}
