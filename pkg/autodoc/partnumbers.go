package autodoc

import (
	"encoding/json"
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

// проверить цену на запчасть
func (session AutodocSession) CheckPartnumber(partnumber string) (priceResponse *PartnumberPriceResponse, err error) {
	var url = "https://webapi.autodoc.ru/api/spareparts/511/" + partnumber + "/2?isrecross=false"
	var response, authError = session.makeAuthorizedGetRequest(url)

	if authError != nil {
		log.Println("Error on partnumber request: %s", err)
		return nil, authError
	}
	return parseCheckPartnumberResponse(response)
}
