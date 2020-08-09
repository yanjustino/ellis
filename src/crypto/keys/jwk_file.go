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

func JwkToJson(file Jwk) string {

	b, err := json.Marshal(file)

	if err != nil {
		checkError(err)
	}

	return string(b)
}

func JwkFromJson(filename string) Jwk {
	value, err := ioutil.ReadFile(filename)
	checkError(err)

	var result Jwk
	err2 := json.Unmarshal(value, &result)
	checkError(err2)

	return result
}
