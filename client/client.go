package client

import (
	"github.com/ncaak/pifiabot/models"
)

type client struct {
	Url string
}

func Build(botToken string) client {
	return client{
		Url: "https://api.telegram.org/bot" + botToken + "/",
	}
}

func (c client) SendMessage(data models.Output) {
	var endpoint = c.Url + "sendMessage"
	var body = getReplyJson(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}
