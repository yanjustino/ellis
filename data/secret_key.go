package data

import (
	"bufio"
	"ellis.com/application"
	"ellis.com/crypto/keys"
	"ellis.com/crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type SecretKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Holder struct {
	Items []SecretKey
}

func StoreSecretKey(path string, key string, value string) {
	jwk := keys.JsonToJwk(application.ReadFile(path))
	secretKey := prepareRegister(jwk, key, value)
	secretsByKey := groupSecretsByKey(path, secretKey)

	var values []string
	for _, v := range secretsByKey {
		toByteArray, e := json.Marshal(v)
		application.HandleError(e)

		values = append(values, string(toByteArray))
	}

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	content := strings.Join(values, "\n")

	e := application.CreateFile(fileName, []byte(content))
	application.HandleError(e)
}

func FetchAllSecretKeys(jwkFileName string) {
	holder := getSecretKeysHolder(jwkFileName)

	for i, s := range holder.Items {
		line := fmt.Sprintf("[%v] key: %v - value: %v", i, s.Key, s.Value)
		fmt.Println(line)
	}
}

func FetchSecretKeysHolder(jwkFileName string, eject bool) {
	holder := getSecretKeysHolder(jwkFileName)

	result, e := json.MarshalIndent(&holder, "", " ")
	application.HandleError(e)

	if eject {
		ejectSecretKeyHolder(jwkFileName, result)
	} else {
		println(string(result))
	}
}

// PRIVATE SCOPE

func ejectSecretKeyHolder(path string, holder []byte) {
	data := application.ReadFile(path)
	jwk := keys.JsonToJwk(data)

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "settings.json")

	e := application.CreateFile(fileName, holder)
	application.HandleError(e)
}

func groupSecretsByKey(jwkFileName string, defaultSecret SecretKey) map[string]SecretKey {
	holder := getSecretKeysHolder(jwkFileName)

	secretKeys := append(holder.Items, defaultSecret)
	mappings := map[string]SecretKey{}

	for _, s := range secretKeys {
		mappings[s.Key] = s
	}

	return mappings
}

func getSecretKeysHolder(path string) Holder {
	data := application.ReadFile(path)
	jwk := keys.JsonToJwk(data)

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	if !application.FileExists(fileName) {
		return Holder{Items: []SecretKey{}}
	}

	list := parseFileToSecretKeyArray(fileName)
	return Holder{Items: list}
}

func parseFileToSecretKeyArray(path string) []SecretKey {
	var list []SecretKey

	file := application.OpenFile(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		secret := parseBytesToSecret(scanner.Bytes())
		list = append(list, secret)
	}

	e := scanner.Err()
	application.HandleError(e)

	return list
}

func parseBytesToSecret(bytes []byte) SecretKey {
	var secret SecretKey
	e := json.Unmarshal(bytes, &secret)
	application.HandleError(e)
	return secret
}

func prepareRegister(jwk keys.Jwk, key string, value string) SecretKey {
	publicKey := keys.LoadPublicKey(jwk)

	encryptWithPublicKey := rsa.EncryptWithPublicKey([]byte(value), publicKey)
	encodeToString := base64.StdEncoding.EncodeToString(encryptWithPublicKey)

	return SecretKey{Key: key, Value: encodeToString}
}
