package client

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ncaak/pifiabot/models"
)

type API struct {
	Client *http.Client
	Url    string
}

var client *API

func Get() *API {
	return client
}

func Setup(botToken string) {
	var api = API{
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
		Url: "https://api.telegram.org/bot" + botToken + "/",
	}

	client = &api
}

func (api API) SetWebhook(data models.SetWebhook) error {
	var req, errReq = getMultipartRequest(api.getEndpoint("setWebhook"), data)
	if errReq != nil {
		log.Println("ERROR :: Setting up SetWebhook request")
		return errReq
	}

	//debugRequest(req) TODO: explicit debug mode

	var resp, errDo = api.Client.Do(req)
	if errDo != nil {
		log.Println("ERROR :: Sending HTTP request")
		return errDo
	}

	//debugResponse(resp)

	if resp.StatusCode != 200 {
		log.Println("ERROR :: Request was not successful")
		return errors.New(handleFailedResponse(resp))
	}

	return nil
}

func (api API) SendMessage(data models.Output) {
	var endpoint = api.getEndpoint("sendMessage")
	var body = getReplyJson(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}

func (api API) getEndpoint(method string) string {
	return api.Url + method
}
