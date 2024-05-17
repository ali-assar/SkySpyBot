package main

import (
	"log"
	"net/http"

	"github.com/Ali-Assar/SkySpyBot/handler"
)

func main() {
	http.HandleFunc("/webhook", handler.HandleTelegramWebHook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
