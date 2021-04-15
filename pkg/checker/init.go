package checker

import (
	"errors"
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"github.com/abramov-ks/autodoc-helper/pkg/db"
	"github.com/abramov-ks/autodoc-helper/pkg/db/models"
	"log"
	"time"
)

// SavePartNumberCheckHistory
func (config Config) SavePartNumberCheckHistory(partNumberInfo *autodoc.PartnumberPriceResponse) bool {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()
	partNumberRecord := models.PartnumberPricesTable{
		Partnumber:   partNumberInfo.PartNumber,
		DateChecked:  time.Now(),
		MinimalPrice: partNumberInfo.MinimalPrice,
		Info:         *partNumberInfo,
	}
	_, insertError := dbInstance.Model(&partNumberRecord).Insert()

	if insertError != nil {
		log.Printf("Error while saving partnumber history: %s", insertError)
		return false
	}

	checkListItem, checkListItemErr := config.UpdatePartNumberChecklistHistory(partNumberInfo.PartNumber)
	if checkListItemErr != nil {
		log.Println("Error on updating checklist", checkListItemErr)
	}

	log.Printf("Деталь: %s мин. цена %.2f", partNumberRecord.Partnumber, partNumberRecord.MinimalPrice)
	log.Printf("Изменение с первоначальной ценой %.2f", (partNumberRecord.MinimalPrice - checkListItem.InitalPrice))
	return true
}

// UpdatePartNumberChecklistHistory
func (config Config) UpdatePartNumberChecklistHistory(partNumber string) (*models.PartnumberChecklist, error) {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()

	var partnumberChecklist = new(models.PartnumberChecklist)
	selectError := dbInstance.Model(partnumberChecklist).Where("partnumber = ?", partNumber).Limit(1).Select()

	if selectError != nil {
		return nil, selectError
	}

	partnumberChecklist.DateLastChecked = time.Now()

	_, updateError := dbInstance.Model(partnumberChecklist).WherePK().Update()
	if updateError != nil {
		return nil, updateError
	}
	return partnumberChecklist, nil
}

// InsertPartNumberToChecklist
func (config Config) InsertPartNumberToChecklist(partNumberInfo *autodoc.PartnumberPriceResponse) (*models.PartnumberChecklist, error) {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()

	var partnumberChecklist = new(models.PartnumberChecklist)
	counter, selectError := dbInstance.Model(partnumberChecklist).Where("partnumber = ?", partNumberInfo.PartNumber).Limit(1).Count()
	if selectError != nil {
		return nil, selectError
	}

	if counter > 0 {
		return nil, errors.New("partnumber already exists in list")
	}

	newPartnumberChecklistItem := models.PartnumberChecklist{
		Partnumber:      partNumberInfo.PartNumber,
		InitalPrice:     partNumberInfo.MinimalPrice,
		DateLastChecked: time.Now(),
		Actual:          true,
	}

	_, insertError := dbInstance.Model(&newPartnumberChecklistItem).Insert()

	if insertError != nil {
		return nil, insertError
	}

	return &newPartnumberChecklistItem, nil
}

//
func (config Config) doPartnumberCheck(session *autodoc.AutodocSession, partNumber string) {
	if partNumber == "" {
		log.Println("No partNumber to check")
		return
	}

	partNumberInfo, partNumberInfoErr := session.GetPartnumberPrices(partNumber)
	if partNumberInfoErr != nil {
		log.Printf("Cannot check partNumber price: %s\n", partNumberInfoErr)
		return
	}

	config.SavePartNumberCheckHistory(partNumberInfo)

	return
}

//
func (config Config) doAddPartnumberForChecking(session *autodoc.AutodocSession, partNumber string) {
	if partNumber == "" {
		log.Println("No partnumber to check")
		return
	}

	partNumberInfo, partNumberInfoErr := session.GetPartnumberPrices(partNumber)

	if partNumberInfoErr != nil {
		return
	}

	insertedItem, err := config.InsertPartNumberToChecklist(partNumberInfo)
	if err != nil {
		log.Println("Cannot add to checklist: ", err)
		return
	}

	log.Println("Successfully added to checklist with id%", insertedItem.ID)

	return
}

// Run Запуск
func (config Config) Run(action *AppAction) {
	log.Println("run checker for user", config.Autodoc.Username)
	var autodocSession autodoc.AutodocSession
	autodocSession.FillFromConfig(&config.Autodoc)
	ok := autodocSession.Auth()
	if !ok {
		log.Println("Cannot create autodoc session")
	}

	if action.Action == "check" {
		config.doPartnumberCheck(&autodocSession, action.Value)
		return
	}

	if action.Action == "add" {
		config.doAddPartnumberForChecking(&autodocSession, action.Value)
		return
	}

}
