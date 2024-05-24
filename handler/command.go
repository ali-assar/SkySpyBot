package handler

import "github.com/enescakir/emoji"

func startMessage(chatID int64) {
	// Handle /start command
	msg := emoji.Sprint("Hi I'm Sky Spy :wave:! \nA spy in the sky :cloud:. \nI can help you to know about current weather just by sending your city name. Use the /help command to learn how to use me.")
	SendMessage(chatID, msg)
}

func helpMessage(chatID int64) {
	// Handle /help command
	msg := "Send me a location and I'll tell you the current weather there."
	SendMessage(chatID, msg)

}
