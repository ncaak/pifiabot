package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

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

func getBytes(data interface{}) ([]byte, error) {
	var ba, err = json.Marshal(data)
	if err != nil {
		log.Println("ERROR :: Encoding json object")
		return []byte{}, err
	}

	return ba, nil
}

func getReplyBytes(data models.Output) []byte {
	var ba, err = getBytes(models.Reply{
		ChatId:  data.ChatId,
		ReplyId: data.MessageId,
		Text:    data.Text,
	})
	if err != nil {
		log.Println("ERROR :: JSON Marshaling Reply model : " + err.Error())
		return []byte{}
	}
	return ba
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
	//var body = requests.GetReplyBytes(data)
	var body = getReplyBytes(data)

	if len(body) > 0 {
		send(endpoint, body)
	}
}
