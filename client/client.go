package client

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/requests"
)

type client struct {
	Url string
}

func Build(botToken string) client {
	return client{
		Url: "https://api.telegram.org/bot" + botToken + "/",
	}
}

func send(url string, body []byte) {
	var req, errReq = http.NewRequest("POST", url, bytes.NewBuffer(body))
	if errReq != nil {
		log.Println("ERROR :: Creating Request Object : " + errReq.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{
		Timeout: 30 * time.Second,
	}

	var _, errDo = client.Do(req)
	if errDo != nil {
		log.Println("ERROR :: Sending HTTP request : " + errReq.Error())
	}
}

func (c client) SendMessage(data models.Output) {
	var endpoint = c.Url + "sendMessage"
	var body = requests.GetReplyBytes(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}
