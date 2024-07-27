package db

import (
	. "github.com/garbage-project/trash_api.git/types"
	"github.com/upper/db/v4"
)

func (s *SQL) StockList(req StockListReq) ([]*StockList, error) {
	// TODO image가 필요할까...?

	where := db.And(
		db.Cond{"is_valid": req.IsValid},
	)

	if len(req.Search) > 0 {
		// 	TODO Search
	}

	iterator := paging(
		req.Paging,
		s.sql().Select().From(s.joinAs[s.stock]).Where(where),
	).(db.Selector).Iterator()

	defer func() {
		if err := iterator.Close(); err != nil {
			s.log.ErrLog("Failed To Close Iterator", "err", err.Error())
		}
	}()

	var res []*StockList

	if err := iterator.All(&res); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (s *SQL) StockDetail(stockId int64) (*Stock, error) {
	return nil, nil
}
