package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ali-Assar/SkySpyBot/types"
	owm "github.com/briandowns/openweathermap"
	"github.com/enescakir/emoji"
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

	switch body.Message.Text {
	case "/start":
		startMessage(body.Message.Chat.ID)
	case "/help":
		helpMessage(body.Message.Chat.ID)
	default:
		if err := SendWeather(body.Message.Chat.ID, body.Message.Text); err != nil {
			fmt.Println("error in sending reply:", err)
			return
		}
	}

	// log a confirmation message if the message is sent successfully
	log.Printf("reply sent To %v", body.Message.Chat.ID)
}

func SendWeather(chatID int64, cityLocation string) error {
	var Description string
	w, err := owm.NewCurrent("C", "EN", OWMApiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(cityLocation)

	// Assuming 'w' contains the weather data
	if w.Weather != nil {
		var conditionCode = w.Weather[0].ID // Get the first condition code
		description, ok := weatherDescriptions[conditionCode]
		if !ok {
			description = "Unknown" // Fallback if the code is not in the map
		}
		Description = description
	} else {
		msg := "data for requested location is not valid"
		return SendMessage(chatID, msg)
	}
	// Create a single message with all the weather data
	msg := emoji.Sprintf(
		"Temperature: :thermometer: Temperature %.3f \nFeels Like: %.3f\nHumidity: :droplet: %v\nSunrise: :sunrise: %s\nSunset: :city_sunset: %s\nWind Speed: :dash: %.3f\nWeather: %s",
		w.Main.Temp,
		w.Main.FeelsLike,
		w.Main.Humidity,
		time.Unix(int64(w.Sys.Sunrise), 0).Format("15:04 MST"),
		time.Unix(int64(w.Sys.Sunset), 0).Format("15:04 MST"),
		w.Wind.Speed,
		Description,
	)

	// Send the message
	SendMessage(chatID, msg)
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
