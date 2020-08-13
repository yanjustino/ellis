package keys

import (
	"ellis.com/application"
	"encoding/json"
	"github.com/google/uuid"
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

func JwkToJson(jwk Jwk) string {
	b, e := json.Marshal(jwk)
	application.HandleError(e)
	return string(b)
}

func JsonToJwk(value []byte) Jwk {
	var result Jwk
	e := json.Unmarshal(value, &result)
	application.HandleError(e)
	return result
}
