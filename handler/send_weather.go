package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Ali-Assar/SkySpyBot/types"
	owm "github.com/briandowns/openweathermap"
)

func SendWeather(chatID int64, text string) error {

	w, err := owm.NewCurrent("C", "EN", OWMApiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(text)
	// Create the request body struct
	reqBody := &types.SendMessageReqBody{
		ChatID: chatID,
		Text:   fmt.Sprintf("temp= %v", int(w.Main.Temp)),
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	telegramApi := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TelegramApikey)
	// Send a post request with your token
	res, err := http.Post(telegramApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func SendMessage(chatID int64, text string) error {

	reqBody := &types.SendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	telegramApi := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TelegramApikey)
	// Send a post request with your token
	res, err := http.Post(telegramApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
