package repository

import (
	"github.com/abramov-ks/autodoc-helper/pkg/db"
	"github.com/abramov-ks/autodoc-helper/pkg/db/models"
)

func GetPartnumbersChecklist(dbConfig db.DatabaseConfig) ([]models.PartnumberChecklist, error) {
	var connection = db.GetConnection(&dbConfig)
	var records []models.PartnumberChecklist

	err := connection.Model(&records).Order("id ASC").Where("actual = ?", true).Select()
	if err != nil {
		return nil, err
	}

	return records, nil
}
