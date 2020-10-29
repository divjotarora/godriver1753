package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

var (
	regularCert   = readFile("cert.pem")
	regularKey    = readFile("key.pem")
	passwordCert  = readFile("cert_pass.pem")
	passwordKey   = readEncryptedKey("key_pass.pem")
	combinedCerts = readFile("combined_cert.pem")
)

func main() {
	_, err := tls.X509KeyPair(regularCert, regularKey)
	if err != nil {
		log.Printf("regular err: %v", err)
	}

	_, err = tls.X509KeyPair(passwordCert, passwordKey)
	if err != nil {
		log.Printf("password err: %v", err)
	}

	_, err = tls.X509KeyPair(combinedCerts, regularKey)
	if err != nil {
		log.Printf("combined + regular err: %v", err)
	}

	// Fails with "private key does not match public key"
	_, err = tls.X509KeyPair(combinedCerts, passwordKey)
	if err != nil {
		log.Printf("combined + password err: %v", err)
	}
}

func readFile(path string) []byte {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return res
}

func readEncryptedKey(path string) []byte {
	data := readFile(path)

	block, _ := pem.Decode(data)
	buf, err := x509.DecryptPEMBlock(block, []byte("passphrase"))
	if err != nil {
		panic(err)
	}

	var encoded bytes.Buffer
	pem.Encode(&encoded, &pem.Block{Type: block.Type, Bytes: buf})
	return encoded.Bytes()
}
