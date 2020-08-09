package tests

import (
	"ellis.io/command"
	"os"
	"strings"
	"testing"
)

func TestAllCommands(t *testing.T)  {
	cleanAll()
	generateKeys(t)
	listKeys(t)
	setKeys(t)
	listKeys(t)
	viewKeys(t)
	ejectSecret(t)
}

func generateKeys(t *testing.T) {
	cmd := command.KeyGenerateArgs{Args: []string{"ellis", "keys", "-g", "project"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ := cmd.Execute()

	if !ok {
		t.Errorf("Fail while generate keys")
	}

	if !fileExists("project.key") {
		t.Errorf("private key not created")
	}

	if !fileExists("project.json") {
		t.Errorf("private key not created")
	}
}

func setKeys(t *testing.T) {
	cmd := command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "keyA", "valueA"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}

	cmd = command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "keyB", "valueB"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ = cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}

	if !fileExists("project.data") {
		t.Errorf("data file was not created")
	}

}

func listKeys(t *testing.T)  {
	cmd := command.ListSecretsArgs{Args: []string{"ellis", "list", "-k", "project.json"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}
}

func viewKeys(t *testing.T)  {
	cmd := command.ViewSecretsArgs{Args: []string{"ellis", "view", "-k", "project.json"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}
}

func ejectSecret(t *testing.T)  {
	cmd := command.EjectSecretsArgs{Args: []string{"ellis", "eject", "-k", "project.json"}}
	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}

	ok, _ := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try eject secret")
		return
	}

	if fileExists("project.json"){
		t.Errorf("Fail: project.json was not DELETED")
	}

	if fileExists("project.data"){
		t.Errorf("Fail: project.data was not DELETED")
	}

	if !fileExists("project.settings.json"){
		t.Errorf("Fail: project.settings.json was not CREATED")
	}

}

func cleanAll() {
	if fileExists("project.key") {
		os.Remove("project.key")
	}

	if fileExists("project.json") {
		os.Remove("project.json")
	}

	if fileExists("project.data") {
		os.Remove("project.data")
	}

	if fileExists("project.settings.json") {
		os.Remove("project.settings.json")
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
