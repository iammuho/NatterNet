package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

// loadPublicKey loads the PUBLIC RSA file from directory
func (j *jwt) loadPublicKey() {
	var err error
	var ok bool

	rsaPublicKey := loadRSAFile(j.options.PublicKeyPath)

	var parsedKey interface{}

	if rsaPublicKey.Type != "RSA PUBLIC KEY" {
		log.Fatal("RSA Public Key File is wrong")
	}

	if parsedKey, err = x509.ParsePKCS1PublicKey(rsaPublicKey.Bytes); err != nil {
		log.Fatal("Unable to decode RSA Public key")
	}

	var pubKey *rsa.PublicKey
	if pubKey, ok = parsedKey.(*rsa.PublicKey); !ok {
		log.Fatal("Unable to parse RSA Public key")
	}

	j.publicKey = pubKey
}

// loadPrivateKey loads the PRIVATE RSA file from directory
func (j *jwt) loadPrivateKey() {

	var err error
	var ok bool

	rsaPrivKey := loadRSAFile(j.options.PrivateKeyPath)

	if rsaPrivKey.Type != "RSA PRIVATE KEY" {
		log.Fatal("RSA Private Key File is wrong")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(rsaPrivKey.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(rsaPrivKey.Bytes); err != nil { // note this returns type `interface{}`
			log.Fatal("Couldn't Decode the Private Pem File")
		}
	}

	var privateKey *rsa.PrivateKey
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("Unable to parse RSA private key")
	}

	j.privateKey = privateKey
}

// loadRSAFile loads the RSA file from directory
func loadRSAFile(filePath string) *pem.Block {

	if filePath == "" {
		log.Fatal("No file for RSA Certificate")
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("No RSA Certificate file found")
	}

	decodedPem, _ := pem.Decode(file)

	return decodedPem
}
