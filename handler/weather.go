package handler

import (
	"fmt"
	"log"
	"time"

	database "github.com/Ali-Assar/SkySpyBot/db"
	"github.com/Ali-Assar/SkySpyBot/types"
	owm "github.com/briandowns/openweathermap"
	"github.com/enescakir/emoji"
)

var OWMApiKey string
var RedisClient database.RedisClient

// SendWeather sends weather information for the given city to the given chat ID
func SendWeather(chatID int64, cityLocation string) error {
	var (
		description string
		icon        string
	)

	w, err := owm.NewCurrent("C", "EN", OWMApiKey) // define w here
	if err != nil {
		log.Println("Error creating OWM client:", err)
		return err
	}

	// Try to get the weather data from Redis
	weatherData, iconData, err := RedisClient.GetWeather(cityLocation)
	if err != nil {
		log.Println("Error getting weather data from Redis:", err)
	} else if weatherData != nil && iconData != nil {
		// If the weather data is available in Redis, use it
		description = weatherData["description"]
		icon = iconData["icon"]
	} else {
		// If the weather data is not available in Redis, send a query to the API
		w.CurrentByName(cityLocation)

		if w.Weather != nil {
			description = w.Weather[0].Description
			icon = w.Weather[0].Icon

			// Cache the weather data in Redis
			err = RedisClient.SetWeather(cityLocation, description, icon)
			if err != nil {
				log.Println("Error setting weather data in Redis:", err)
			}
		} else {
			msg := "Data for requested location is not valid"
			return SendMessage(chatID, msg)
		}
	}

	iconURL := fmt.Sprintf("http://openweathermap.org/img/wn/%s@4x.png", icon)

	if err := SendPhoto(chatID, iconURL); err != nil {
		log.Println("Error in sending icon:", err)
		return err
	}

	data := types.WeatherData{
		Description: description,
		Temperature: w.Main.Temp,
		FeelsLike:   w.Main.FeelsLike,
		Humidity:    w.Main.Humidity,
		Sunset:      w.Sys.Sunrise,
		Sunrise:     w.Sys.Sunset,
		WindSpeed:   w.Wind.Speed,
		Dt:          w.Dt,
	}

	msg := CreateWeatherMsg(data)

	return SendMessage(chatID, msg)
}

func CreateWeatherMsg(data types.WeatherData) string {
	percentString := "%"
	msg := emoji.Sprintf(
		":satellite:Weather: %s\n:thermometer:Temperature: %.3f (Feels Like: %.3f)\n:droplet:Humidity: %v%s\n:sunrise:Sunrise: %s\n:sunset:Sunset: %s\n:dash:Wind Speed: %.3f KpH\n %v",
		data.Description,
		data.Temperature,
		data.FeelsLike,
		data.Humidity,
		percentString,
		time.Unix(int64(data.Sunrise), 0).Format("15:04 MST"),
		time.Unix(int64(data.Sunset), 0).Format("15:04 MST"),
		data.WindSpeed,
		time.Unix(int64(data.Dt), 0),
	)
	return msg
}
