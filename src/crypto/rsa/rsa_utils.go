/*
this module is a collection of RSA utilities.
RFC: https://www.rfc-editor.org/rfc/rfc8017.html
based on: https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a
*/
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	bitSize := 2048
	privies, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		checkError(err)
	}
	return privies, &privies.PublicKey
}

func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	encodeToMemory := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return encodeToMemory
}

func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	publicKey, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		checkError(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKey,
	})

	return pubBytes
}

func BytesToPrivateKey(bytes []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(bytes)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			checkError(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		checkError(err)
	}
	return key
}

func BytesToPublicKey(bytes []byte) *rsa.PublicKey {
	block, _ := pem.Decode(bytes)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			checkError(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		checkError(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		checkError(err)
	}
	return key
}

func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		checkError(err)
	}
	return ciphertext
}

func DecryptWithPrivateKey(ciphertext []byte, key *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, key, ciphertext, nil)
	if err != nil {
		checkError(err)
	}
	return plaintext
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
