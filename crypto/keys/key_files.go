package keys

import (
	key "crypto/rsa"
	"ellis.io/crypto/rsa"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteKeys(name string) {
	private, public := rsa.GenerateKeyPair(1024)
	SavePrivateKey(private, name)
	SavePublicKey(public, name)
}

func SavePrivateKey(key *key.PrivateKey, name string) {
	bytes := rsa.PrivateKeyToBytes(key)
	bytesToString := b64.StdEncoding.EncodeToString(bytes)

	data := []byte(bytesToString)

	err := ioutil.WriteFile(fmt.Sprintf("%v.%v", name, "key"), data, 0)

	if err != nil {
		checkError(err)
	}
}

func SavePublicKey(key *key.PublicKey, name string) {
	keyToBytes := rsa.PublicKeyToBytes(key)
	bytesToString := b64.StdEncoding.EncodeToString(keyToBytes)

	data := []byte(bytesToString)

	err := ioutil.WriteFile(fmt.Sprintf("%v.%v", name, "json"), data, 0)

	if err != nil {
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
