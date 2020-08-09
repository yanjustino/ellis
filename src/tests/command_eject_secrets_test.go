package tests

import (
	"ellis.io/command"
	"strings"
	"testing"
)

// Eject Command Can Be Executed
func TestEjectCommandCanBeExecute(t *testing.T) {
	cmd := command.EjectSecretsArgs{ Args: []string{"ellis", "eject", "-k", "project"}}

	if !cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}
}

func TestEjectCommandCanNotBeExecute(t *testing.T) {
	cmd := command.EjectSecretsArgs{ Args: []string{"ellis", "eject", "-k"}}

	if cmd.CanExecute() {
		t.Errorf("Invalid arguments [%v]", strings.Join(cmd.Args, " "))
	}
}
