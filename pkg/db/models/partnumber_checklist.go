package models

import "time"

type PartnumberChecklist struct {
	tableName struct{} `pg:"partnumber_checklist"`

	ID              int
	Partnumber      string
	InitalPrice     float32
	DateLastChecked time.Time
	Actual          bool
}
