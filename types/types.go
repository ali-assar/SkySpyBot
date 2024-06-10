package types

type WebhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type SendPhotoReqBody struct {
	ChatID int64  `json:"chat_id"`
	Photo  string `json:"photo"`
}

type WeatherData struct {
	Description string
	Temperature float64
	FeelsLike   float64
	Humidity    int
	Sunset      int
	Sunrise     int
	WindSpeed   float64
	Dt          int
}
