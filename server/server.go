package server

import (
	"encoding/json"
	"io"
	"log"

	"github.com/ncaak/pifiabot/models"
)

func GetInput(body io.ReadCloser) models.Input {
	var input = getInputFromUpdate(
		getUpdate(body),
	)

	return input
}

func getInputFromUpdate(update models.Update) models.Input {
	var entities = update.Message.Entities

	return models.Input{
		ChatId:    update.Message.Chat.Id,
		IsCommand: len(entities) > 0 && entities[0].Type == "bot_command",
		MessageId: update.Message.Id,
		Text:      update.Message.Text,
	}
}

func getUpdate(body io.ReadCloser) (update models.Update) {
	var decoder = json.NewDecoder(body)
	if err := decoder.Decode(&update); err != nil {
		log.Println("ERROR :: Decoding request body : " + err.Error())
	}

	return
}
