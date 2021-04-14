package autodoc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Парсим ответ авторзиации
func parseAuthResponse(response []byte) (AuthResult, error) {
	var result = new(AuthResult)
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Printf("Error on unmarshaling auth response: %s", err)
		return *result, err
	}

	return *result, nil
}

// Авторизуемся в автодоке
func (session *AutodocSession) Auth() bool {
	response, err := http.PostForm(session.AuthUrl, url.Values{"username": {session.Username}, "password": {session.Password}, "grant_type": {"password"}})

	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Auth respnse code is %d", response.StatusCode)
		return false
	}

	bodyString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Can't read auth response")
		return false
	}

	var result, parseErr = parseAuthResponse(bodyString)

	if parseErr != nil {
		log.Printf("Auth response parse error %s", parseErr)
	}
	session.AuthData = result
	return true
}
