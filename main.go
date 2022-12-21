package main

import (
	"log"
	"net/http"

	"github.com/ncaak/pifiabot/client"
	"github.com/ncaak/pifiabot/config"
	"github.com/ncaak/pifiabot/models"
	"github.com/ncaak/pifiabot/server"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO :: Request received")

	defer r.Body.Close()

	var input = server.GetInput(r.Body)

	if input.IsCommand {

		// TODO : Handle command

		// stub
		var output = models.Output{
			ChatId:    input.ChatId,
			MessageId: input.MessageId,
			Text:      "pong",
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
		Certificate: config.Get().Certificate,
	}
	if err := client.Get().SetWebhook(webhook); err != nil {
		log.Fatal(err)
	}

	log.Println("INFO :: Starting the server...")
	var service, err = server.Build(config.Get().Certificate, config.Get().PrivateKey)
	if err != nil {
		log.Println("ERROR :: There was an error when building the service")
		log.Fatal(err)
	}

	service.AddRoute(config.Get().Url.Path, handleUpdate)

	log.Fatal(
		service.Listen(config.Get().Url.Port),
	)
}
