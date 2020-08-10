package data

import (
	"bufio"
	"ellis.com/crypto/keys"
	"ellis.com/crypto/rsa"
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

func StoreSecretKey(jwkFileName string, key string, value string) error {
	jwk, e := keys.JwkFromJsonFilePath(jwkFileName)
	if e != nil {
		return e
	}

	register, e := prepareRegister(jwk, key, value)
	if e != nil {
		return e
	}

	mps, e := groupSecretKeysByKey(jwkFileName, register)
	if e != nil {
		return e
	}

	var values []string
	for _, v := range mps {
		toByteArray, e := parseSecretToByteArray(v)
		if e != nil {
			return e
		}

		values = append(values, string(toByteArray))
	}

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	content := strings.Join(values, "\n")

	return createFileIfNonExists(fileName, []byte(content))
}

func FetchAllSecretKeys(jwkFileName string) error {
	holder, e := getSecretKeysHolder(jwkFileName)
	if e != nil {
		return e
	}

	for i, s := range holder.Items {
		line := fmt.Sprintf("[%v] key: %v - value: %v", i, s.Key, s.Value)
		fmt.Println(line)
	}

	return nil
}

func FetchSecretKeysHolder(jwkFileName string, eject bool) error {
	holder, e := getSecretKeysHolder(jwkFileName)
	if e != nil {
		return e
	}

	result, e := json.MarshalIndent(&holder, "", " ")
	if e != nil {
		return e
	}

	if eject {
		return ejectSecretKeyHolder(jwkFileName, result)
	} else {
		println(string(result))
		return nil
	}
}

// PRIVATE SCOPE

func ejectSecretKeyHolder(jwkFileName string, holder []byte) error {
	jwk, e := keys.JwkFromJsonFilePath(jwkFileName)
	if e != nil {
		return e
	}

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "settings.json")
	//fileData := fmt.Sprintf("%v.%v", jwk.Kid, "data")

	return createFileIfNonExists(fileName, holder)

	//e = os.Remove(jwkFileName)
	//if e != nil {
	//	return e
	//}
	//
	//e = os.Remove(fileData)
	//if e != nil {
	//	return e
	//}
}

func groupSecretKeysByKey(jwkFileName string, fallbackSecret SecretKey) (map[string]SecretKey, error) {
	holder, e := getSecretKeysHolder(jwkFileName)
	if e != nil {
		return nil, e
	}

	list := append(holder.Items, fallbackSecret)

	m := map[string]SecretKey{}
	for _, s := range list {
		m[s.Key] = s
	}

	return m, nil
}

func getSecretKeysHolder(jwkFileName string) (Holder, error) {
	jwk, e := keys.JwkFromJsonFilePath(jwkFileName)
	if e != nil {
		return Holder{}, e
	}

	fileName := fmt.Sprintf("%v.%v", jwk.Kid, "data")
	if !fileExists(fileName) {
		return Holder{Items: []SecretKey{}}, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return Holder{}, err
	}

	defer file.Close()

	list := parseFileToSecretKeyArray(file)
	return Holder{Items: list}, nil
}

func parseFileToSecretKeyArray(file *os.File) []SecretKey {
	var list []SecretKey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		secret, _ := parseBytesToSecret(scanner.Bytes())
		list = append(list, secret)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func parseBytesToSecret(bytes []byte) (SecretKey, error) {
	var secret SecretKey
	err := json.Unmarshal(bytes, &secret)
	return secret, err
}

func parseSecretToByteArray(key SecretKey) ([]byte, error) {
	return json.Marshal(&key)
}

func prepareRegister(jwk keys.Jwk, key string, value string) (SecretKey, error) {
	publicKey, err := keys.LoadPublicKey(jwk)
	if err != nil {
		return SecretKey{}, err
	}

	encryptWithPublicKey := rsa.EncryptWithPublicKey([]byte(value), publicKey)
	encodeToString := base64.StdEncoding.EncodeToString(encryptWithPublicKey)

	return SecretKey{Key: key, Value: encodeToString}, nil
}

func createFileIfNonExists(fileName string, content []byte) error {
	data := []byte(string(content) + "\n")
	return ioutil.WriteFile(fileName, data, os.FileMode.Perm(0644))
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
