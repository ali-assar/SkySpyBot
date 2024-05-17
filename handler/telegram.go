package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ali-Assar/SkySpyBot/services"
	"github.com/Ali-Assar/SkySpyBot/types"
)

// parseTelegramRequest handles incoming update from the Telegram web hook

func parseTelegramRequest(r *http.Request) (*types.Update, error) {
	var update types.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	update, err := parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Assume the text is the location
	location := update.Message.Text

	// Fetch weather data for the location
	// For now, we'll just use a placeholder
	weatherData := "Weather data for " + location

	_, err = services.SendTextToTelegramChat(update.Message.Chat.Id, weatherData)
	if err != nil {
		log.Printf("error sending message to Telegram, %s", err.Error())
	}
}
