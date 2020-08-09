package main

import (
	"ellis.io/command"
	"os"
)

func main() {
	command.KeyGenerateArgs{Args: os.Args}.Execute()
	command.SetSecretArgs{Args: os.Args}.Execute()
	command.ListSecretsArgs{Args: os.Args}.Execute()
	command.ViewSecretsArgs{Args: os.Args}.Execute()
}
