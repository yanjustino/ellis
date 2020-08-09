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
	private, public := rsa.GenerateKeyPair()
	savePrivateKey(private, name)
	savePublicKey(public, name)
}

func LoadPublicKey(jwk Jwk) *key.PublicKey {
	toByte, err := b64.RawURLEncoding.DecodeString(jwk.Mod)
	checkError(err)

	return rsa.BytesToPublicKey(toByte)
}

func savePrivateKey(key *key.PrivateKey, name string) {
	bytes := rsa.PrivateKeyToBytes(key)
	bytesToString := b64.StdEncoding.EncodeToString(bytes)

	data := []byte(bytesToString)

	filename := fmt.Sprintf("%v.%v", name, "key")
	os.Remove(filename)
	err := ioutil.WriteFile(filename, data, 0777)

	if err != nil {
		checkError(err)
	}
}

func savePublicKey(key *key.PublicKey, name string) {
	keyToBytes := rsa.PublicKeyToBytes(key)
	bytesToString := b64.RawURLEncoding.EncodeToString(keyToBytes)

	jwk := NewJwk(bytesToString, name)
	json := JwkToJson(jwk)
	data := []byte(json)

	filename := fmt.Sprintf("%v.%v", name, "json")
	os.Remove(filename)
	err := ioutil.WriteFile(filename, data, 0777)

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
