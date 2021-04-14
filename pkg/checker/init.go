package checker

import (
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"github.com/abramov-ks/autodoc-helper/pkg/db"
	"github.com/abramov-ks/autodoc-helper/pkg/db/models"
	"log"
	"time"
)

func (config Config) SavePartNumberCheckHistory(partNumberInfo *autodoc.PartnumberPriceResponse) bool {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()
	partNumberRecord := models.PartnumberPricesTable{
		Partnumber:   partNumberInfo.PartNumber,
		DateChecked:  time.Now(),
		MinimalPrice: partNumberInfo.MinimalPrice,
		Info:         *partNumberInfo,
	}
	insertResult, insertError := dbInstance.Model(&partNumberRecord).Insert()

	if insertError != nil {
		log.Printf("Error while saving partnumber history: %s", insertError)
		return false
	}
	log.Printf("Saved partnumber history: %s", insertResult)
	return true
}

// Run Запуск
func (config Config) Run(partnumber string) {
	log.Println("run checker for user", config.Autodoc.Username)
	var autodocSession autodoc.AutodocSession
	autodocSession.FillFromConfig(&config.Autodoc)
	ok := autodocSession.Auth()
	if !ok {
		log.Println("Cannot create autodoc session")
	}

	if partnumber == "" {
		log.Println("No partnumber to check")
		return
	}

	partNumberInfo, partNumberInfoErr := autodocSession.CheckPartnumber(partnumber)
	if partNumberInfoErr != nil {
		log.Println("Cannot check partnumber price: %s", partNumberInfoErr)
		return
	}

	log.Printf("Деталь: %s цена: %.2f руб.", partNumberInfo.PartNumber, partNumberInfo.MinimalPrice)
	config.SavePartNumberCheckHistory(partNumberInfo)
}
