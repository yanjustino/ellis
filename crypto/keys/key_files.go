package keys

import (
	key "crypto/rsa"
	"ellis.com/application"
	"ellis.com/crypto/rsa"
	b64 "encoding/base64"
	"fmt"
)

func WriteKeys(name string) {
	private, public := rsa.GenerateKeyPair()
	savePrivateKey(private, name)
	savePublicKey(public, name)
}

func LoadPublicKey(jwk Jwk) *key.PublicKey {
	toByte, err := b64.RawURLEncoding.DecodeString(jwk.Mod)
	application.HandleError(err)

	return rsa.BytesToPublicKey(toByte)
}

func LoadPrivateKey(jwk Jwk) *key.PrivateKey {
	bytes := application.ReadFile(jwk.Kid + ".key")
	privateKey := rsa.BytesToPrivateKey(bytes)
	return privateKey
}

func savePrivateKey(key *key.PrivateKey, name string) {
	bytes := rsa.PrivateKeyToBytes(key)
	filename := fmt.Sprintf("%v.%v", name, "key")

	application.RemoveFile(filename)
	err := application.CreateFile(filename, bytes)
	application.HandleError(err)
}

func savePublicKey(key *key.PublicKey, name string) {
	keyToBytes := rsa.PublicKeyToBytes(key)
	bytesToString := b64.RawURLEncoding.EncodeToString(keyToBytes)
	jwk := NewJwk(bytesToString, name)

	json := JwkToJson(jwk)
	bytes := []byte(json)

	filename := fmt.Sprintf("%v.%v", name, "json")

	application.RemoveFile(filename)
	err := application.CreateFile(filename, bytes)
	application.HandleError(err)
}
