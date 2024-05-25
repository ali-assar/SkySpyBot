package handler

import (
	"fmt"
	"log"
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/enescakir/emoji"
)

var OWMApiKey string

// SendWeather sends weather information for the given city to the given chat ID
func SendWeather(chatID int64, cityLocation string) error {
	var (
		description string
		icon        string
	)
	w, err := owm.NewCurrent("C", "EN", OWMApiKey)
	if err != nil {
		log.Println("Error creating OWM client:", err)
		return err
	}

	w.CurrentByName(cityLocation)

	if w.Weather != nil {
		description = w.Weather[0].Description
		icon = w.Weather[0].Icon
	} else {
		msg := "Data for requested location is not valid"
		return SendMessage(chatID, msg)
	}

	iconURL := fmt.Sprintf("http://openweathermap.org/img/wn/%s@4x.png", icon)

	if err := SendPhoto(chatID, iconURL); err != nil {
		log.Println("Error in sending icon:", err)
		return err
	}

	percentString := "%"
	msg := emoji.Sprintf(
		":satellite:Weather: %s\n:thermometer:Temperature: %.3f (Feels Like: %.3f)\n:droplet:Humidity: %v%s\n:sunrise:Sunrise: %s\n:sunset:Sunset: %s\n:dash:Wind Speed: %.3f KpH\n %v",
		description,
		w.Main.Temp,
		w.Main.FeelsLike,
		w.Main.Humidity,
		percentString,
		time.Unix(int64(w.Sys.Sunrise), 0).Format("15:04 MST"),
		time.Unix(int64(w.Sys.Sunset), 0).Format("15:04 MST"),
		w.Wind.Speed,
		time.Unix(int64(w.Dt), 0),
	)

	return SendMessage(chatID, msg)
}
