package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ali-Assar/SkySpyBot/types"
)

var TelegramApikey string

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &types.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("Could not decode request body:", err)
		return
	}

	switch body.Message.Text {
	case "/start":
		startMessage(body.Message.Chat.ID)
	case "/help":
		helpMessage(body.Message.Chat.ID)
	default:
		if err := SendWeather(body.Message.Chat.ID, body.Message.Text); err != nil {
			log.Println("Error in sending weather:", err)
			return
		}
	}

	log.Printf("Reply sent to %v", body.Message.Chat.ID)
}

func SendMessage(chatID int64, text string) error {
	reqBody := &types.SendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %w", err)
	}

	telegramApi := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TelegramApikey)
	res, err := http.Post(telegramApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s", res.Status)
	}

	return nil
}

func SendPhoto(chatID int64, photoURL string) error {
	reqBody := &types.SendPhotoReqBody{
		ChatID: chatID,
		Photo:  photoURL,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %w", err)
	}

	telegramApi := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", TelegramApikey)
	res, err := http.Post(telegramApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error sending photo: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s", res.Status)
	}

	return nil
}
