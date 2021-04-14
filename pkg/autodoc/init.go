package autodoc

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// FillFromConfig заполняем сессию из конфига
func (session *AutodocSession) FillFromConfig(config *AutodocConfig) {
	session.Username = config.Username
	session.Password = config.Password
	session.BaseUrl = config.BaseUrl
	session.AuthUrl = config.AuthUrl
	session.PartnumbersUrl = config.PartnumbersUrl
}

// GET запрос с авторизацией
func (session AutodocSession) makeAuthorizedGetRequest(url string) (response []byte, err error) {
	if session.AuthData.AccessToken == "" {
		return nil, errors.New("not authorized")
	}
	var bearer = "Bearer " + session.AuthData.AccessToken

	// Create a new request using http
	req, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return nil, requestErr
	}
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil, err
	}

	return body, nil
}
