package db

import (
	. "github.com/garbage-project/trash_api.git/types"
	"github.com/upper/db/v4"
)

func paging(page Paging, query interface{}) interface{} {
	if page.Page == 0 {
		page.Page = 1
	}

	if page.PageSize == 0 {
		page.PageSize = 10
	}

	switch query.(type) {
	case db.Selector:
		query = query.(db.Selector).Limit(int(page.PageSize)).Offset(int((page.Page - 1) * page.PageSize))
		return query
	case db.Result:
		query = query.(db.Result).Limit(int(page.PageSize)).Offset(int((page.Page - 1) * page.PageSize))
		return query
	default:
		return nil
	}
}
