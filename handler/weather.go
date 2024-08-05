package handler

import (
	"errors"
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

	weatherData, iconData, err := RedisClient.GetWeather(cityLocation)

	if (err != nil) && (!errors.Is(err, database.WeatherGetError)) {
		log.Println("Error getting weather data from Redis:", err)
	} else if weatherData != nil && iconData != nil {
		description = string(weatherData)
		icon = string(iconData)
		log.Println("data is cached")
	} else {

		w.CurrentByName(cityLocation)
		log.Println("sent a query to get data")

		if w.Weather != nil {
			data := types.WeatherData{
				State:       w.Weather[0].Description,
				Temperature: w.Main.Temp,
				FeelsLike:   w.Main.FeelsLike,
				Humidity:    w.Main.Humidity,
				Sunset:      w.Sys.Sunrise,
				Sunrise:     w.Sys.Sunset,
				WindSpeed:   w.Wind.Speed,
			}

			description = CreateWeatherMsg(data) + fmt.Sprintf("\nCity:%s", cityLocation)
			icon = w.Weather[0].Icon

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

	return SendMessage(chatID, description)
}

func CreateWeatherMsg(data types.WeatherData) string {
	percentString := "%"
	msg := emoji.Sprintf(
		":satellite:Weather: %s\n:thermometer:Temperature: %.3f (Feels Like: %.3f)\n:droplet:Humidity: %v%s\n:sunrise:Sunrise: %s\n:sunset:Sunset: %s\n:dash:Wind Speed: %.3f KpH\n",
		data.State,
		data.Temperature,
		data.FeelsLike,
		data.Humidity,
		percentString,
		time.Unix(int64(data.Sunrise), 0).Format("15:04 MST"),
		time.Unix(int64(data.Sunset), 0).Format("15:04 MST"),
		data.WindSpeed,
	)
	return msg
}
