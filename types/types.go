package types

type defaultStock struct {
	Name           string    `json:"name" db:"name"`
	Location       string    `json:"location" db:"location"`
	Management     string    `json:"management" db:"management"`
	DetailLocation string    `json:"detailLocation" db:"detail_location"`
	Description    string    `json:"description" db:"description"`
	Roles          StockRole `json:"roles" db:"roles"`

	Position

	RelatedPhoneNumber string `json:"relatedPhoneNumber" db:"related_phone_number"`
	IsValid            bool   `json:"isValid" db:"is_valid"`
}

type StockList struct {
	defaultStock
}

type Stock struct {
	defaultStock

	Image []string `json:"image" db:"image"`
}

type Reviews struct {
	User        string       `json:"user" db:"user"`
	Description string       `json:"description" db:"description"`
	Roles       StockRole    `json:"roles" db:"stock_roles"`
	CreatedDate int64        `json:"createdDate" db:"created_date"`
	Status      ReviewStatus `json:"status" db:"status"`
}
