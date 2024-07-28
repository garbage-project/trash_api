package types

type StockListReq struct {
	Search  string `json:"search"`
	IsValid bool   `json:"isValid"`

	Paging
}

type StokeIdReq struct {
	ID int64 `uri:"id" binding:"required"`
}
