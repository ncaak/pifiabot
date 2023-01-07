package main

import (
	"log"
	"net/http"

	"github.com/ncaak/pifiabot/actions"
	"github.com/ncaak/pifiabot/client"
	"github.com/ncaak/pifiabot/config"
	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/utils"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO :: Request received")
	var input, err = utils.GetRequestInput(r.Body)
	if err != nil {
		log.Println("ERROR :: Reading Telegram request : " + err.Error())
		return
	}

	if input.IsCommand {

		var command = actions.Factory(input.Text)
		var response, err = command.Resolve()
		if err != nil {
			log.Println("ERROR :: Resolving a command : " + err.Error()) // TODO: send message with command help
			return
		}

		var output = models.Output{
			ChatId:    input.ChatId,
			MessageId: input.MessageId,
			Text:      response,
		}
		client.Get().SendMessage(output)
	}
}

func main() {

	log.Println("INFO :: Setting up Configuration")
	if err := config.Setup(); err != nil {
		log.Fatal(err)
	}

	log.Println("INFO :: Setting up Webhook")
	client.Setup(config.Get().BotToken)
	var webhook = models.SetWebhook{
		Url:         config.Get().GetEndpoint(),
		Certificate: config.Get().CertBytes,
	}
	if err := client.Get().SetWebhook(webhook); err != nil {
		log.Fatal(err)
	}

	log.Println("INFO :: Starting the server...")
	http.HandleFunc("/"+config.Get().Url.Path, handleUpdate)

	log.Fatal(
		http.ListenAndServeTLS(
			":"+config.Get().Url.Port,
			config.Get().CertFilePath,
			config.Get().KeyFilePath,
			nil),
	)
}
