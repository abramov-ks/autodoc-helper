package autodoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// парсим ответ на цену детали
func parseCheckPartnumberResponse(response []byte) (responseStruc *PartnumberPriceResponse, err error) {
	var result = new(PartnumberPriceResponse)
	jsonError := json.Unmarshal(response, &result)
	if jsonError != nil {
		log.Printf("Error on unmarshaling parseCheckPartnumberResponse: %s", jsonError)
		return result, err
	}

	return result, nil
}

//getPartnumbersUrl
func (session *AutodocSession) getPartnumbersUrl(manufacterId int, partnumber string) string {
	return fmt.Sprintf(session.ApiUrl+"/api/spareparts/%d/%s/2?isrecross=false", manufacterId, partnumber)
}

func (session *AutodocSession) getManufacterInfoUrl(partnumber string) string {
	return fmt.Sprintf(session.ApiUrl+"/api/manufacturers/%s?showAll=false", partnumber)
}

func (session AutodocSession) CheckManufacter(partnumber string) (info *[]ManufacterInfo, err error) {
	var url = session.getManufacterInfoUrl(partnumber)
	var response, requestError = session.makeAuthorizedGetRequest(url)

	if requestError != nil {
		return nil, requestError
	}
	return parseCheckManufacterInfoResponse(response)

}

func parseCheckManufacterInfoResponse(response []byte) (*[]ManufacterInfo, error) {
	var result = new([]ManufacterInfo)
	jsonError := json.Unmarshal(response, &result)
	if jsonError != nil {
		return nil, jsonError
	}

	return result, nil
}

// проверить цену на запчасть
func (session AutodocSession) GetPartnumberPrices(partnumber string) (priceResponse *PartnumberPriceResponse, err error) {
	manufacterInfos, manufacterInfosError := session.CheckManufacter(partnumber)
	if manufacterInfosError != nil {
		log.Println("Cannot check manufacter info: %s", manufacterInfosError)
	}
	if len((*manufacterInfos)) < 1 {
		return nil, errors.New("Manufacter not found")
	}
	var url = session.getPartnumbersUrl((*manufacterInfos)[0].Id, partnumber)
	var response, authError = session.makeAuthorizedGetRequest(url)

	if authError != nil {
		log.Printf("Error on partnumber request: %s", err)
		return nil, authError
	}
	return parseCheckPartnumberResponse(response)
}
