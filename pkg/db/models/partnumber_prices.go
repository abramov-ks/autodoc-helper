package models

import (
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"time"
)

type PartnumberPricesTable struct {
	tableName struct{} `pg:"partnumber_prices"`

	ID           int
	Partnumber   string
	DateChecked  time.Time
	MinimalPrice float32
	Info         autodoc.PartnumberPriceResponse
}
