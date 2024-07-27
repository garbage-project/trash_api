package service

import (
	"errors"

	. "github.com/garbage-project/trash_api.git/types"
	"github.com/upper/db/v4"
)

func (s *Service) Reviews(stockID int64) ([]*Reviews, error) {
	if res, err := s.repository.DB.GetReviews(stockID); err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return []*Reviews{}, nil
		} else {
			s.log.ErrLog("Failed To Get Reviews", "err", err.Error())
			return nil, err
		}
	} else {
		return res, nil
	}
}
