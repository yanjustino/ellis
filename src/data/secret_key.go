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

const SecretFolder = "secrets/"

func StoreSecretKey(jwkFileName string, key string, value string) {
	jwk := keys.JwkFromJsonFilePath(SecretFolder + jwkFileName)
	new := prepareRegister(jwk, key, value)
	mps := groupSecretKeysByKey(jwkFileName, new)

	var values []string
	for _, v := range mps {
		json := parseSecretToByteArray(v)
		values = append(values, string(json))
	}

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	content := strings.Join(values, "\n")

	createFileIfNonExists(fileName, []byte(content))
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
	utils.CheckError(e)

	if eject {
		ejectSecretKeyHolder(jwkFileName, result)
	} else {
		println(string(result))
	}
}

// PRIVATE SCOPE

func ejectSecretKeyHolder(jwkFileName string, holder []byte) {
	jwk := keys.JwkFromJsonFilePath(SecretFolder + jwkFileName)
	fileName := SecretFolder + fmt.Sprintf("%v.%v", jwk.Kid, "settings.json")
	fileData := SecretFolder + fmt.Sprintf("%v.%v", jwk.Kid, "data")

	createFileIfNonExists(fileName, holder)

	errJwk := os.Remove(SecretFolder + jwkFileName)
	utils.CheckError(errJwk)

	errData := os.Remove(SecretFolder + fileData)
	utils.CheckError(errData)
}

func groupSecretKeysByKey(jwkFileName string, fallbackSecret SecretKey) map[string]SecretKey {
	holder := getSecretKeysHolder(jwkFileName)
	list := append(holder.Items, fallbackSecret)

	m := map[string]SecretKey{}
	for _, s := range list {
		m[s.Key] = s
	}

	return m
}

func getSecretKeysHolder(jwkFileName string) Holder {
	jwk := keys.JwkFromJsonFilePath(jwkFileName)

	fileName := SecretFolder + fmt.Sprintf("%v.%v", jwk.Kid, "data")
	if !fileExists(fileName) {
		return Holder{Items: []SecretKey{}}
	}

	file, err := os.Open(fileName)
	utils.CheckError(err)
	defer file.Close()

	list := parseFileToSecretKeyArray(file)
	return Holder{Items: list}
}

func parseFileToSecretKeyArray(file *os.File) []SecretKey {
	var list []SecretKey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		secret := parseBytesToSecret(scanner.Bytes())
		list = append(list, secret)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func parseBytesToSecret(bytes []byte) SecretKey {
	var secret SecretKey
	err := json.Unmarshal(bytes, &secret)
	utils.CheckError(err)
	return secret
}

func parseSecretToByteArray(key SecretKey) []byte {
	json, err := json.Marshal(&key)
	utils.CheckError(err)
	return json
}

func prepareRegister(jwk keys.Jwk, key string, value string) SecretKey {
	publicKey := keys.LoadPublicKey(jwk)
	encryptWithPublicKey := rsa.EncryptWithPublicKey([]byte(value), publicKey)
	encodeToString := base64.StdEncoding.EncodeToString(encryptWithPublicKey)

	return SecretKey{Key: key, Value: encodeToString}
}

func createFileIfNonExists(fileName string, content []byte) {
	data := []byte(string(content) + "\n")
	e := ioutil.WriteFile(SecretFolder+fileName, data, os.FileMode.Perm(0644))
	utils.CheckError(e)
}

func createDirectoryIfNonExists() {
	_, err := os.Stat(SecretFolder)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("secrets", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
