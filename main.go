package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ali-Assar/SkySpyBot/handler"
	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	apiKey := os.Getenv("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("mashhad")

	fmt.Println(w.Weather)
	fmt.Println(w.Main.Temp)

	http.HandleFunc("/webhook", handler.HandleTelegramWebHook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
