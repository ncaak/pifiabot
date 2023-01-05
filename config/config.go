package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

type config struct {
	BotToken     string
	CertFilePath string
	CertBytes    []byte
	KeyFilePath  string
	Url          struct {
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
	config.Url.Path = "updates"
	config.Url.Port = "8443"
	config.CertFilePath = "cert.pem"
	config.KeyFilePath = "private.key"

	if err := config.getEnvData(); err != nil {
		log.Println("ERROR :: Retrieving Env variables")
		return err
	}

	b, err := config.getCertificateBytes()
	if err != nil {
		log.Println("ERROR :: Reading Certificate file")
	}
	config.CertBytes = b

	configuration = &config
	return nil
}

func (c *config) GetEndpoint() string {
	return fmt.Sprintf("https://%s:%s/%s", c.Url.Endpoint, c.Url.Port, c.Url.Path)
}

func (c *config) getCertificateBytes() ([]byte, error) {
	var file, err = os.Open(c.CertFilePath)
	if err != nil {
		log.Println("ERROR :: Opening Certificate file")
		return []byte{}, err
	}

	defer file.Close()

	return io.ReadAll(file)
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

func getEnvVariable(name string) (string, error) {
	var value, isSet = os.LookupEnv(name)
	if !isSet {
		return "", fmt.Errorf("ERROR :: %s was not set", name)
	}

	return value, nil
}
