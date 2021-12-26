package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ncaak/pifiabot/models"
)

func getBytes(data interface{}) ([]byte, error) {
	var ba, err = json.Marshal(data)
	if err != nil {
		log.Println("ERROR :: Encoding json object")
		return []byte{}, err
	}

	return ba, nil
}

func Reply(url string, data models.Output) (*http.Request, error) {
	var body, errGetBytes = getBytes(models.Reply{
		ChatId:  data.ChatId,
		ReplyId: data.MessageId,
		Text:    data.Text,
	})
	if errGetBytes != nil {
		log.Println("ERROR :: Before creating Reply request")
		return &http.Request{}, errGetBytes
	}

	var req, errNewRequest = http.NewRequest("POST", url, bytes.NewBuffer(body))
	if errNewRequest != nil {
		log.Println("ERROR :: Creating Reply request")
		return &http.Request{}, errNewRequest
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
