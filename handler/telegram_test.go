package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Ali-Assar/SkySpyBot/types"
)

func TestParseUpdateMessageWithText(t *testing.T) {
	chat := types.Chat{Id: 12345}
	msg := types.Message{
		Text: "tehran",
		Chat: chat,
	}

	update := types.Update{
		UpdateId: 1,
		Message:  msg,
	}

	requestBody, err := json.Marshal(update)
	if err != nil {
		t.Errorf("Failed to marshal update in json, got %s", err.Error())
	}
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	updateToTest, errParse := parseTelegramRequest(req)
	if errParse != nil {
		t.Errorf("Expected a <nil> error, got %s", errParse.Error())
	}
	if *updateToTest != update {
		t.Errorf("Expected update %v, got %v", update, updateToTest)
	}

}
