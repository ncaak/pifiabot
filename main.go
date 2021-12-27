package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ncaak/pifiabot/client"
	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/server"
)

func getInput(body io.ReadCloser) (input models.Input) {
	var update models.Update
	var decoder = json.NewDecoder(body)
	if err := decoder.Decode(&update); err != nil {
		log.Println("INFO :: Decoding request body : " + err.Error())
		return
	}

	if entities := update.Message.Entities; len(entities) == 0 || entities[0].Type != "bot_command" {
		log.Println("INFO :: Request body was not a Bot Command")
		return
	}

	return models.Input{
		ChatId:    update.Message.Chat.Id,
		MessageId: update.Message.Id,
		Text:      update.Message.Text,
	}
}

func webhook(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO :: Request received")

	var input = getInput(r.Body)

	var emptyInput models.Input
	if input != emptyInput {

		// TODO : Handle command

		var telegram = client.Build(os.Getenv("BOT_TOKEN"))

		// stub
		var output = models.Output{
			ChatId:    input.ChatId,
			MessageId: input.MessageId,
			Text:      "pong",
		}

		telegram.SendReply("sendMessage", output)
	}

}

func main() {

	fmt.Println("INFO :: Starting the server...")

	var service, err = server.Build()
	if err != nil {
		log.Println("ERROR :: There was an error when building the service")
		log.Fatal(err)
	}

	service.AddRoute("/v1/bot-api", webhook)

	log.Fatal(
		service.Listen("8443"),
	)
}
