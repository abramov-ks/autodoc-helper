package checker

import (
	"errors"
	"fmt"
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"github.com/abramov-ks/autodoc-helper/pkg/db"
	"github.com/abramov-ks/autodoc-helper/pkg/db/models"
	"github.com/abramov-ks/autodoc-helper/pkg/db/repository"
	"log"
	"strconv"
	"strings"
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

	_, checkListItemErr := config.UpdatePartNumberChecklistHistory(*partNumberInfo)
	if checkListItemErr != nil {
		log.Println("Error on updating checklist", checkListItemErr)
	}

	return true
}

func (config Config) getLastPartnumberInfo(partNumber string) (*autodoc.PartnumberPriceResponse, error) {

	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()

	var record = new(models.PartnumberPricesTable)
	selectError := dbInstance.Model(record).Where("partnumber = ?", partNumber).Limit(1).Order("id DESC").Select()

	if selectError != nil {
		return nil, selectError
	}

	return &record.Info, nil
}

// UpdatePartNumberChecklistHistory
func (config Config) UpdatePartNumberChecklistHistory(partNumber autodoc.PartnumberPriceResponse) (*models.PartnumberChecklist, error) {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()

	var partnumberChecklist = new(models.PartnumberChecklist)
	selectError := dbInstance.Model(partnumberChecklist).Where("partnumber = ?", partNumber.PartNumber).Limit(1).Select()

	if selectError != nil {
		return nil, selectError
	}

	partnumberChecklist.DateLastChecked = time.Now()
	partnumberChecklist.ManufacterId = partNumber.Manufacturer.Id

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
		return nil, errors.New("partnumber " + partNumberInfo.PartNumber + " already exists in list")
	}

	newPartnumberChecklistItem := models.PartnumberChecklist{
		Partnumber:      partNumberInfo.PartNumber,
		InitalPrice:     partNumberInfo.MinimalPrice,
		DateLastChecked: time.Now(),
		Name:            partNumberInfo.Name,
		Actual:          true,
		ManufacterId:    partNumberInfo.Manufacturer.Id,
	}

	_, insertError := dbInstance.Model(&newPartnumberChecklistItem).Insert()

	if insertError != nil {
		return nil, insertError
	}

	return &newPartnumberChecklistItem, nil
}

//
func (config Config) doPartnumberCheck(channel chan int, session *autodoc.AutodocSession, partNumber []string) {
	if partNumber[0] == "" {
		log.Println("No partNumber to check")
		return
	}
	manufacterId, _ := strconv.Atoi(partNumber[1])

	partNumberInfo, partNumberInfoErr := session.GetPartnumberPrices(partNumber[0], manufacterId)
	if partNumberInfoErr != nil {
		log.Printf("Cannot check partNumber price: %s\n", partNumberInfoErr)
		return
	}

	var previousPartNumberPriceInfo, checkError = config.getLastPartnumberInfo(partNumberInfo.PartNumber)

	config.SavePartNumberCheckHistory(partNumberInfo)

	if checkError != nil {
		log.Printf("No previous price info for %s", partNumberInfo.PartNumber)
		return
	}

	if previousPartNumberPriceInfo.MinimalPrice != partNumberInfo.MinimalPrice {
		var emoji string
		if partNumberInfo.MinimalPrice > previousPartNumberPriceInfo.MinimalPrice {
			emoji = "&#8599;"
		} else {
			emoji = "&#8600;"
		}
		var message = fmt.Sprintf("%s %s (%s) текущая мин. цена: %.2f руб., ", emoji, partNumberInfo.Name, partNumberInfo.PartNumber, partNumberInfo.MinimalPrice)
		message += fmt.Sprintf("изменение %.2f руб.", partNumberInfo.MinimalPrice-previousPartNumberPriceInfo.MinimalPrice)

		log.Printf("Send to telegram: %s", message)
		_, sendError := config.Telegram.SendTelegramNotification(message, false, "html")
		if sendError != nil {
			log.Printf("Telegram send error: %s", sendError)
			return
		}
	} else {
		log.Printf("No price changes for %s (%s): %8.2f\u20BD", partNumberInfo.DisplayPartNumber, partNumberInfo.Name, partNumberInfo.MinimalPrice)
	}
	channel <- 1
	return
}

func (config Config) doCheckAll(session *autodoc.AutodocSession) {
	log.Println("Run check all in list...")
	checkRecords, err := repository.GetPartnumbersChecklist(config.DataBase)
	log.Printf("Found %d for checking", len(checkRecords))
	if err != nil {
		log.Println("Cannot get checklist: ", err)
	}
	channel := make(chan int)
	counter := 0
	for _, record := range checkRecords {
		go config.doPartnumberCheck(channel, session, []string{record.Partnumber, strconv.Itoa(record.ManufacterId)})
		counter++
	}

	for i := 0; i < counter; i++ {
		<-channel
	}
}

//
func (config Config) doAddPartnumberForChecking(session *autodoc.AutodocSession, partNumbersWithComma []string) {
	log.Println("Run add action...")
	if partNumbersWithComma[0] == "" {
		log.Println("No partnumber to add")
		return
	}

	partNumbersArray := strings.Split(partNumbersWithComma[0], ",")
	manufacterId, _ := strconv.Atoi(partNumbersWithComma[1])
	for _, partNumber := range partNumbersArray {
		partNumberInfo, partNumberInfoErr := session.GetPartnumberPrices(partNumber, manufacterId)
		if partNumberInfoErr != nil {
			return
		}
		insertedItem, err := config.InsertPartNumberToChecklist(partNumberInfo)
		if err != nil {
			log.Println("Cannot add to checklist:", err)
		} else {
			log.Printf("%s successfully added to checklist with id #%d", insertedItem.Name, insertedItem.ID)
		}
	}

	return
}

func (config Config) doCleanup() int {
	dbInstance := db.GetConnection(&config.DataBase)
	defer dbInstance.Close()
	log.Println("Do cleanup...")
	partNumberRecord := models.PartnumberPricesTable{}

	res, err := dbInstance.Model(&partNumberRecord).Where("date_checked < NOW() - INTERVAL '1 week'").Delete()

	if err != nil {
		log.Printf("Cannot delete records: %s", err)
	}

	return res.RowsAffected()

}

// Run Запуск
func (config Config) Run(action *AppAction) int {
	log.Println("run checker for user", config.Autodoc.Username)
	var autodocSession autodoc.AutodocSession
	autodocSession.FillFromConfig(&config.Autodoc)
	ok := autodocSession.Auth()
	if !ok {
		log.Println("Cannot create autodoc session. Check login and password")
		return 1
	}

	if action.Action == "check" {
		channel := make(chan int)
		config.doPartnumberCheck(channel, &autodocSession, action.Value)
		<-channel
	}

	if action.Action == "add" {
		config.doAddPartnumberForChecking(&autodocSession, action.Value)
	}

	if action.Action == "check-all" {
		config.doCheckAll(&autodocSession)
	}

	if action.Action == "cleanup" {
		config.doCleanup()
	}

	return 0

}
