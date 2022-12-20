package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

type Config struct {
	Certificate []byte
	PrivateKey  []byte
	BotToken    string
	Url         struct {
		Endpoint string
		Path     string
		Port     string
	}
}

var configuration *Config

func Get() *Config {
	return configuration
}

func Setup() error {
	var config = Config{}
	config.Url.Path = "/v1/bot-api"
	config.Url.Port = "8443"

	if err := config.getEnvData(); err != nil {
		log.Println("ERROR :: Retrieving Env variables")
		return err
	}

	if err := config.setKeyPair(); err != nil {
		log.Println("ERROR :: Setting up Server keys")
		return err
	}

	configuration = &config
	return nil
}

func (c *Config) getEnvData() (err error) {
	if c.BotToken, err = getEnvVariable("BOT_TOKEN"); err != nil {
		return err
	}

	if c.Url.Endpoint, err = getEnvVariable("ENDPOINT"); err != nil {
		return err
	}

	return
}

func (c *Config) setKeyPair() error {
	var rsaKey, errGenKey = rsa.GenerateKey(rand.Reader, 2048)
	if errGenKey != nil {
		return errors.New("ERROR :: Generating RSA key : " + errGenKey.Error())
	}
	var privateKey = x509.MarshalPKCS1PrivateKey(rsaKey)
	c.PrivateKey = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKey})

	var certificate, errCert = getCertificateBytes(rsaKey)
	if errCert != nil {
		return errCert
	}
	c.Certificate = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate})

	return nil
}

func getCertificateBytes(privateKey *rsa.PrivateKey) ([]byte, error) {
	var template, errTemplate = getCertificateTemplate()
	if errTemplate != nil {
		log.Println("ERROR :: Creating certificate's template")
		return []byte{}, errTemplate
	}

	var certBytes, errCert = x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if errCert != nil {
		log.Println("ERROR :: Creating certificate")
		return []byte{}, errCert
	}

	return certBytes, nil
}

func getCertificateTemplate() (x509.Certificate, error) {
	var limit = new(big.Int).Lsh(big.NewInt(1), 128)
	var serial, err = rand.Int(rand.Reader, limit)
	if err != nil {
		return x509.Certificate{}, errors.New("ERROR :: Generating random serial : " + err.Error())
	}

	var template = x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Organization: []string{"pifiabot"},
			Country:      []string{"ES"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	return template, nil
}

func getEnvVariable(name string) (string, error) {
	var value, isSet = os.LookupEnv(name)
	if !isSet {
		return "", fmt.Errorf("ERROR :: %s was not set", name)
	}

	return value, nil
}
