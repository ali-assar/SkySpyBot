package services

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id

func SendTextToTelegramChat(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)

	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error in parsing telegram answer %s", err.Error())
		return "", err
	}

	return string(bodyBytes), nil
}
