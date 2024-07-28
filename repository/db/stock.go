package db

import (
	"github.com/garbage-project/trash_api.git/repository/db/query"
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

	//query.StockDetailSelect...
	iterator := s.sql().
		Select(query.StockDetailSelect...).
		From(s.joinAs[s.stock]).
		LeftJoin(s.joinAs[s.stockImageInfo]).
		On("s.t_id = sii.t_id").
		Where(db.Cond{"s.t_id": stockId}).Iterator()

	defer func() {
		if err := iterator.Close(); err != nil {
			s.log.ErrLog("Failed To Close Iterator", "err", err.Error())
		}
	}()

	var res Stock
	if err := iterator.One(&res); err != nil {
		return nil, err
	} else {
		if len(res.Image) > 0 {
			if res.Image[0] == "" {
				res.Image = []string{}
			}
		}

		return &res, nil
	}
}
