package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Ali-Assar/SkySpyBot/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	handler.OWMApiKey = os.Getenv("OWM_API_KEY")
	handler.TelegramApikey = os.Getenv("TELEGRAM_BOT_TOKEN")

	log.Println("running")
	http.ListenAndServe(":8080", http.HandlerFunc(handler.Handler))
}
