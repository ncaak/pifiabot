package client

import (
	"log"
	"net/http"
	"time"

	"github.com/ncaak/pifiabot/models"
)

type telegram struct {
	Client *http.Client
	Url    string
}

var client *telegram

func Get() *telegram {
	return client
}

func Setup(botToken string) {
	var api = telegram{
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
		Url: "https://api.telegram.org/bot" + botToken + "/",
	}

	client = &api
}

func (api telegram) SetWebhook(webhook models.SetWebhook) error {
	body, contentType := getMultipartBody(webhook)
	var req, errReq = getMultipartRequest(api.getEndpoint("setWebhook"), body, contentType)
	if errReq != nil {
		log.Println("ERROR :: Setting up SetWebhook request")
		return errReq
	}

	debugRequest(req)

	var resp, errDo = api.Client.Do(req)
	if errDo != nil {
		log.Println("ERROR :: Sending HTTP request")
		return errDo
	}

	debugResponse(resp)

	return nil
}

func (api telegram) SendMessage(data models.Output) {
	var endpoint = api.getEndpoint("sendMessage")
	var body = getReplyJson(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}

func (api telegram) getEndpoint(method string) string {
	return api.Url + method
}
