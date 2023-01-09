package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type config struct {
	BotToken  string
	CertBytes []byte
	File      struct {
		Certificate string
		Messages    string
		PrivateKey  string
	}
	MessageDict map[string]string
	Url         struct {
		Endpoint string
		Path     string
		Port     string
	}
}

var configuration *config

func Get() *config {
	return configuration
}

func Setup() error {
	var config = config{}
	config.Url.Path = "updates" // TODO: Randomize url path
	config.Url.Port = "8443"
	config.File.Certificate = "cert.pem"
	config.File.PrivateKey = "private.key"
	config.File.Messages = "messages.json"

	if err := config.getEnvData(); err != nil {
		log.Println("ERROR :: Retrieving Env variables")
		return err
	}

	if err := config.getFileData(); err != nil {
		log.Println("ERROR :: Retrieving files")
		return err
	}

	configuration = &config
	return nil
}

func (c *config) GetEndpoint() string {
	return fmt.Sprintf("https://%s:%s/%s", c.Url.Endpoint, c.Url.Port, c.Url.Path)
}

func (c *config) Message(entry string) string {
	var msg, exist = c.MessageDict[entry]
	if exist {
		return msg
	}

	return fmt.Sprintf("No message for '%s' entry", entry)
}

func (c *config) getEnvData() (err error) {
	if c.BotToken, err = getEnvVariable("BOT_TOKEN"); err != nil {
		return err
	}

	if c.Url.Endpoint, err = getEnvVariable("ENDPOINT"); err != nil {
		return err
	}

	return
}

func (c *config) getFileData() (err error) {
	c.CertBytes, err = getFileBytes(c.File.Certificate)
	if err != nil {
		log.Println("ERROR :: Accesing Certificate file")
		return
	}

	ba, err := getFileBytes(c.File.Messages)
	if err != nil {
		log.Println("ERROR :: Accesing Messages dictionary file")
		return
	}
	json.Unmarshal(ba, &c.MessageDict)

	return
}

func getEnvVariable(name string) (string, error) {
	var value, isSet = os.LookupEnv(name)
	if !isSet {
		return "", fmt.Errorf("ERROR :: %s was not set", name)
	}

	return value, nil
}

func getFileBytes(filename string) ([]byte, error) {
	var file, err = os.Open(filename)
	if err != nil {
		log.Println("ERROR :: Opening file")
		return []byte{}, err
	}

	defer file.Close()

	return io.ReadAll(file)
}
