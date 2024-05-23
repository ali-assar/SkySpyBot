package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ali-Assar/SkySpyBot/types"
	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

var (
	OWMApiKey      string
	TelegramApikey string
)

func Handler(res http.ResponseWriter, req *http.Request) {

	body := &types.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if err := SendLocation(body.Message.Chat.ID, body.Message.Text); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

func SendLocation(chatID int64, text string) error {

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

// FInally, the main funtion starts our server on port 3000
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	OWMApiKey = os.Getenv("OWM_API_KEY")
	TelegramApikey = os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Println("running")
	http.ListenAndServe(":8080", http.HandlerFunc(Handler))
}
