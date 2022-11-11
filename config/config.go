package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"math/big"
)

type Config struct {
	Certificate []byte
	PrivateKey  []byte
}

func Setup() (Config, error) {
	var config = Config{}

	var err error = config.setKeyPair()

	return config, err
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
		SerialNumber:          serial,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	return template, nil
}
