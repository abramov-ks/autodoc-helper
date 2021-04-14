package autodoc

import (
	"encoding/json"
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
func (session *AutodocSession) getPartnumbersUrl(partnumber string) string {
	return fmt.Sprintf(session.PartnumbersUrl, partnumber)
}

// проверить цену на запчасть
func (session AutodocSession) CheckPartnumber(partnumber string) (priceResponse *PartnumberPriceResponse, err error) {
	var url = session.getPartnumbersUrl(partnumber)
	var response, authError = session.makeAuthorizedGetRequest(url)

	if authError != nil {
		log.Printf("Error on partnumber request: %s", err)
		return nil, authError
	}
	return parseCheckPartnumberResponse(response)
}
