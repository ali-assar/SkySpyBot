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
	/*
		redisAddress := os.Getenv("REDIS_ADDRESS")

		redisClient, err := database.NewRedisClient(redisAddress)
		if err != nil {
			log.Fatalf("Failed to create Redis client: %v", err)
		}

		httpHandler := func(res http.ResponseWriter, req *http.Request) {
			handler.Handler(res, req, redisClient)
		}
			}
	*/
	log.Println("running")
	http.ListenAndServe(":8080", http.HandlerFunc(handler.Handler))
}
