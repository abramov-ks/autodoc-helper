package telegram

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func (telegram *TelegramConfig) SendTelegramNotification(text string, shouldMute bool, parseMode string) (string, error) {

	if telegram.Token == "" || telegram.ChatId == 0 {
		return "", errors.New("TelegramConfig not cofigured")
	}

	log.Printf("Sending %s to chat_id: %d", text, telegram.ChatId)
	var telegramApi string = "https://api.telegram.org/bot" + telegram.Token + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":              {strconv.Itoa(telegram.ChatId)},
			"text":                 {text},
			"parse_mode":           {parseMode},
			"disable_notification": {strconv.FormatBool(shouldMute)},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of TelegramConfig Response: %s", bodyString)

	return bodyString, nil
}
