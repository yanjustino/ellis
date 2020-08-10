package keys

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"strings"
)

type Jwk struct {
	Alg string `json:"alg"`
	Mod string `json:"mod"`
	Exp string `json:"exp"`
	Kid string `json:"kid"`
}

func NewJwk(mod string, name string) Jwk {
	return Jwk{
		Alg: "RSA",
		Mod: mod,
		Exp: strings.Replace(uuid.New().String(), "-", "", -1),
		Kid: name,
	}
}

func JwkToJson(jwk Jwk) (string, error) {
	b, err := json.Marshal(jwk)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func JwkFromJsonFilePath(filename string) (Jwk, error) {
	value, err := ioutil.ReadFile(filename)
	if err != nil {
		return Jwk{}, err
	}
	return JwkFromJsonByteArray(value)
}

func JwkFromJsonByteArray(value []byte) (Jwk, error) {
	var result Jwk

	err := json.Unmarshal(value, &result)
	if err != nil {
		return Jwk{}, err
	}

	return result, nil
}
