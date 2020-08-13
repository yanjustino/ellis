package tests

import (
	"ellis.com/command"
	"strings"
	"testing"
)

func TestCanExecuteSetKeys(t *testing.T) {
	var args []string

	args = []string{"ellis", "set", "-k", "project.json", "KeyA", "1234"}
	executeCommand(t, args)

	args = []string{"ellis", "set", "-k", "project.json", "KeyA", "Only letters secret"}
	executeCommand(t, args)

	args = []string{"ellis", "set", "-k", "project.json", "KeyA", "p@ssw0rd#2020b3"}
	executeCommand(t, args)

	args = []string{"ellis", "set", "-k", "folder/project.json", "Key", "12$%*&3@#)("}
	executeCommand(t, args)

	args = []string{"ellis", "set", "-k", "folder/project.json", "Key.z", "Value_A"}
	executeCommand(t, args)

	args = []string{"ellis", "set", "-k", "folder/project.json", "Key", "Server=myServerName\\myInstanceName;Database=myDataBase;User Id=myUsername;Password=myPassword;"}
	executeCommand(t, args)
}

func executeCommand(t *testing.T, args []string) {
	cmd := command.SetSecretArgs{Args: args}

	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(args, " "))
	} else {
		t.Logf("Valid arguments [%v]", strings.Join(args, " "))
	}
}
