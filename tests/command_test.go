package tests

import (
	"ellis.com/application"
	"ellis.com/command"
	"os"
	"testing"
)

func TestAllCommands(t *testing.T) {
	cleanAll()
	generateKeys(t)
	setKeys(t)
	getKeys(t)
	listKeys(t)
	viewKeys(t)
	ejectSecret(t)
}

func generateKeys(t *testing.T) {
	cmd := command.KeyGenerateArgs{Args: []string{"ellis", "keys", "-g", "project"}}
	ok := cmd.Execute()

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
	cmd := command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "USER", "Username"}}
	ok := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}

	cmd = command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "USER", "$625asdf%@"}}
	ok = cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}

	cmd = command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "PASS", "Server=myServerName\\myInstanceName;Database=myDataBase;User Id=myUsername;Password=myPassword;"}}
	ok = cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}

	if !fileExists("project.data") {
		t.Errorf("data file was not created")
	}
}

func getKeys(t *testing.T) {
	set := command.SetSecretArgs{Args: []string{"ellis", "set", "-k", "project.json", "MYSEC", "$625asdf%@"}}
	if !set.Execute() {
		t.Errorf("Fail when try set secret")
	}

	get := command.GetSecretArgs{Args: []string{"ellis", "get", "-k", "project.json", "MYSEC"}}

	if !get.CanExecute() {
		t.Errorf("Inválid Get command")
	}

	if !get.Execute() {
		t.Errorf("Fail when try set secret")
	}
}

func listKeys(t *testing.T) {
	cmd := command.ListSecretsArgs{Args: []string{"ellis", "list", "-k", "project.json"}}
	ok := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}
}

func viewKeys(t *testing.T) {
	cmd := command.ViewSecretsArgs{Args: []string{"ellis", "view", "-k", "project.json"}}
	ok := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try set secret")
	}
}

func ejectSecret(t *testing.T) {
	cmd := command.EjectSecretsArgs{Args: []string{"ellis", "eject", "-k", "project.json"}}
	ok := cmd.Execute()
	if !ok {
		t.Errorf("Fail when try eject secret")
		return
	}

	if !fileExists("project.settings.json") {
		t.Errorf("Fail: project.settings.json was not CREATED")
	}
}

func cleanAll() {
	if fileExists("project.key") {
		application.RemoveFile("project.key")
	}

	if fileExists("project.json") {
		application.RemoveFile("project.json")
	}

	if fileExists("project.data") {
		application.RemoveFile("project.data")
	}

	if fileExists("project.settings.json") {
		application.RemoveFile("project.settings.json")
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
