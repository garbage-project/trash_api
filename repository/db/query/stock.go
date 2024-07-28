package query

import "github.com/upper/db/v4"

var StockDetailSelect = []interface{}{
	"s.t_id",
	"s.location",
	"s.management",
	"s.detail_location",
	"s.description",
	"s.roles",
	"s.related_phone_number",
	"s.is_valid",
	"s.latitude",
	"s.hardness",

	db.Raw("JSON_ARRAYAGG(sii.image) AS images"),
}
