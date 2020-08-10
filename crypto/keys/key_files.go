package keys

import (
	key "crypto/rsa"
	"ellis.com/crypto/rsa"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteKeys(name string) error {
	private, public := rsa.GenerateKeyPair()

	epriv := savePrivateKey(private, name)
	if epriv != nil {
		return epriv
	}

	epub := savePublicKey(public, name)
	if epub != nil {
		return epub
	}

	return nil
}

func LoadPublicKey(jwk Jwk) (key *key.PublicKey, err error) {
	toByte, err := b64.RawURLEncoding.DecodeString(jwk.Mod)

	if err != nil {
		return nil, err
	}

	return rsa.BytesToPublicKey(toByte), nil
}

func savePrivateKey(key *key.PrivateKey, name string) error {
	bytes := rsa.PrivateKeyToBytes(key)
	//bytesToString := b64.StdEncoding.EncodeToString(bytes)

	//data := []byte(bytesToString)

	filename := fmt.Sprintf("%v.%v", name, "key")
	os.Remove(filename)
	err := ioutil.WriteFile(filename, bytes, 0777)

	if err != nil {
		return err
	}
	return nil
}

func savePublicKey(key *key.PublicKey, name string) error {
	keyToBytes := rsa.PublicKeyToBytes(key)
	bytesToString := b64.RawURLEncoding.EncodeToString(keyToBytes)

	jwk := NewJwk(bytesToString, name)

	json, e := JwkToJson(jwk)
	if e != nil {
		return e
	}

	data := []byte(json)

	filename := fmt.Sprintf("%v.%v", name, "json")
	os.Remove(filename)
	err := ioutil.WriteFile(filename, data, 0777)

	if err != nil {
		return err
	}

	return nil
}
