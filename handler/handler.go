package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ali-Assar/SkySpyBot/types"
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

	if err := SendLocation(body.Message.Chat.ID, body.Message.Text); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}
