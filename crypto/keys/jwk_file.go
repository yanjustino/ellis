package keys

import (
	"encoding/json"
	"github.com/google/uuid"
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
		Exp: name,
		Kid: uuid.New().String(),
	}
}

func JwkToJson(file Jwk) string {

	b, err := json.MarshalIndent(file, "", " ")

	if err != nil {
		checkError(err)
	}

	return string(b)
}
