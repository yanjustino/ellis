package data

import (
	"bufio"
	"ellis.io/crypto/keys"
	"ellis.io/crypto/rsa"
	"ellis.io/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type SecretKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Holder struct {
	Items []SecretKey
}

func StoreSecretKey(jwkFileName string, key string, value string) {
	jwk := keys.JwkFromJson(jwkFileName)
	new := prepareRegister(jwk, key, value)

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")

	holder := viewSecretKeysHolder(jwkFileName)
	list := append(holder.Items, new)

	m := map[string]SecretKey{}
	for _, s := range list {
		m[s.Key] = s
	}

	var values []string
	for _, v := range m {
		json, err := json.Marshal(&v)
		utils.CheckError(err)
		values = append(values, string(json))
	}

	content := strings.Join(values, "\n")

	createFileIfNonExists(fileName, []byte(content))
}

func FetchAllSecretKeys(jwkFileName string) {
	holder := viewSecretKeysHolder(jwkFileName)

	for i, s := range holder.Items {
		line := fmt.Sprintf("[%v] key: %v - value: %v", i, s.Key, s.Value)
		fmt.Println(line)
	}
}

func FetchSecretKeysHolder(jwkFileName string) {
	holder := viewSecretKeysHolder(jwkFileName)
	result, e := json.MarshalIndent(&holder, "", " ")
	utils.CheckError(e)

	println(string(result))
}

func viewSecretKeysHolder(jwkFileName string) Holder {
	jwk := keys.JwkFromJson(jwkFileName)

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	if fileExists(fileName) {
		file, err := os.Open(fileName)
		utils.CheckError(err)
		defer file.Close()

		var list []SecretKey
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var secret SecretKey
			err := json.Unmarshal(scanner.Bytes(), &secret)
			utils.CheckError(err)

			list = append(list, secret)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return Holder{Items: list}
	}

	return Holder{Items: []SecretKey{}}
}

func prepareRegister(jwk keys.Jwk, key string, value string) SecretKey {
	publicKey := keys.LoadPublicKey(jwk)
	encryptWithPublicKey := rsa.EncryptWithPublicKey([]byte(value), publicKey)
	encodeToString := base64.StdEncoding.EncodeToString(encryptWithPublicKey)

	return SecretKey{Key: key, Value: encodeToString}
}

func createFileIfNonExists(fileName string, content []byte) {
	data := []byte(string(content) + "\n")
	e := ioutil.WriteFile(fileName, data, os.FileMode.Perm(0644))
	utils.CheckError(e)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
