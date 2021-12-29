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

func (c client) sendRequest(r *http.Request) error {
	var client = &http.Client{
		Timeout: 30 * time.Second,
	}

	_, err := client.Do(r)
	if err != nil {
		log.Println("ERROR :: Sending HTTP request")
	}

	return err
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

func (c client) SendReply(method string, body models.Output) {
	var req, err = requests.Reply(c.Url+method, body)
	if err != nil {
		log.Println("ERROR :: Retrieving the request")
		return
	}

	var errRequest = c.sendRequest(req)
	if errRequest != nil {
		log.Println("ERROR :: Sending Reply request")
		return
	}
}

func (c client) SendMessage(data models.Output) {
	var endpoint = c.Url + "sendMessage"
	var body = requests.GetReplyBytes(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}
