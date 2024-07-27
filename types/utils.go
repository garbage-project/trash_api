package types

type StockRole int8
type ReviewStatus int8

const (
	RestRoom StockRole = iota
	Trash
	Smoking
)

const (
	Removed ReviewStatus = iota
	Hide
	Expose
)

type Paging struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type Position struct {
	Hardness float64 `json:"hardness" db:"hardness"`
	Latitude float64 `json:"latitude" db:"latitude"`
}
